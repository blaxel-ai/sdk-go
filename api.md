# Agents

Params Types:

- <a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#AgentParam">AgentParam</a>
- <a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#AgentRuntimeParam">AgentRuntimeParam</a>
- <a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#AgentSpecParam">AgentSpecParam</a>
- <a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#CoreEventParam">CoreEventParam</a>
- <a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#MetadataParam">MetadataParam</a>
- <a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#RepositoryParam">RepositoryParam</a>
- <a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#RevisionConfigurationParam">RevisionConfigurationParam</a>
- <a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#Status">Status</a>
- <a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#TriggerParam">TriggerParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#Agent">Agent</a>
- <a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#AgentRuntime">AgentRuntime</a>
- <a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#AgentSpec">AgentSpec</a>
- <a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#CoreEvent">CoreEvent</a>
- <a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#Metadata">Metadata</a>
- <a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#Repository">Repository</a>
- <a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#RevisionConfiguration">RevisionConfiguration</a>
- <a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#RevisionMetadata">RevisionMetadata</a>
- <a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#Status">Status</a>
- <a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#Trigger">Trigger</a>

Methods:

- <code title="post /agents">client.Agents.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#AgentService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#AgentNewParams">AgentNewParams</a>) (<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#Agent">Agent</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /agents/{agentName}">client.Agents.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#AgentService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, agentName <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#AgentGetParams">AgentGetParams</a>) (<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#Agent">Agent</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /agents/{agentName}">client.Agents.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#AgentService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, agentName <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#AgentUpdateParams">AgentUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#Agent">Agent</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /agents">client.Agents.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#AgentService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) ([]<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#Agent">Agent</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /agents/{agentName}">client.Agents.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#AgentService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, agentName <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#Agent">Agent</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /agents/{agentName}/revisions">client.Agents.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#AgentService.ListRevisions">ListRevisions</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, agentName <a href="https://pkg.go.dev/builtin#string">string</a>) ([]<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#RevisionMetadata">RevisionMetadata</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Functions

Params Types:

- <a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#FunctionParam">FunctionParam</a>
- <a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#FunctionRuntimeParam">FunctionRuntimeParam</a>
- <a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#FunctionSpecParam">FunctionSpecParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#Function">Function</a>
- <a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#FunctionRuntime">FunctionRuntime</a>
- <a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#FunctionSpec">FunctionSpec</a>

Methods:

- <code title="post /functions">client.Functions.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#FunctionService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#FunctionNewParams">FunctionNewParams</a>) (<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#Function">Function</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /functions/{functionName}">client.Functions.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#FunctionService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, functionName <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#FunctionGetParams">FunctionGetParams</a>) (<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#Function">Function</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /functions/{functionName}">client.Functions.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#FunctionService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, functionName <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#FunctionUpdateParams">FunctionUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#Function">Function</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /functions">client.Functions.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#FunctionService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) ([]<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#Function">Function</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /functions/{functionName}">client.Functions.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#FunctionService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, functionName <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#Function">Function</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /functions/{functionName}/revisions">client.Functions.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#FunctionService.ListRevisions">ListRevisions</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, functionName <a href="https://pkg.go.dev/builtin#string">string</a>) ([]<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#RevisionMetadata">RevisionMetadata</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Integrations

Response Types:

- <a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#IntegrationGetResponse">IntegrationGetResponse</a>

Methods:

- <code title="get /integrations/{integrationName}">client.Integrations.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#IntegrationService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, integrationName <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#IntegrationGetResponse">IntegrationGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Connections

Params Types:

- <a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#IntegrationConnectionParam">IntegrationConnectionParam</a>
- <a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#IntegrationConnectionSpecParam">IntegrationConnectionSpecParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#IntegrationConnection">IntegrationConnection</a>
- <a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#IntegrationConnectionSpec">IntegrationConnectionSpec</a>

Methods:

- <code title="post /integrations/connections">client.Integrations.Connections.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#IntegrationConnectionService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#IntegrationConnectionNewParams">IntegrationConnectionNewParams</a>) (<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#IntegrationConnection">IntegrationConnection</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /integrations/connections/{connectionName}">client.Integrations.Connections.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#IntegrationConnectionService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, connectionName <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#IntegrationConnection">IntegrationConnection</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /integrations/connections/{connectionName}">client.Integrations.Connections.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#IntegrationConnectionService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, connectionName <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#IntegrationConnectionUpdateParams">IntegrationConnectionUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#IntegrationConnection">IntegrationConnection</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /integrations/connections">client.Integrations.Connections.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#IntegrationConnectionService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) ([]<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#IntegrationConnection">IntegrationConnection</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /integrations/connections/{connectionName}">client.Integrations.Connections.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#IntegrationConnectionService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, connectionName <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#IntegrationConnection">IntegrationConnection</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Images

Response Types:

- <a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#Image">Image</a>
- <a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#ImageCleanupResponse">ImageCleanupResponse</a>

Methods:

- <code title="get /images/{resourceType}/{imageName}">client.Images.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#ImageService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, imageName <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#ImageGetParams">ImageGetParams</a>) (<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#Image">Image</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /images">client.Images.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#ImageService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) ([]<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#Image">Image</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /images/{resourceType}/{imageName}">client.Images.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#ImageService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, imageName <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#ImageDeleteParams">ImageDeleteParams</a>) (<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#Image">Image</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /images">client.Images.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#ImageService.Cleanup">Cleanup</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#ImageCleanupResponse">ImageCleanupResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Tags

Methods:

- <code title="delete /images/{resourceType}/{imageName}/tags/{tagName}">client.Images.Tags.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#ImageTagService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, tagName <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#ImageTagDeleteParams">ImageTagDeleteParams</a>) (<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#Image">Image</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Jobs

Params Types:

- <a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#CreateJobExecutionRequestParam">CreateJobExecutionRequestParam</a>
- <a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#JobParam">JobParam</a>
- <a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#JobRuntimeParam">JobRuntimeParam</a>
- <a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#JobSpecParam">JobSpecParam</a>
- <a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#ModelParam">ModelParam</a>
- <a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#ModelRuntimeParam">ModelRuntimeParam</a>
- <a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#ModelSpecParam">ModelSpecParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#Job">Job</a>
- <a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#JobExecution">JobExecution</a>
- <a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#JobRuntime">JobRuntime</a>
- <a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#JobSpec">JobSpec</a>
- <a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#Model">Model</a>
- <a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#ModelRuntime">ModelRuntime</a>
- <a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#ModelSpec">ModelSpec</a>

Methods:

- <code title="post /jobs">client.Jobs.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#JobService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#JobNewParams">JobNewParams</a>) (<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#Job">Job</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /jobs/{jobId}">client.Jobs.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#JobService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, jobID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#JobGetParams">JobGetParams</a>) (<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#Job">Job</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /jobs/{jobId}">client.Jobs.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#JobService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, jobID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#JobUpdateParams">JobUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#Job">Job</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /jobs">client.Jobs.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#JobService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) ([]<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#Job">Job</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /jobs/{jobId}">client.Jobs.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#JobService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, jobID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#Job">Job</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /jobs/{jobId}/revisions">client.Jobs.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#JobService.ListRevisions">ListRevisions</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, jobID <a href="https://pkg.go.dev/builtin#string">string</a>) ([]<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#RevisionMetadata">RevisionMetadata</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Executions

Response Types:

- <a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#JobExecutionNewResponse">JobExecutionNewResponse</a>

Methods:

- <code title="post /jobs/{jobId}/executions">client.Jobs.Executions.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#JobExecutionService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, jobID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#JobExecutionNewParams">JobExecutionNewParams</a>) (<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#JobExecutionNewResponse">JobExecutionNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /jobs/{jobId}/executions/{executionId}">client.Jobs.Executions.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#JobExecutionService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, executionID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#JobExecutionGetParams">JobExecutionGetParams</a>) (<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#JobExecution">JobExecution</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /jobs/{jobId}/executions">client.Jobs.Executions.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#JobExecutionService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, jobID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#JobExecutionListParams">JobExecutionListParams</a>) ([]<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#JobExecution">JobExecution</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /jobs/{jobId}/executions/{executionId}">client.Jobs.Executions.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#JobExecutionService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, executionID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#JobExecutionDeleteParams">JobExecutionDeleteParams</a>) (<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#JobExecution">JobExecution</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Models

Methods:

- <code title="post /models">client.Models.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#ModelService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#ModelNewParams">ModelNewParams</a>) (<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#Model">Model</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /models/{modelName}">client.Models.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#ModelService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, modelName <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#Model">Model</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /models/{modelName}">client.Models.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#ModelService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, modelName <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#ModelUpdateParams">ModelUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#Model">Model</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /models">client.Models.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#ModelService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) ([]<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#Model">Model</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /models/{modelName}">client.Models.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#ModelService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, modelName <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#Model">Model</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /models/{modelName}/revisions">client.Models.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#ModelService.ListRevisions">ListRevisions</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, modelName <a href="https://pkg.go.dev/builtin#string">string</a>) ([]<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#RevisionMetadata">RevisionMetadata</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Policies

Params Types:

- <a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#PolicyParam">PolicyParam</a>
- <a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#PolicyLocationParam">PolicyLocationParam</a>
- <a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#PolicyMaxTokensParam">PolicyMaxTokensParam</a>
- <a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#PolicySpecParam">PolicySpecParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#Policy">Policy</a>
- <a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#PolicyLocation">PolicyLocation</a>
- <a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#PolicyMaxTokens">PolicyMaxTokens</a>
- <a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#PolicySpec">PolicySpec</a>

Methods:

- <code title="post /policies">client.Policies.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#PolicyService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#PolicyNewParams">PolicyNewParams</a>) (<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#Policy">Policy</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /policies/{policyName}">client.Policies.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#PolicyService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, policyName <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#Policy">Policy</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /policies/{policyName}">client.Policies.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#PolicyService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, policyName <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#PolicyUpdateParams">PolicyUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#Policy">Policy</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /policies">client.Policies.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#PolicyService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) ([]<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#Policy">Policy</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /policies/{policyName}">client.Policies.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#PolicyService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, policyName <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#Policy">Policy</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# PublicIPs

Response Types:

- <a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#PublicIPListResponse">PublicIPListResponse</a>

Methods:

- <code title="get /publicIps">client.PublicIPs.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#PublicIPService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#PublicIPListParams">PublicIPListParams</a>) (<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#PublicIPListResponse">PublicIPListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# VolumeTemplates

Params Types:

- <a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#VolumeTemplateParam">VolumeTemplateParam</a>
- <a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#VolumeTemplateSpecParam">VolumeTemplateSpecParam</a>
- <a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#VolumeTemplateStateParam">VolumeTemplateStateParam</a>
- <a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#VolumeTemplateVersionParam">VolumeTemplateVersionParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#VolumeTemplate">VolumeTemplate</a>
- <a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#VolumeTemplateSpec">VolumeTemplateSpec</a>
- <a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#VolumeTemplateState">VolumeTemplateState</a>
- <a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#VolumeTemplateVersion">VolumeTemplateVersion</a>
- <a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#VolumeTemplateDeleteVersionResponse">VolumeTemplateDeleteVersionResponse</a>

Methods:

- <code title="post /volume_templates">client.VolumeTemplates.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#VolumeTemplateService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, params <a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#VolumeTemplateNewParams">VolumeTemplateNewParams</a>) (<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#VolumeTemplate">VolumeTemplate</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /volume_templates/{volumeTemplateName}">client.VolumeTemplates.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#VolumeTemplateService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, volumeTemplateName <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#VolumeTemplate">VolumeTemplate</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /volume_templates">client.VolumeTemplates.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#VolumeTemplateService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) ([]<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#VolumeTemplate">VolumeTemplate</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /volume_templates/{volumeTemplateName}">client.VolumeTemplates.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#VolumeTemplateService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, volumeTemplateName <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#VolumeTemplate">VolumeTemplate</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /volume_templates/{volumeTemplateName}/versions/{versionName}">client.VolumeTemplates.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#VolumeTemplateService.DeleteVersion">DeleteVersion</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, versionName <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#VolumeTemplateDeleteVersionParams">VolumeTemplateDeleteVersionParams</a>) (<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#VolumeTemplateDeleteVersionResponse">VolumeTemplateDeleteVersionResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /volume_templates/{volumeTemplateName}">client.VolumeTemplates.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#VolumeTemplateService.Upsert">Upsert</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, volumeTemplateName <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#VolumeTemplateUpsertParams">VolumeTemplateUpsertParams</a>) (<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#VolumeTemplate">VolumeTemplate</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Volumes

Params Types:

- <a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#VolumeParam">VolumeParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#Volume">Volume</a>

Methods:

- <code title="post /volumes">client.Volumes.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#VolumeService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#VolumeNewParams">VolumeNewParams</a>) (<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#Volume">Volume</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /volumes/{volumeName}">client.Volumes.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#VolumeService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, volumeName <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#Volume">Volume</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /volumes">client.Volumes.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#VolumeService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) ([]<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#Volume">Volume</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /volumes/{volumeName}">client.Volumes.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#VolumeService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, volumeName <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#Volume">Volume</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Workspaces

Response Types:

- <a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#Workspace">Workspace</a>
- <a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#WorkspaceRuntime">WorkspaceRuntime</a>

Methods:

- <code title="get /workspaces/{workspaceName}">client.Workspaces.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#WorkspaceService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, workspaceName <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#Workspace">Workspace</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /workspaces">client.Workspaces.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#WorkspaceService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) ([]<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#Workspace">Workspace</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Sandboxes

Params Types:

- <a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#ExpirationPolicyParam">ExpirationPolicyParam</a>
- <a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#PortParam">PortParam</a>
- <a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#SandboxParam">SandboxParam</a>
- <a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#SandboxLifecycleParam">SandboxLifecycleParam</a>
- <a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#SandboxRuntimeParam">SandboxRuntimeParam</a>
- <a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#SandboxSpecParam">SandboxSpecParam</a>
- <a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#VolumeAttachmentParam">VolumeAttachmentParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#ExpirationPolicy">ExpirationPolicy</a>
- <a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#Port">Port</a>
- <a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#Sandbox">Sandbox</a>
- <a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#SandboxLifecycle">SandboxLifecycle</a>
- <a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#SandboxRuntime">SandboxRuntime</a>
- <a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#SandboxSpec">SandboxSpec</a>
- <a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#VolumeAttachment">VolumeAttachment</a>
- <a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#SandboxGetHubResponse">SandboxGetHubResponse</a>

Methods:

- <code title="post /sandboxes">client.Sandboxes.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#SandboxService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#SandboxNewParams">SandboxNewParams</a>) (<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#Sandbox">Sandbox</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /sandboxes/{sandboxName}">client.Sandboxes.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#SandboxService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, sandboxName <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#SandboxGetParams">SandboxGetParams</a>) (<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#Sandbox">Sandbox</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /sandboxes/{sandboxName}">client.Sandboxes.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#SandboxService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, sandboxName <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#SandboxUpdateParams">SandboxUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#Sandbox">Sandbox</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /sandboxes">client.Sandboxes.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#SandboxService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) ([]<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#Sandbox">Sandbox</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /sandboxes/{sandboxName}">client.Sandboxes.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#SandboxService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, sandboxName <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#Sandbox">Sandbox</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /sandbox/hub">client.Sandboxes.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#SandboxService.GetHub">GetHub</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) ([]<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#SandboxGetHubResponse">SandboxGetHubResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Process

Params Types:

- <a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#ProcessRequestParam">ProcessRequestParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#ProcessLogs">ProcessLogs</a>
- <a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#ProcessResponse">ProcessResponse</a>
- <a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#SandboxProcessKillResponse">SandboxProcessKillResponse</a>
- <a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#SandboxProcessStopResponse">SandboxProcessStopResponse</a>

Methods:

- <code title="post /process">client.Sandboxes.Process.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#SandboxProcessService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#SandboxProcessNewParams">SandboxProcessNewParams</a>) (<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#ProcessResponse">ProcessResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /process/{identifier}">client.Sandboxes.Process.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#SandboxProcessService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, identifier <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#ProcessResponse">ProcessResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /process">client.Sandboxes.Process.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#SandboxProcessService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) ([]<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#ProcessResponse">ProcessResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /process/{identifier}/kill">client.Sandboxes.Process.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#SandboxProcessService.Kill">Kill</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, identifier <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#SandboxProcessKillResponse">SandboxProcessKillResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /process/{identifier}/logs">client.Sandboxes.Process.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#SandboxProcessService.GetLogs">GetLogs</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, identifier <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#ProcessLogs">ProcessLogs</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /process/{identifier}">client.Sandboxes.Process.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#SandboxProcessService.Stop">Stop</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, identifier <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#SandboxProcessStopResponse">SandboxProcessStopResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Filesystem

Params Types:

- <a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#FilesystemWriteRequestParam">FilesystemWriteRequestParam</a>
- <a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#TreeRequestParam">TreeRequestParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#ContentSearchMatch">ContentSearchMatch</a>
- <a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#ContentSearchResponse">ContentSearchResponse</a>
- <a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#Directory">Directory</a>
- <a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#FilesystemRead">FilesystemRead</a>
- <a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#FilesystemReadWithContent">FilesystemReadWithContent</a>
- <a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#FindMatch">FindMatch</a>
- <a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#FindResponse">FindResponse</a>
- <a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#FuzzySearchMatch">FuzzySearchMatch</a>
- <a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#FuzzySearchResponse">FuzzySearchResponse</a>
- <a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#Subdirectory">Subdirectory</a>
- <a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#SandboxFilesystemDeleteResponse">SandboxFilesystemDeleteResponse</a>
- <a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#SandboxFilesystemDeleteTreeResponse">SandboxFilesystemDeleteTreeResponse</a>
- <a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#SandboxFilesystemGetResponseUnion">SandboxFilesystemGetResponseUnion</a>
- <a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#SandboxFilesystemGetTreeResponseUnion">SandboxFilesystemGetTreeResponseUnion</a>
- <a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#SandboxFilesystemWriteResponse">SandboxFilesystemWriteResponse</a>
- <a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#SandboxFilesystemWriteTreeResponseUnion">SandboxFilesystemWriteTreeResponseUnion</a>

Methods:

- <code title="delete /filesystem/{filePath}">client.Sandboxes.Filesystem.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#SandboxFilesystemService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, filePath <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#SandboxFilesystemDeleteParams">SandboxFilesystemDeleteParams</a>) (<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#SandboxFilesystemDeleteResponse">SandboxFilesystemDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /filesystem-content-search/{filePath}">client.Sandboxes.Filesystem.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#SandboxFilesystemService.ContentSearch">ContentSearch</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, filePath <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#SandboxFilesystemContentSearchParams">SandboxFilesystemContentSearchParams</a>) (<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#ContentSearchResponse">ContentSearchResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /filesystem/tree/{filePath}">client.Sandboxes.Filesystem.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#SandboxFilesystemService.DeleteTree">DeleteTree</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, filePath <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#SandboxFilesystemDeleteTreeParams">SandboxFilesystemDeleteTreeParams</a>) (<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#SandboxFilesystemDeleteTreeResponse">SandboxFilesystemDeleteTreeResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /filesystem-find/{filePath}">client.Sandboxes.Filesystem.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#SandboxFilesystemService.Find">Find</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, filePath <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#SandboxFilesystemFindParams">SandboxFilesystemFindParams</a>) (<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#FindResponse">FindResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /filesystem/{filePath}">client.Sandboxes.Filesystem.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#SandboxFilesystemService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, filePath <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#SandboxFilesystemGetParams">SandboxFilesystemGetParams</a>) (<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#SandboxFilesystemGetResponseUnion">SandboxFilesystemGetResponseUnion</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /filesystem/tree/{filePath}">client.Sandboxes.Filesystem.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#SandboxFilesystemService.GetTree">GetTree</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, filePath <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#SandboxFilesystemGetTreeResponseUnion">SandboxFilesystemGetTreeResponseUnion</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /filesystem-search/{filePath}">client.Sandboxes.Filesystem.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#SandboxFilesystemService.Search">Search</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, filePath <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#SandboxFilesystemSearchParams">SandboxFilesystemSearchParams</a>) (<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#FuzzySearchResponse">FuzzySearchResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /filesystem/{filePath}">client.Sandboxes.Filesystem.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#SandboxFilesystemService.Write">Write</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, filePath <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#SandboxFilesystemWriteParams">SandboxFilesystemWriteParams</a>) (<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#SandboxFilesystemWriteResponse">SandboxFilesystemWriteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /filesystem/tree/{filePath}">client.Sandboxes.Filesystem.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#SandboxFilesystemService.WriteTree">WriteTree</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, filePath <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#SandboxFilesystemWriteTreeParams">SandboxFilesystemWriteTreeParams</a>) (<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#SandboxFilesystemWriteTreeResponseUnion">SandboxFilesystemWriteTreeResponseUnion</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Watch

Methods:

- <code title="get /watch/filesystem/{filePath}">client.Sandboxes.Filesystem.Watch.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#SandboxFilesystemWatchService.Start">Start</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, filePath <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#SandboxFilesystemWatchStartParams">SandboxFilesystemWatchStartParams</a>) (<a href="https://pkg.go.dev/builtin#string">string</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Multipart

Params Types:

- <a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#CompleteRequestParam">CompleteRequestParam</a>
- <a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#InitiateRequestParam">InitiateRequestParam</a>
- <a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#PartInfoParam">PartInfoParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#InitiateResponse">InitiateResponse</a>
- <a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#ListPartsResponse">ListPartsResponse</a>
- <a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#ListUploadsResponse">ListUploadsResponse</a>
- <a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#UploadPartResponse">UploadPartResponse</a>
- <a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#SandboxFilesystemMultipartAbortResponse">SandboxFilesystemMultipartAbortResponse</a>
- <a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#SandboxFilesystemMultipartCompleteResponse">SandboxFilesystemMultipartCompleteResponse</a>

Methods:

- <code title="get /filesystem-multipart">client.Sandboxes.Filesystem.Multipart.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#SandboxFilesystemMultipartService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#ListUploadsResponse">ListUploadsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /filesystem-multipart/{uploadId}/abort">client.Sandboxes.Filesystem.Multipart.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#SandboxFilesystemMultipartService.Abort">Abort</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, uploadID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#SandboxFilesystemMultipartAbortResponse">SandboxFilesystemMultipartAbortResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /filesystem-multipart/{uploadId}/complete">client.Sandboxes.Filesystem.Multipart.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#SandboxFilesystemMultipartService.Complete">Complete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, uploadID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#SandboxFilesystemMultipartCompleteParams">SandboxFilesystemMultipartCompleteParams</a>) (<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#SandboxFilesystemMultipartCompleteResponse">SandboxFilesystemMultipartCompleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /filesystem-multipart/initiate/{filePath}">client.Sandboxes.Filesystem.Multipart.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#SandboxFilesystemMultipartService.Initiate">Initiate</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, filePath <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#SandboxFilesystemMultipartInitiateParams">SandboxFilesystemMultipartInitiateParams</a>) (<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#InitiateResponse">InitiateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /filesystem-multipart/{uploadId}/parts">client.Sandboxes.Filesystem.Multipart.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#SandboxFilesystemMultipartService.ListParts">ListParts</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, uploadID <a href="https://pkg.go.dev/builtin#string">string</a>) (<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#ListPartsResponse">ListPartsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /filesystem-multipart/{uploadId}/part">client.Sandboxes.Filesystem.Multipart.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#SandboxFilesystemMultipartService.UploadPart">UploadPart</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, uploadID <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#SandboxFilesystemMultipartUploadPartParams">SandboxFilesystemMultipartUploadPartParams</a>) (<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#UploadPartResponse">UploadPartResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Codegen

Params Types:

- <a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#ApplyEditRequestParam">ApplyEditRequestParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#ApplyEditResponse">ApplyEditResponse</a>
- <a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#RerankingResponse">RerankingResponse</a>

Methods:

- <code title="put /codegen/fastapply/{filePath}">client.Sandboxes.Codegen.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#SandboxCodegenService.Fastapply">Fastapply</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, filePath <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#SandboxCodegenFastapplyParams">SandboxCodegenFastapplyParams</a>) (<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#ApplyEditResponse">ApplyEditResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /codegen/reranking/{filePath}">client.Sandboxes.Codegen.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#SandboxCodegenService.Reranking">Reranking</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, filePath <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#SandboxCodegenRerankingParams">SandboxCodegenRerankingParams</a>) (<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#RerankingResponse">RerankingResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Previews

Params Types:

- <a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#PreviewParam">PreviewParam</a>
- <a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#PreviewMetadataParam">PreviewMetadataParam</a>
- <a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#PreviewSpecParam">PreviewSpecParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#Preview">Preview</a>
- <a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#PreviewMetadata">PreviewMetadata</a>
- <a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#PreviewSpec">PreviewSpec</a>

Methods:

- <code title="post /sandboxes/{sandboxName}/previews">client.Sandboxes.Previews.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#SandboxPreviewService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, sandboxName <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#SandboxPreviewNewParams">SandboxPreviewNewParams</a>) (<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#Preview">Preview</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /sandboxes/{sandboxName}/previews/{previewName}">client.Sandboxes.Previews.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#SandboxPreviewService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, previewName <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#SandboxPreviewGetParams">SandboxPreviewGetParams</a>) (<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#Preview">Preview</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="put /sandboxes/{sandboxName}/previews/{previewName}">client.Sandboxes.Previews.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#SandboxPreviewService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, previewName <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#SandboxPreviewUpdateParams">SandboxPreviewUpdateParams</a>) (<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#Preview">Preview</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /sandboxes/{sandboxName}/previews">client.Sandboxes.Previews.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#SandboxPreviewService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, sandboxName <a href="https://pkg.go.dev/builtin#string">string</a>) ([]<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#Preview">Preview</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /sandboxes/{sandboxName}/previews/{previewName}">client.Sandboxes.Previews.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#SandboxPreviewService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, previewName <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#SandboxPreviewDeleteParams">SandboxPreviewDeleteParams</a>) (<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#Preview">Preview</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Tokens

Params Types:

- <a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#PreviewTokenParam">PreviewTokenParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#PreviewToken">PreviewToken</a>
- <a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#SandboxPreviewTokenDeleteResponse">SandboxPreviewTokenDeleteResponse</a>

Methods:

- <code title="post /sandboxes/{sandboxName}/previews/{previewName}/tokens">client.Sandboxes.Previews.Tokens.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#SandboxPreviewTokenService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, previewName <a href="https://pkg.go.dev/builtin#string">string</a>, params <a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#SandboxPreviewTokenNewParams">SandboxPreviewTokenNewParams</a>) (<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#PreviewToken">PreviewToken</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /sandboxes/{sandboxName}/previews/{previewName}/tokens">client.Sandboxes.Previews.Tokens.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#SandboxPreviewTokenService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, previewName <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#SandboxPreviewTokenGetParams">SandboxPreviewTokenGetParams</a>) ([]<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#PreviewToken">PreviewToken</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /sandboxes/{sandboxName}/previews/{previewName}/tokens/{tokenName}">client.Sandboxes.Previews.Tokens.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#SandboxPreviewTokenService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, tokenName <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#SandboxPreviewTokenDeleteParams">SandboxPreviewTokenDeleteParams</a>) (<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go">blaxel</a>.<a href="https://pkg.go.dev/github.com/blaxel-ai/sdk-go#SandboxPreviewTokenDeleteResponse">SandboxPreviewTokenDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
