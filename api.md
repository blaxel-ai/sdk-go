# Agents

Params Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#AgentParam">AgentParam</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#AgentRuntimeParam">AgentRuntimeParam</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#AgentSpecParam">AgentSpecParam</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#CoreEventParam">CoreEventParam</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#MetadataParam">MetadataParam</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#RepositoryParam">RepositoryParam</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#RevisionConfigurationParam">RevisionConfigurationParam</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#Status">Status</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#TriggerParam">TriggerParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#Agent">Agent</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#AgentRuntime">AgentRuntime</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#AgentSpec">AgentSpec</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#CoreEvent">CoreEvent</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#Metadata">Metadata</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#Repository">Repository</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#RevisionConfiguration">RevisionConfiguration</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#RevisionMetadata">RevisionMetadata</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#Status">Status</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#Trigger">Trigger</a>

Methods:

- <code title="post /agents">client.Agents.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#AgentService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#AgentNewParams">AgentNewParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#Agent">Agent</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /agents/{agentName}">client.Agents.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#AgentService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, agentName <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#AgentGetParams">AgentGetParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#Agent">Agent</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /agents/{agentName}">client.Agents.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#AgentService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, agentName <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#AgentUpdateParams">AgentUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#Agent">Agent</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /agents">client.Agents.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#AgentService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) ([]<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#Agent">Agent</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /agents/{agentName}">client.Agents.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#AgentService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, agentName <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#Agent">Agent</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /agents/{agentName}/revisions">client.Agents.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#AgentService.ListRevisions">ListRevisions</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, agentName <a href="https://pkg.go.dev/builtin#string">string</a>) ([]<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#RevisionMetadata">RevisionMetadata</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Configuration

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#ConfigurationListResponse">ConfigurationListResponse</a>

Methods:

- <code title="get /configuration">client.Configuration.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#ConfigurationService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#ConfigurationListResponse">ConfigurationListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Customdomains

Params Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#CustomDomainParam">CustomDomainParam</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#CustomDomainMetadataParam">CustomDomainMetadataParam</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#CustomDomainSpecParam">CustomDomainSpecParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#CustomDomain">CustomDomain</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#CustomDomainMetadata">CustomDomainMetadata</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#CustomDomainSpec">CustomDomainSpec</a>

Methods:

- <code title="post /customdomains">client.Customdomains.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#CustomdomainService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#CustomdomainNewParams">CustomdomainNewParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#CustomDomain">CustomDomain</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /customdomains/{domainName}">client.Customdomains.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#CustomdomainService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, domainName <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#CustomDomain">CustomDomain</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /customdomains/{domainName}">client.Customdomains.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#CustomdomainService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, domainName <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#CustomdomainUpdateParams">CustomdomainUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#CustomDomain">CustomDomain</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /customdomains">client.Customdomains.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#CustomdomainService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) ([]<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#CustomDomain">CustomDomain</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /customdomains/{domainName}">client.Customdomains.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#CustomdomainService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, domainName <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#CustomDomain">CustomDomain</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /customdomains/{domainName}/verify">client.Customdomains.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#CustomdomainService.Verify">Verify</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, domainName <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#CustomDomain">CustomDomain</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Functions

Params Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#FunctionParam">FunctionParam</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#FunctionRuntimeParam">FunctionRuntimeParam</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#FunctionSpecParam">FunctionSpecParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#Function">Function</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#FunctionRuntime">FunctionRuntime</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#FunctionSpec">FunctionSpec</a>

Methods:

- <code title="post /functions">client.Functions.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#FunctionService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#FunctionNewParams">FunctionNewParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#Function">Function</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /functions/{functionName}">client.Functions.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#FunctionService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, functionName <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#FunctionGetParams">FunctionGetParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#Function">Function</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /functions/{functionName}">client.Functions.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#FunctionService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, functionName <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#FunctionUpdateParams">FunctionUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#Function">Function</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /functions">client.Functions.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#FunctionService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) ([]<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#Function">Function</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /functions/{functionName}">client.Functions.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#FunctionService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, functionName <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#Function">Function</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /functions/{functionName}/revisions">client.Functions.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#FunctionService.ListRevisions">ListRevisions</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, functionName <a href="https://pkg.go.dev/builtin#string">string</a>) ([]<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#RevisionMetadata">RevisionMetadata</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Integrations

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#IntegrationGetResponse">IntegrationGetResponse</a>

Methods:

- <code title="get /integrations/{integrationName}">client.Integrations.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#IntegrationService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, integrationName <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#IntegrationGetResponse">IntegrationGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Connections

Params Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#IntegrationConnectionParam">IntegrationConnectionParam</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#IntegrationConnectionSpecParam">IntegrationConnectionSpecParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#IntegrationConnection">IntegrationConnection</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#IntegrationConnectionSpec">IntegrationConnectionSpec</a>

Methods:

- <code title="post /integrations/connections">client.Integrations.Connections.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#IntegrationConnectionService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#IntegrationConnectionNewParams">IntegrationConnectionNewParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#IntegrationConnection">IntegrationConnection</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /integrations/connections/{connectionName}">client.Integrations.Connections.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#IntegrationConnectionService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, connectionName <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#IntegrationConnection">IntegrationConnection</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /integrations/connections/{connectionName}">client.Integrations.Connections.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#IntegrationConnectionService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, connectionName <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#IntegrationConnectionUpdateParams">IntegrationConnectionUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#IntegrationConnection">IntegrationConnection</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /integrations/connections">client.Integrations.Connections.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#IntegrationConnectionService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) ([]<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#IntegrationConnection">IntegrationConnection</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /integrations/connections/{connectionName}">client.Integrations.Connections.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#IntegrationConnectionService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, connectionName <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#IntegrationConnection">IntegrationConnection</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /integrations/connections/{connectionName}/endpointConfigurations">client.Integrations.Connections.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#IntegrationConnectionService.ListEndpointConfigurations">ListEndpointConfigurations</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, connectionName <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

### Models

Methods:

- <code title="get /integrations/connections/{connectionName}/models/{modelId}">client.Integrations.Connections.Models.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#IntegrationConnectionModelService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, modelID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#IntegrationConnectionModelGetParams">IntegrationConnectionModelGetParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="get /integrations/connections/{connectionName}/models">client.Integrations.Connections.Models.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#IntegrationConnectionModelService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, connectionName <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

# Images

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#Image">Image</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#ImageCleanupResponse">ImageCleanupResponse</a>

Methods:

- <code title="get /images/{resourceType}/{imageName}">client.Images.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#ImageService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, imageName <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#ImageGetParams">ImageGetParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#Image">Image</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /images">client.Images.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#ImageService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) ([]<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#Image">Image</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /images/{resourceType}/{imageName}">client.Images.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#ImageService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, imageName <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#ImageDeleteParams">ImageDeleteParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#Image">Image</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /images">client.Images.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#ImageService.Cleanup">Cleanup</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#ImageCleanupResponse">ImageCleanupResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Tags

Methods:

- <code title="delete /images/{resourceType}/{imageName}/tags/{tagName}">client.Images.Tags.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#ImageTagService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, tagName <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#ImageTagDeleteParams">ImageTagDeleteParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#Image">Image</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Jobs

Params Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#CreateJobExecutionRequestParam">CreateJobExecutionRequestParam</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#JobParam">JobParam</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#JobRuntimeParam">JobRuntimeParam</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#JobSpecParam">JobSpecParam</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#ModelParam">ModelParam</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#ModelRuntimeParam">ModelRuntimeParam</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#ModelSpecParam">ModelSpecParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#Job">Job</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#JobExecution">JobExecution</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#JobRuntime">JobRuntime</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#JobSpec">JobSpec</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#Model">Model</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#ModelRuntime">ModelRuntime</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#ModelSpec">ModelSpec</a>

Methods:

- <code title="post /jobs">client.Jobs.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#JobService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#JobNewParams">JobNewParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#Job">Job</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /jobs/{jobId}">client.Jobs.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#JobService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, jobID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#JobGetParams">JobGetParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#Job">Job</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /jobs/{jobId}">client.Jobs.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#JobService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, jobID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#JobUpdateParams">JobUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#Job">Job</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /jobs">client.Jobs.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#JobService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) ([]<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#Job">Job</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /jobs/{jobId}">client.Jobs.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#JobService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, jobID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#Job">Job</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /jobs/{jobId}/revisions">client.Jobs.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#JobService.ListRevisions">ListRevisions</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, jobID <a href="https://pkg.go.dev/builtin#string">string</a>) ([]<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#RevisionMetadata">RevisionMetadata</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Executions

Methods:

- <code title="post /jobs/{jobId}/executions">client.Jobs.Executions.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#JobExecutionService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, jobID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#JobExecutionNewParams">JobExecutionNewParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#JobExecution">JobExecution</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /jobs/{jobId}/executions/{executionId}">client.Jobs.Executions.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#JobExecutionService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, executionID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#JobExecutionGetParams">JobExecutionGetParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#JobExecution">JobExecution</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /jobs/{jobId}/executions">client.Jobs.Executions.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#JobExecutionService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, jobID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#JobExecutionListParams">JobExecutionListParams</a>) ([]<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#JobExecution">JobExecution</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /jobs/{jobId}/executions/{executionId}">client.Jobs.Executions.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#JobExecutionService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, executionID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#JobExecutionDeleteParams">JobExecutionDeleteParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#JobExecution">JobExecution</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Locations

Params Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#FlavorParam">FlavorParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#Flavor">Flavor</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#LocationListResponse">LocationListResponse</a>

Methods:

- <code title="get /locations">client.Locations.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#LocationService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) ([]<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#LocationListResponse">LocationListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Mcp

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#McpGetHubResponse">McpGetHubResponse</a>

Methods:

- <code title="get /mcp/hub">client.Mcp.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#McpService.GetHub">GetHub</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) ([]<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#McpGetHubResponse">McpGetHubResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Models

Methods:

- <code title="post /models">client.Models.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#ModelService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#ModelNewParams">ModelNewParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#Model">Model</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /models/{modelName}">client.Models.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#ModelService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, modelName <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#Model">Model</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /models/{modelName}">client.Models.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#ModelService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, modelName <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#ModelUpdateParams">ModelUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#Model">Model</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /models">client.Models.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#ModelService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) ([]<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#Model">Model</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /models/{modelName}">client.Models.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#ModelService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, modelName <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#Model">Model</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /models/{modelName}/revisions">client.Models.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#ModelService.ListRevisions">ListRevisions</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, modelName <a href="https://pkg.go.dev/builtin#string">string</a>) ([]<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#RevisionMetadata">RevisionMetadata</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Policies

Params Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#PolicyParam">PolicyParam</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#PolicyLocationParam">PolicyLocationParam</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#PolicyMaxTokensParam">PolicyMaxTokensParam</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#PolicySpecParam">PolicySpecParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#Policy">Policy</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#PolicyLocation">PolicyLocation</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#PolicyMaxTokens">PolicyMaxTokens</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#PolicySpec">PolicySpec</a>

Methods:

- <code title="post /policies">client.Policies.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#PolicyService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#PolicyNewParams">PolicyNewParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#Policy">Policy</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /policies/{policyName}">client.Policies.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#PolicyService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, policyName <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#Policy">Policy</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /policies/{policyName}">client.Policies.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#PolicyService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, policyName <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#PolicyUpdateParams">PolicyUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#Policy">Policy</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /policies">client.Policies.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#PolicyService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) ([]<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#Policy">Policy</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /policies/{policyName}">client.Policies.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#PolicyService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, policyName <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#Policy">Policy</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Profile

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#ProfileListInvitationsResponse">ProfileListInvitationsResponse</a>

Methods:

- <code title="get /profile/invitations">client.Profile.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#ProfileService.ListInvitations">ListInvitations</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) ([]<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#ProfileListInvitationsResponse">ProfileListInvitationsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# PublicIPs

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#PublicIPListResponse">PublicIPListResponse</a>

Methods:

- <code title="get /publicIps">client.PublicIPs.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#PublicIPService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#PublicIPListParams">PublicIPListParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#PublicIPListResponse">PublicIPListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# ServiceAccounts

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#ServiceAccountNewResponse">ServiceAccountNewResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#ServiceAccountUpdateResponse">ServiceAccountUpdateResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#ServiceAccountListResponse">ServiceAccountListResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#ServiceAccountDeleteResponse">ServiceAccountDeleteResponse</a>

Methods:

- <code title="post /service_accounts">client.ServiceAccounts.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#ServiceAccountService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#ServiceAccountNewParams">ServiceAccountNewParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#ServiceAccountNewResponse">ServiceAccountNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /service_accounts/{clientId}">client.ServiceAccounts.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#ServiceAccountService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, clientID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#ServiceAccountUpdateParams">ServiceAccountUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#ServiceAccountUpdateResponse">ServiceAccountUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /service_accounts">client.ServiceAccounts.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#ServiceAccountService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) ([]<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#ServiceAccountListResponse">ServiceAccountListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /service_accounts/{clientId}">client.ServiceAccounts.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#ServiceAccountService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, clientID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#ServiceAccountDeleteResponse">ServiceAccountDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## APIKeys

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#APIKey">APIKey</a>

Methods:

- <code title="post /service_accounts/{clientId}/api_keys">client.ServiceAccounts.APIKeys.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#ServiceAccountAPIKeyService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, clientID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#ServiceAccountAPIKeyNewParams">ServiceAccountAPIKeyNewParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#APIKey">APIKey</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /service_accounts/{clientId}/api_keys">client.ServiceAccounts.APIKeys.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#ServiceAccountAPIKeyService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, clientID <a href="https://pkg.go.dev/builtin#string">string</a>) ([]<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#APIKey">APIKey</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /service_accounts/{clientId}/api_keys/{apiKeyId}">client.ServiceAccounts.APIKeys.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#ServiceAccountAPIKeyService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, apiKeyID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#ServiceAccountAPIKeyDeleteParams">ServiceAccountAPIKeyDeleteParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>

# Templates

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#Template">Template</a>

Methods:

- <code title="get /templates/{templateName}">client.Templates.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#TemplateService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, templateName <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#Template">Template</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /templates">client.Templates.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#TemplateService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) ([]<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#Template">Template</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Users

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#PendingInvitation">PendingInvitation</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#WorkspaceUser">WorkspaceUser</a>

Methods:

- <code title="get /users">client.Users.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#UserService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) ([]<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#WorkspaceUser">WorkspaceUser</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /users">client.Users.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#UserService.Invite">Invite</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#UserInviteParams">UserInviteParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#PendingInvitation">PendingInvitation</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /users/{subOrEmail}">client.Users.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#UserService.Remove">Remove</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, subOrEmail <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="put /users/{subOrEmail}">client.Users.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#UserService.UpdateRole">UpdateRole</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, subOrEmail <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#UserUpdateRoleParams">UserUpdateRoleParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#WorkspaceUser">WorkspaceUser</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# VolumeTemplates

Params Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#VolumeTemplateParam">VolumeTemplateParam</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#VolumeTemplateSpecParam">VolumeTemplateSpecParam</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#VolumeTemplateStateParam">VolumeTemplateStateParam</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#VolumeTemplateVersionParam">VolumeTemplateVersionParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#VolumeTemplate">VolumeTemplate</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#VolumeTemplateSpec">VolumeTemplateSpec</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#VolumeTemplateState">VolumeTemplateState</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#VolumeTemplateVersion">VolumeTemplateVersion</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#VolumeTemplateDeleteVersionResponse">VolumeTemplateDeleteVersionResponse</a>

Methods:

- <code title="post /volume_templates">client.VolumeTemplates.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#VolumeTemplateService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#VolumeTemplateNewParams">VolumeTemplateNewParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#VolumeTemplate">VolumeTemplate</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /volume_templates/{volumeTemplateName}">client.VolumeTemplates.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#VolumeTemplateService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, volumeTemplateName <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#VolumeTemplate">VolumeTemplate</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /volume_templates">client.VolumeTemplates.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#VolumeTemplateService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) ([]<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#VolumeTemplate">VolumeTemplate</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /volume_templates/{volumeTemplateName}">client.VolumeTemplates.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#VolumeTemplateService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, volumeTemplateName <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#VolumeTemplate">VolumeTemplate</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /volume_templates/{volumeTemplateName}/versions/{versionName}">client.VolumeTemplates.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#VolumeTemplateService.DeleteVersion">DeleteVersion</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, versionName <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#VolumeTemplateDeleteVersionParams">VolumeTemplateDeleteVersionParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#VolumeTemplateDeleteVersionResponse">VolumeTemplateDeleteVersionResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /volume_templates/{volumeTemplateName}">client.VolumeTemplates.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#VolumeTemplateService.Upsert">Upsert</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, volumeTemplateName <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#VolumeTemplateUpsertParams">VolumeTemplateUpsertParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#VolumeTemplate">VolumeTemplate</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Volumes

Params Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#VolumeParam">VolumeParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#Volume">Volume</a>

Methods:

- <code title="post /volumes">client.Volumes.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#VolumeService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#VolumeNewParams">VolumeNewParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#Volume">Volume</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /volumes/{volumeName}">client.Volumes.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#VolumeService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, volumeName <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#Volume">Volume</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /volumes">client.Volumes.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#VolumeService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) ([]<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#Volume">Volume</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /volumes/{volumeName}">client.Volumes.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#VolumeService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, volumeName <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#Volume">Volume</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Workspaces

Params Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#WorkspaceParam">WorkspaceParam</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#WorkspaceRuntimeParam">WorkspaceRuntimeParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#Workspace">Workspace</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#WorkspaceRuntime">WorkspaceRuntime</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#WorkspaceAcceptInvitationResponse">WorkspaceAcceptInvitationResponse</a>

Methods:

- <code title="post /workspaces">client.Workspaces.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#WorkspaceService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#WorkspaceNewParams">WorkspaceNewParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#Workspace">Workspace</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /workspaces/{workspaceName}">client.Workspaces.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#WorkspaceService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, workspaceName <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#Workspace">Workspace</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /workspaces/{workspaceName}">client.Workspaces.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#WorkspaceService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, workspaceName <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#WorkspaceUpdateParams">WorkspaceUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#Workspace">Workspace</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /workspaces">client.Workspaces.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#WorkspaceService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) ([]<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#Workspace">Workspace</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /workspaces/{workspaceName}">client.Workspaces.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#WorkspaceService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, workspaceName <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#Workspace">Workspace</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /workspaces/{workspaceName}/join">client.Workspaces.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#WorkspaceService.AcceptInvitation">AcceptInvitation</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, workspaceName <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#WorkspaceAcceptInvitationResponse">WorkspaceAcceptInvitationResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /workspaces/availability">client.Workspaces.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#WorkspaceService.CheckAvailability">CheckAvailability</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#WorkspaceCheckAvailabilityParams">WorkspaceCheckAvailabilityParams</a>) (<a href="https://pkg.go.dev/builtin#bool">bool</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /workspaces/{workspaceName}/decline">client.Workspaces.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#WorkspaceService.DeclineInvitation">DeclineInvitation</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, workspaceName <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#PendingInvitation">PendingInvitation</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /workspaces/{workspaceName}/leave">client.Workspaces.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#WorkspaceService.Leave">Leave</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, workspaceName <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#Workspace">Workspace</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Sandboxes

Params Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#ExpirationPolicyParam">ExpirationPolicyParam</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#PortParam">PortParam</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#SandboxParam">SandboxParam</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#SandboxLifecycleParam">SandboxLifecycleParam</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#SandboxRuntimeParam">SandboxRuntimeParam</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#SandboxSpecParam">SandboxSpecParam</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#VolumeAttachmentParam">VolumeAttachmentParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#ExpirationPolicy">ExpirationPolicy</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#Port">Port</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#Sandbox">Sandbox</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#SandboxLifecycle">SandboxLifecycle</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#SandboxRuntime">SandboxRuntime</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#SandboxSpec">SandboxSpec</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#VolumeAttachment">VolumeAttachment</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#SandboxGetHubResponse">SandboxGetHubResponse</a>

Methods:

- <code title="post /sandboxes">client.Sandboxes.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#SandboxService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#SandboxNewParams">SandboxNewParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#Sandbox">Sandbox</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /sandboxes/{sandboxName}">client.Sandboxes.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#SandboxService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, sandboxName <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#SandboxGetParams">SandboxGetParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#Sandbox">Sandbox</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /sandboxes/{sandboxName}">client.Sandboxes.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#SandboxService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, sandboxName <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#SandboxUpdateParams">SandboxUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#Sandbox">Sandbox</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /sandboxes">client.Sandboxes.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#SandboxService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) ([]<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#Sandbox">Sandbox</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /sandboxes/{sandboxName}">client.Sandboxes.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#SandboxService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, sandboxName <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#Sandbox">Sandbox</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /sandbox/hub">client.Sandboxes.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#SandboxService.GetHub">GetHub</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) ([]<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#SandboxGetHubResponse">SandboxGetHubResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Process

Params Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#ProcessRequestParam">ProcessRequestParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#ProcessLogs">ProcessLogs</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#ProcessResponse">ProcessResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#SandboxProcessKillResponse">SandboxProcessKillResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#SandboxProcessStopResponse">SandboxProcessStopResponse</a>

Methods:

- <code title="post /process">client.Sandboxes.Process.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#SandboxProcessService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#SandboxProcessNewParams">SandboxProcessNewParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#ProcessResponse">ProcessResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /process/{identifier}">client.Sandboxes.Process.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#SandboxProcessService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, identifier <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#ProcessResponse">ProcessResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /process">client.Sandboxes.Process.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#SandboxProcessService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) ([]<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#ProcessResponse">ProcessResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /process/{identifier}/kill">client.Sandboxes.Process.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#SandboxProcessService.Kill">Kill</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, identifier <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#SandboxProcessKillResponse">SandboxProcessKillResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /process/{identifier}/logs">client.Sandboxes.Process.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#SandboxProcessService.GetLogs">GetLogs</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, identifier <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#ProcessLogs">ProcessLogs</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /process/{identifier}">client.Sandboxes.Process.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#SandboxProcessService.Stop">Stop</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, identifier <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#SandboxProcessStopResponse">SandboxProcessStopResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Fs

Params Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#FileRequestParam">FileRequestParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#Directory">Directory</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#File">File</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#Subdirectory">Subdirectory</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#SandboxFGetResponseUnion">SandboxFGetResponseUnion</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#SandboxFUpdateResponse">SandboxFUpdateResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#SandboxFListResponseUnion">SandboxFListResponseUnion</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#SandboxFDeleteResponse">SandboxFDeleteResponse</a>

Methods:

- <code title="get /filesystem/{path}">client.Sandboxes.Fs.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#SandboxFService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, path <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#SandboxFGetParams">SandboxFGetParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#SandboxFGetResponseUnion">SandboxFGetResponseUnion</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /filesystem/{path}">client.Sandboxes.Fs.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#SandboxFService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, path <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#SandboxFUpdateParams">SandboxFUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#SandboxFUpdateResponse">SandboxFUpdateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /filesystem/{path}">client.Sandboxes.Fs.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#SandboxFService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, path <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#SandboxFListParams">SandboxFListParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#SandboxFListResponseUnion">SandboxFListResponseUnion</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /filesystem/{path}">client.Sandboxes.Fs.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#SandboxFService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, path <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#SandboxFDeleteParams">SandboxFDeleteParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#SandboxFDeleteResponse">SandboxFDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Codegen

Params Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#ApplyEditRequestParam">ApplyEditRequestParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#ApplyEditResponse">ApplyEditResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#RerankingResponse">RerankingResponse</a>

Methods:

- <code title="put /codegen/fastapply/{path}">client.Sandboxes.Codegen.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#SandboxCodegenService.Fastapply">Fastapply</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, path <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#SandboxCodegenFastapplyParams">SandboxCodegenFastapplyParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#ApplyEditResponse">ApplyEditResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /codegen/reranking/{path}">client.Sandboxes.Codegen.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#SandboxCodegenService.Reranking">Reranking</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, path <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#SandboxCodegenRerankingParams">SandboxCodegenRerankingParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#RerankingResponse">RerankingResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Previews

Params Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#PreviewParam">PreviewParam</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#PreviewMetadataParam">PreviewMetadataParam</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#PreviewSpecParam">PreviewSpecParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#Preview">Preview</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#PreviewMetadata">PreviewMetadata</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#PreviewSpec">PreviewSpec</a>

Methods:

- <code title="post /sandboxes/{sandboxName}/previews">client.Sandboxes.Previews.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#SandboxPreviewService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, sandboxName <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#SandboxPreviewNewParams">SandboxPreviewNewParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#Preview">Preview</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /sandboxes/{sandboxName}/previews/{previewName}">client.Sandboxes.Previews.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#SandboxPreviewService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, previewName <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#SandboxPreviewGetParams">SandboxPreviewGetParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#Preview">Preview</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /sandboxes/{sandboxName}/previews/{previewName}">client.Sandboxes.Previews.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#SandboxPreviewService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, previewName <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#SandboxPreviewUpdateParams">SandboxPreviewUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#Preview">Preview</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /sandboxes/{sandboxName}/previews">client.Sandboxes.Previews.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#SandboxPreviewService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, sandboxName <a href="https://pkg.go.dev/builtin#string">string</a>) ([]<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#Preview">Preview</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /sandboxes/{sandboxName}/previews/{previewName}">client.Sandboxes.Previews.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#SandboxPreviewService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, previewName <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#SandboxPreviewDeleteParams">SandboxPreviewDeleteParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#Preview">Preview</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Tokens

Params Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#PreviewTokenParam">PreviewTokenParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#PreviewToken">PreviewToken</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#SandboxPreviewTokenDeleteResponse">SandboxPreviewTokenDeleteResponse</a>

Methods:

- <code title="post /sandboxes/{sandboxName}/previews/{previewName}/tokens">client.Sandboxes.Previews.Tokens.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#SandboxPreviewTokenService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, previewName <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#SandboxPreviewTokenNewParams">SandboxPreviewTokenNewParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#PreviewToken">PreviewToken</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /sandboxes/{sandboxName}/previews/{previewName}/tokens">client.Sandboxes.Previews.Tokens.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#SandboxPreviewTokenService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, previewName <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#SandboxPreviewTokenGetParams">SandboxPreviewTokenGetParams</a>) ([]<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#PreviewToken">PreviewToken</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /sandboxes/{sandboxName}/previews/{previewName}/tokens/{tokenName}">client.Sandboxes.Previews.Tokens.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#SandboxPreviewTokenService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, tokenName <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#SandboxPreviewTokenDeleteParams">SandboxPreviewTokenDeleteParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#SandboxPreviewTokenDeleteResponse">SandboxPreviewTokenDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
