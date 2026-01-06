// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package requestconfig

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"sync"
	"time"
)

var OAuth2Cache = newOAuth2Cache("/oauth/token")

// OAuth2RefreshCache is a global cache for OAuth2 refresh token flow
var OAuth2RefreshCache = &OAuth2RefreshState{}

func newOAuth2Cache(tokenUrls ...string) map[string]*OAuth2State {
	state := make(map[string]*OAuth2State, len(tokenUrls))
	for _, url := range tokenUrls {
		state[url] = &OAuth2State{authPath: url}
	}
	return state
}

// OAuth2State represents the OAuth2 provider configuration and manages token
// caching.
type OAuth2State struct {
	authPath string
	entries  sync.Map // map[oAuth2ClientCredentials]*oAuthTokenInfo
}

type oAuth2ClientCredentials struct {
	ClientID     string
	ClientSecret string
}

type oAuthTokenInfo struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int    `json:"expires_in"`

	expiresAt     time.Time
	reqTokenMutex sync.Mutex
}

func (tok *oAuthTokenInfo) reset() {
	tok.AccessToken = ""
	tok.ExpiresIn = 0
	tok.expiresAt = time.Time{}
}

func (auth *oAuthTokenInfo) validToken() string {
	if auth.AccessToken != "" && time.Now().Before(auth.expiresAt.Add(-10*time.Second)) {
		return auth.AccessToken
	}
	return ""
}

func (state *OAuth2State) GetToken(cfg *RequestConfig) (string, error) {
	tokenInfo := state.find(cfg)

	valid := tokenInfo.validToken()
	if valid != "" {
		return valid, nil
	}

	tokenInfo.reqTokenMutex.Lock()
	defer tokenInfo.reqTokenMutex.Unlock()

	tokenInfo.reset()

	authUrl, err := state.authUrl(cfg)
	if err != nil {
		return "", err
	}

	// Create JSON body with grant_type
	bodyJSON := []byte(`{"grant_type":"client_credentials"}`)
	oAuthReq, err := http.NewRequestWithContext(cfg.Context, http.MethodPost, authUrl, strings.NewReader(string(bodyJSON)))
	if err != nil {
		return "", fmt.Errorf("requestconfig: failed to create OAuth2 token request: %w", err)
	}

	encoded := base64.StdEncoding.EncodeToString(fmt.Appendf(nil, "%s:%s", cfg.ClientID, cfg.ClientSecret))
	oAuthReq.Header.Set("Authorization", fmt.Sprintf("Basic %s", encoded))
	oAuthReq.Header.Set("Content-Type", "application/json")

	handler := cfg.HTTPClient.Do
	if cfg.CustomHTTPDoer != nil {
		handler = cfg.CustomHTTPDoer.Do
	}

	oauthRes, err := handler(oAuthReq)
	if err != nil {
		return "", fmt.Errorf("requestconfig: OAuth2 token request failed: %w", err)
	}
	defer oauthRes.Body.Close()

	if oauthRes.StatusCode != http.StatusOK {
		return "", fmt.Errorf("requestconfig: OAuth2 token request returned status %d, expected %d", oauthRes.StatusCode, http.StatusOK)
	}

	contents, err := io.ReadAll(oauthRes.Body)
	if err != nil {
		return "", fmt.Errorf("requestconfig: failed to read OAuth2 token response: %w", err)
	}

	err = json.Unmarshal(contents, tokenInfo)
	if err != nil {
		return "", fmt.Errorf("requestconfig: failed to parse OAuth2 token response: %w", err)
	}

	if tokenInfo.AccessToken == "" {
		return "", fmt.Errorf("requestconfig: OAuth2 token response missing access_token")
	}

	tokenInfo.expiresAt = time.Now().Add(time.Duration(tokenInfo.ExpiresIn) * time.Second)

	return tokenInfo.AccessToken, nil
}

func (r *OAuth2State) find(cfg *RequestConfig) *oAuthTokenInfo {
	key := oAuth2ClientCredentials{cfg.ClientID, cfg.ClientSecret}
	resp, ok := r.entries.Load(key)
	if !ok {
		resp, _ = r.entries.LoadOrStore(key, new(oAuthTokenInfo))
	}

	return resp.(*oAuthTokenInfo)
}

func (state *OAuth2State) authUrl(cfg *RequestConfig) (string, error) {
	authUrl, err := cfg.BaseURL.Parse(strings.TrimLeft(state.authPath, "/"))
	if err != nil {
		err = fmt.Errorf("requestconfig: failed to parse OAuth2 token URL: %w", err)
		return "", err
	}
	// Don't add grant_type as query param - it's sent in JSON body
	return authUrl.String(), nil
}

// OAuth2RefreshState manages OAuth2 refresh token flow with automatic refresh.
// It handles access tokens that can be refreshed using a refresh token.
type OAuth2RefreshState struct {
	mu           sync.Mutex
	entries      sync.Map // map[oauth2RefreshKey]*oauth2RefreshTokenInfo
	saveCallback func(workspace string, accessToken string, refreshToken string, expiresIn int) error
}

type oauth2RefreshKey struct {
	Workspace    string
	AccessToken  string // Initial access token (used as key)
	RefreshToken string
}

type oauth2RefreshTokenInfo struct {
	AccessToken   string
	RefreshToken  string
	DeviceCode    string
	ExpiresAt     time.Time
	TokenLifetime time.Duration
	mu            sync.Mutex
}

// refreshTokenResponse represents the response from the refresh token endpoint
type refreshTokenResponse struct {
	AccessToken  string `json:"access_token"`
	ExpiresIn    int    `json:"expires_in"`
	RefreshToken string `json:"refresh_token"`
	TokenType    string `json:"token_type"`
}

// SetSaveCallback sets the callback function to save credentials after refresh
func (state *OAuth2RefreshState) SetSaveCallback(callback func(workspace string, accessToken string, refreshToken string, expiresIn int) error) {
	state.mu.Lock()
	defer state.mu.Unlock()
	state.saveCallback = callback
}

// GetToken returns a valid access token, refreshing if necessary
func (state *OAuth2RefreshState) GetToken(cfg *RequestConfig) (string, error) {
	if cfg.AccessToken == "" {
		return "", fmt.Errorf("requestconfig: no access token provided")
	}

	tokenInfo := state.findOrCreate(cfg)

	tokenInfo.mu.Lock()
	defer tokenInfo.mu.Unlock()

	// Check if token needs refresh
	if needsRefresh(tokenInfo) {
		if err := state.doRefresh(cfg, tokenInfo); err != nil {
			return "", err
		}
	}

	return tokenInfo.AccessToken, nil
}

func (state *OAuth2RefreshState) findOrCreate(cfg *RequestConfig) *oauth2RefreshTokenInfo {
	key := oauth2RefreshKey{
		Workspace:    cfg.Workspace,
		AccessToken:  cfg.AccessToken,
		RefreshToken: cfg.RefreshToken,
	}

	if existing, ok := state.entries.Load(key); ok {
		return existing.(*oauth2RefreshTokenInfo)
	}

	// Parse JWT to get expiration info
	expTime, iatTime, err := parseJWTExpiration(cfg.AccessToken)
	tokenLifetime := time.Duration(0)
	expiresAt := time.Time{}

	if err == nil {
		tokenLifetime = expTime.Sub(iatTime)
		expiresAt = expTime
	}

	info := &oauth2RefreshTokenInfo{
		AccessToken:   cfg.AccessToken,
		RefreshToken:  cfg.RefreshToken,
		DeviceCode:    cfg.DeviceCode,
		ExpiresAt:     expiresAt,
		TokenLifetime: tokenLifetime,
	}

	actual, _ := state.entries.LoadOrStore(key, info)
	return actual.(*oauth2RefreshTokenInfo)
}

// needsRefresh checks if the token should be refreshed.
// It uses a dynamic threshold: refresh when 20% or less of token lifetime remains.
func needsRefresh(info *oauth2RefreshTokenInfo) bool {
	if info.AccessToken == "" {
		return true
	}

	if info.ExpiresAt.IsZero() {
		// Can't determine expiration, don't refresh
		return false
	}

	if info.RefreshToken == "" {
		// No refresh token available
		return false
	}

	// Calculate refresh threshold (20% of token lifetime)
	// For a 30s token: refreshes 6s before expiry
	// For a 2h token: refreshes 24min before expiry
	var refreshThreshold time.Duration
	if info.TokenLifetime > 0 {
		refreshThreshold = time.Duration(float64(info.TokenLifetime) * 0.2)
	} else {
		// Default to 10 seconds if we can't determine lifetime
		refreshThreshold = 10 * time.Second
	}

	timeUntilExpiry := time.Until(info.ExpiresAt)
	return timeUntilExpiry <= refreshThreshold
}

// parseJWTExpiration extracts exp and iat claims from a JWT token
func parseJWTExpiration(token string) (expTime time.Time, iatTime time.Time, err error) {
	parts := strings.Split(token, ".")
	if len(parts) != 3 {
		return time.Time{}, time.Time{}, fmt.Errorf("invalid JWT token format")
	}

	// Decode the claims (second part)
	claimsBytes, err := base64.RawURLEncoding.DecodeString(parts[1])
	if err != nil {
		return time.Time{}, time.Time{}, fmt.Errorf("failed to decode JWT claims: %w", err)
	}

	var claims map[string]interface{}
	if err := json.Unmarshal(claimsBytes, &claims); err != nil {
		return time.Time{}, time.Time{}, fmt.Errorf("failed to parse JWT claims: %w", err)
	}

	// Extract exp claim
	if exp, ok := claims["exp"].(float64); ok {
		expTime = time.Unix(int64(exp), 0)
	} else {
		return time.Time{}, time.Time{}, fmt.Errorf("JWT missing exp claim")
	}

	// Extract iat claim
	if iat, ok := claims["iat"].(float64); ok {
		iatTime = time.Unix(int64(iat), 0)
	} else {
		// If no iat, estimate based on expiry (assume 1 hour token)
		iatTime = expTime.Add(-1 * time.Hour)
	}

	return expTime, iatTime, nil
}

func (state *OAuth2RefreshState) doRefresh(cfg *RequestConfig, info *oauth2RefreshTokenInfo) error {
	if info.RefreshToken == "" {
		return fmt.Errorf("requestconfig: no refresh token available")
	}

	authUrl, err := cfg.BaseURL.Parse("oauth/token")
	if err != nil {
		return fmt.Errorf("requestconfig: failed to parse refresh URL: %w", err)
	}

	refreshData := map[string]string{
		"grant_type":    "refresh_token",
		"refresh_token": info.RefreshToken,
		"client_id":     "blaxel",
	}
	if info.DeviceCode != "" {
		refreshData["device_code"] = info.DeviceCode
	}

	jsonData, err := json.Marshal(refreshData)
	if err != nil {
		return fmt.Errorf("requestconfig: failed to marshal refresh data: %w", err)
	}

	req, err := http.NewRequestWithContext(cfg.Context, http.MethodPost, authUrl.String(), strings.NewReader(string(jsonData)))
	if err != nil {
		return fmt.Errorf("requestconfig: failed to create refresh request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")

	handler := cfg.HTTPClient.Do
	if cfg.CustomHTTPDoer != nil {
		handler = cfg.CustomHTTPDoer.Do
	}

	res, err := handler(req)
	if err != nil {
		return fmt.Errorf("requestconfig: refresh token request failed: %w", err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return fmt.Errorf("requestconfig: failed to read refresh response: %w", err)
	}

	if res.StatusCode >= 400 {
		if res.StatusCode == 401 || res.StatusCode == 403 {
			return fmt.Errorf("requestconfig: authentication failed (HTTP %d): %s\nFor more information on authentication, visit: https://docs.blaxel.ai/sdk-reference/introduction#how-authentication-works", res.StatusCode, string(body))
		}
		return fmt.Errorf("requestconfig: refresh token request returned status %d: %s", res.StatusCode, string(body))
	}

	var refreshResp refreshTokenResponse
	if err := json.Unmarshal(body, &refreshResp); err != nil {
		return fmt.Errorf("requestconfig: failed to parse refresh response: %w", err)
	}

	if refreshResp.AccessToken == "" {
		return fmt.Errorf("requestconfig: refresh response missing access_token")
	}

	// Update token info
	info.AccessToken = refreshResp.AccessToken
	if refreshResp.RefreshToken != "" {
		info.RefreshToken = refreshResp.RefreshToken
	}

	// Parse new token expiration
	if expTime, iatTime, err := parseJWTExpiration(refreshResp.AccessToken); err == nil {
		info.ExpiresAt = expTime
		info.TokenLifetime = expTime.Sub(iatTime)
	} else if refreshResp.ExpiresIn > 0 {
		info.ExpiresAt = time.Now().Add(time.Duration(refreshResp.ExpiresIn) * time.Second)
		info.TokenLifetime = time.Duration(refreshResp.ExpiresIn) * time.Second
	}

	// Save credentials if callback is set
	state.mu.Lock()
	saveCallback := state.saveCallback
	state.mu.Unlock()

	if saveCallback != nil && cfg.Workspace != "" {
		_ = saveCallback(cfg.Workspace, info.AccessToken, info.RefreshToken, refreshResp.ExpiresIn)
	}

	return nil
}
