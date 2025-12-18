# Agents

Params Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#AgentParam">AgentParam</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#CoreEventParam">CoreEventParam</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#MetadataParam">MetadataParam</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#RevisionConfigurationParam">RevisionConfigurationParam</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#Status">Status</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#TriggerParam">TriggerParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#Agent">Agent</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#CoreEvent">CoreEvent</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#Metadata">Metadata</a>
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

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#CustomDomain">CustomDomain</a>

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

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#Function">Function</a>

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

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#IntegrationConnection">IntegrationConnection</a>

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

# Jobs

Params Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#JobParam">JobParam</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#ModelParam">ModelParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#Job">Job</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#Model">Model</a>

Methods:

- <code title="post /jobs">client.Jobs.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#JobService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#JobNewParams">JobNewParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#Job">Job</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /jobs/{jobId}">client.Jobs.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#JobService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, jobID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#JobGetParams">JobGetParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#Job">Job</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /jobs/{jobId}">client.Jobs.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#JobService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, jobID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#JobUpdateParams">JobUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#Job">Job</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /jobs">client.Jobs.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#JobService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) ([]<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#Job">Job</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /jobs/{jobId}">client.Jobs.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#JobService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, jobID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#Job">Job</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /jobs/{jobId}/revisions">client.Jobs.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#JobService.ListRevisions">ListRevisions</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, jobID <a href="https://pkg.go.dev/builtin#string">string</a>) ([]<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#RevisionMetadata">RevisionMetadata</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

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

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#Policy">Policy</a>

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

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#VolumeTemplate">VolumeTemplate</a>
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

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#Workspace">Workspace</a>
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

- <a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#PortParam">PortParam</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#SandboxParam">SandboxParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#Port">Port</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#Sandbox">Sandbox</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#SandboxGetHubResponse">SandboxGetHubResponse</a>

Methods:

- <code title="post /sandboxes">client.Sandboxes.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#SandboxService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#SandboxNewParams">SandboxNewParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#Sandbox">Sandbox</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /sandboxes/{sandboxName}">client.Sandboxes.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#SandboxService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, sandboxName <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#SandboxGetParams">SandboxGetParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#Sandbox">Sandbox</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /sandboxes/{sandboxName}">client.Sandboxes.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#SandboxService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, sandboxName <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#SandboxUpdateParams">SandboxUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#Sandbox">Sandbox</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /sandboxes">client.Sandboxes.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#SandboxService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) ([]<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#Sandbox">Sandbox</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /sandboxes/{sandboxName}">client.Sandboxes.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#SandboxService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, sandboxName <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#Sandbox">Sandbox</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /sandbox/hub">client.Sandboxes.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#SandboxService.GetHub">GetHub</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) ([]<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#SandboxGetHubResponse">SandboxGetHubResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Previews

Params Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#PreviewParam">PreviewParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/blaxel-go#Preview">Preview</a>

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
