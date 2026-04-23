package urls

import (
	"fmt"
	"net"
	"net/url"
)

type Builder struct {
	url.URL
	queries map[string]any
	port    int
}

func (b *Builder) WithScheme(scheme string) *Builder {
	b.Scheme = scheme
	return b
}

func (b *Builder) WithHost(host string) *Builder {
	b.Host = host
	return b
}

func (b *Builder) WithPath(path string) *Builder {
	b.Path = path
	return b
}

func (b *Builder) WithQueries(queries map[string]any) *Builder {
	if b.queries == nil {
		b.queries = make(map[string]any)
	}
	for k, v := range queries {
		b.queries[k] = v
	}
	return b
}

func (b *Builder) WithQuery(key string, value any) *Builder {
	b.WithQueries(map[string]any{key: value})
	return b
}

func (b *Builder) WithPort(port int) *Builder {
	b.port = port
	return b
}

func (b *Builder) Build() (string, error) {

	die := func(f string, a ...any) (string, error) {
		return "", fmt.Errorf(f, a...)
	}

	// QUERIES

	if b.queries == nil {
		b.queries = make(map[string]any)
	}

	values := url.Values{}
	for k, v := range b.queries {
		values.Set(k, fmt.Sprintf("%v", v))
	}

	b.RawQuery = values.Encode()

	// PORT

	if b.port < 0 || b.port > 65535 {
		return die("invalid port: %d", b.port)
	}

	host := b.Host
	if h, _, err := net.SplitHostPort(b.Host); err == nil {
		host = h
	}

	if b.port != 0 {
		b.Host = fmt.Sprintf("%s:%d", host, b.port)
	}

	// BUILD

	return b.URL.String(), nil
}

func NewBuilder() *Builder {
	return &Builder{}
}
