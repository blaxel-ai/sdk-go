// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package blaxel_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/blaxel-ai/sdk-go"
	"github.com/blaxel-ai/sdk-go/internal/testutil"
	"github.com/blaxel-ai/sdk-go/option"
	"github.com/blaxel-ai/sdk-go/shared"
)

func TestSandboxNewWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := blaxel.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Sandboxes.New(context.TODO(), blaxel.SandboxNewParams{
		Sandbox: blaxel.SandboxParam{
			Metadata: blaxel.MetadataParam{
				Name:        "my-resource",
				DisplayName: blaxel.String("My Resource"),
				Labels: map[string]string{
					"foo": "string",
				},
			},
			Spec: blaxel.SandboxSpecParam{
				Enabled: blaxel.Bool(true),
				Lifecycle: blaxel.SandboxLifecycleParam{
					ExpirationPolicies: []blaxel.ExpirationPolicyParam{{
						Action: blaxel.ExpirationPolicyActionDelete,
						Type:   blaxel.ExpirationPolicyTypeTtlIdle,
						Value:  blaxel.String("24h"),
					}},
					TerminatedRetention: blaxel.String("24h"),
				},
				Network: blaxel.SandboxNetworkParam{
					AllowedDomains: []string{"string"},
					Egress: blaxel.SandboxNetworkEgressParam{
						Gateway: blaxel.String("egress-ip-gw-1"),
						Mode:    blaxel.String("dedicated"),
						Policies: []blaxel.SandboxNetworkEgressPolicyParam{{
							Destinations: []string{"string"},
							Mode:         blaxel.String("dedicated"),
							Name:         blaxel.String("payment-apis"),
						}},
					},
					ForbiddenDomains: []string{"string"},
					Proxy: blaxel.SandboxNetworkProxyParam{
						Bypass: []string{"string"},
						Routing: []blaxel.SandboxNetworkProxyRoutingParam{{
							Body: map[string]string{
								"0":  "{",
								"1":  "\"",
								"2":  "a",
								"3":  "p",
								"4":  "i",
								"5":  "_",
								"6":  "k",
								"7":  "e",
								"8":  "y",
								"9":  "\"",
								"10": ":",
								"11": " ",
								"12": "\"",
								"13": "{",
								"14": "{",
								"15": "S",
								"16": "E",
								"17": "C",
								"18": "R",
								"19": "E",
								"20": "T",
								"21": ":",
								"22": "s",
								"23": "t",
								"24": "r",
								"25": "i",
								"26": "p",
								"27": "e",
								"28": "-",
								"29": "k",
								"30": "e",
								"31": "y",
								"32": "}",
								"33": "}",
								"34": "\"",
								"35": "}",
							},
							Destinations: []string{"string"},
							Headers: map[string]string{
								"0":  "{",
								"1":  "\"",
								"2":  "A",
								"3":  "u",
								"4":  "t",
								"5":  "h",
								"6":  "o",
								"7":  "r",
								"8":  "i",
								"9":  "z",
								"10": "a",
								"11": "t",
								"12": "i",
								"13": "o",
								"14": "n",
								"15": "\"",
								"16": ":",
								"17": " ",
								"18": "\"",
								"19": "B",
								"20": "e",
								"21": "a",
								"22": "r",
								"23": "e",
								"24": "r",
								"25": " ",
								"26": "{",
								"27": "{",
								"28": "S",
								"29": "E",
								"30": "C",
								"31": "R",
								"32": "E",
								"33": "T",
								"34": ":",
								"35": "s",
								"36": "t",
								"37": "r",
								"38": "i",
								"39": "p",
								"40": "e",
								"41": "-",
								"42": "k",
								"43": "e",
								"44": "y",
								"45": "}",
								"46": "}",
								"47": "\"",
								"48": "}",
							},
							Secrets: map[string]string{
								"0":  "{",
								"1":  "\"",
								"2":  "s",
								"3":  "t",
								"4":  "r",
								"5":  "i",
								"6":  "p",
								"7":  "e",
								"8":  "-",
								"9":  "k",
								"10": "e",
								"11": "y",
								"12": "\"",
								"13": ":",
								"14": " ",
								"15": "\"",
								"16": "s",
								"17": "k",
								"18": "-",
								"19": "l",
								"20": "i",
								"21": "v",
								"22": "e",
								"23": "-",
								"24": "a",
								"25": "b",
								"26": "c",
								"27": "1",
								"28": "2",
								"29": "3",
								"30": ".",
								"31": ".",
								"32": ".",
								"33": "\"",
								"34": "}",
							},
						}},
					},
				},
				Region: blaxel.String("us-pdx-1"),
				Runtime: blaxel.SandboxRuntimeParam{
					Envs: []shared.EnvParam{{
						Name:   blaxel.String("MY_ENV_VAR"),
						Secret: blaxel.Bool(true),
						Value:  blaxel.String("my-value"),
					}},
					Expires: blaxel.String("2025-12-31T23:59:59Z"),
					Image:   blaxel.String("blaxel/base-image:latest"),
					Memory:  blaxel.Int(4096),
					Ports: []blaxel.PortParam{{
						Target:   8080,
						Name:     blaxel.String("http"),
						Protocol: blaxel.PortProtocolHTTP,
					}},
					TerminationGracePeriodSeconds: blaxel.Int(30),
					Ttl:                           blaxel.String("24h"),
				},
				Volumes: []blaxel.VolumeAttachmentParam{{
					MountPath: blaxel.String("/mnt/data"),
					Name:      blaxel.String("my-volume"),
					ReadOnly:  blaxel.Bool(false),
				}},
			},
		},
		CreateIfNotExist: blaxel.Bool(true),
	})
	if err != nil {
		var apierr *blaxel.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSandboxGetWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := blaxel.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Sandboxes.Get(
		context.TODO(),
		"sandboxName",
		blaxel.SandboxGetParams{
			ShowSecrets: blaxel.Bool(true),
		},
	)
	if err != nil {
		var apierr *blaxel.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSandboxUpdateWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := blaxel.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Sandboxes.Update(
		context.TODO(),
		"sandboxName",
		blaxel.SandboxUpdateParams{
			Sandbox: blaxel.SandboxParam{
				Metadata: blaxel.MetadataParam{
					Name:        "my-resource",
					DisplayName: blaxel.String("My Resource"),
					Labels: map[string]string{
						"foo": "string",
					},
				},
				Spec: blaxel.SandboxSpecParam{
					Enabled: blaxel.Bool(true),
					Lifecycle: blaxel.SandboxLifecycleParam{
						ExpirationPolicies: []blaxel.ExpirationPolicyParam{{
							Action: blaxel.ExpirationPolicyActionDelete,
							Type:   blaxel.ExpirationPolicyTypeTtlIdle,
							Value:  blaxel.String("24h"),
						}},
						TerminatedRetention: blaxel.String("24h"),
					},
					Network: blaxel.SandboxNetworkParam{
						AllowedDomains: []string{"string"},
						Egress: blaxel.SandboxNetworkEgressParam{
							Gateway: blaxel.String("egress-ip-gw-1"),
							Mode:    blaxel.String("dedicated"),
							Policies: []blaxel.SandboxNetworkEgressPolicyParam{{
								Destinations: []string{"string"},
								Mode:         blaxel.String("dedicated"),
								Name:         blaxel.String("payment-apis"),
							}},
						},
						ForbiddenDomains: []string{"string"},
						Proxy: blaxel.SandboxNetworkProxyParam{
							Bypass: []string{"string"},
							Routing: []blaxel.SandboxNetworkProxyRoutingParam{{
								Body: map[string]string{
									"0":  "{",
									"1":  "\"",
									"2":  "a",
									"3":  "p",
									"4":  "i",
									"5":  "_",
									"6":  "k",
									"7":  "e",
									"8":  "y",
									"9":  "\"",
									"10": ":",
									"11": " ",
									"12": "\"",
									"13": "{",
									"14": "{",
									"15": "S",
									"16": "E",
									"17": "C",
									"18": "R",
									"19": "E",
									"20": "T",
									"21": ":",
									"22": "s",
									"23": "t",
									"24": "r",
									"25": "i",
									"26": "p",
									"27": "e",
									"28": "-",
									"29": "k",
									"30": "e",
									"31": "y",
									"32": "}",
									"33": "}",
									"34": "\"",
									"35": "}",
								},
								Destinations: []string{"string"},
								Headers: map[string]string{
									"0":  "{",
									"1":  "\"",
									"2":  "A",
									"3":  "u",
									"4":  "t",
									"5":  "h",
									"6":  "o",
									"7":  "r",
									"8":  "i",
									"9":  "z",
									"10": "a",
									"11": "t",
									"12": "i",
									"13": "o",
									"14": "n",
									"15": "\"",
									"16": ":",
									"17": " ",
									"18": "\"",
									"19": "B",
									"20": "e",
									"21": "a",
									"22": "r",
									"23": "e",
									"24": "r",
									"25": " ",
									"26": "{",
									"27": "{",
									"28": "S",
									"29": "E",
									"30": "C",
									"31": "R",
									"32": "E",
									"33": "T",
									"34": ":",
									"35": "s",
									"36": "t",
									"37": "r",
									"38": "i",
									"39": "p",
									"40": "e",
									"41": "-",
									"42": "k",
									"43": "e",
									"44": "y",
									"45": "}",
									"46": "}",
									"47": "\"",
									"48": "}",
								},
								Secrets: map[string]string{
									"0":  "{",
									"1":  "\"",
									"2":  "s",
									"3":  "t",
									"4":  "r",
									"5":  "i",
									"6":  "p",
									"7":  "e",
									"8":  "-",
									"9":  "k",
									"10": "e",
									"11": "y",
									"12": "\"",
									"13": ":",
									"14": " ",
									"15": "\"",
									"16": "s",
									"17": "k",
									"18": "-",
									"19": "l",
									"20": "i",
									"21": "v",
									"22": "e",
									"23": "-",
									"24": "a",
									"25": "b",
									"26": "c",
									"27": "1",
									"28": "2",
									"29": "3",
									"30": ".",
									"31": ".",
									"32": ".",
									"33": "\"",
									"34": "}",
								},
							}},
						},
					},
					Region: blaxel.String("us-pdx-1"),
					Runtime: blaxel.SandboxRuntimeParam{
						Envs: []shared.EnvParam{{
							Name:   blaxel.String("MY_ENV_VAR"),
							Secret: blaxel.Bool(true),
							Value:  blaxel.String("my-value"),
						}},
						Expires: blaxel.String("2025-12-31T23:59:59Z"),
						Image:   blaxel.String("blaxel/base-image:latest"),
						Memory:  blaxel.Int(4096),
						Ports: []blaxel.PortParam{{
							Target:   8080,
							Name:     blaxel.String("http"),
							Protocol: blaxel.PortProtocolHTTP,
						}},
						TerminationGracePeriodSeconds: blaxel.Int(30),
						Ttl:                           blaxel.String("24h"),
					},
					Volumes: []blaxel.VolumeAttachmentParam{{
						MountPath: blaxel.String("/mnt/data"),
						Name:      blaxel.String("my-volume"),
						ReadOnly:  blaxel.Bool(false),
					}},
				},
			},
		},
	)
	if err != nil {
		var apierr *blaxel.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSandboxList(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := blaxel.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Sandboxes.List(context.TODO())
	if err != nil {
		var apierr *blaxel.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSandboxDelete(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := blaxel.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Sandboxes.Delete(context.TODO(), "sandboxName")
	if err != nil {
		var apierr *blaxel.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSandboxGetHub(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := blaxel.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Sandboxes.GetHub(context.TODO())
	if err != nil {
		var apierr *blaxel.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
