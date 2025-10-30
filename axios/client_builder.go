package axios

// ClientBuilder is a builder for constructing axios.Client instances.
type ClientBuilder struct {
	client *Client
}

// NewClientBuilder creates a new ClientBuilder.
func NewClientBuilder() *ClientBuilder {
	return &ClientBuilder{
		client: NewClient(),
	}
}

// WithBaseUrl sets the base URL for the client.
func (b *ClientBuilder) WithBaseUrl(base string) *ClientBuilder {
	b.client.BaseUrl = base
	return b
}

// WithHeader adds a single header to the client.
func (b *ClientBuilder) WithHeader(key, value string) *ClientBuilder {
	if b.client.Headers == nil {
		b.client.Headers = make(map[string]string)
	}
	b.client.Headers[key] = value
	return b
}

// WithHeaders adds multiple headers to the client.
func (b *ClientBuilder) WithHeaders(headers map[string]string) *ClientBuilder {
	if b.client.Headers == nil {
		b.client.Headers = make(map[string]string)
	}
	for k, v := range headers {
		b.client.Headers[k] = v
	}
	return b
}

// Add builder: set multiple environments and return client for chaining
func (b *ClientBuilder) WithEnvironments(envs map[string]map[string]string) *ClientBuilder {
	b.client.Environments = envs
	return b
}

// Add builder: set multiple environments and return client for chaining
func (b *ClientBuilder) WithEnvironment(name string, envs map[string]string) *ClientBuilder {
	b.client.Environments[name] = envs
	return b
}

// Add builder: set active environment name and return client for chaining
func (b *ClientBuilder) WithSelectedEnvironment(name string) *ClientBuilder {
	b.client.SelectedEnvironment = name
	return b
}

// Build returns the configured Client instance.
func (b *ClientBuilder) Build() *Client {
	return b.client
}
