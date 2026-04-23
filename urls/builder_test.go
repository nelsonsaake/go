package urls

import (
	"testing"
)

func TestBuilder(t *testing.T) {

	url, err := NewBuilder().
		WithScheme("https").
		WithHost("example.com").
		WithPath("/api/v1/resource").
		WithQuery("key1", "value1").
		WithQuery("key2", "value2").
		WithPort(8080).
		Build()

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expected := "https://example.com:8080/api/v1/resource?key1=value1&key2=value2"
	if url != expected {
		t.Errorf("expected %q, got %q", expected, url)
	}
}

func TestBuildWithoutPort(t *testing.T) {

	url, err := NewBuilder().
		WithScheme("https").
		WithHost("example.com").
		WithPath("/api/v1/resource").
		WithQuery("key1", "value1").
		WithQuery("key2", "value2").
		Build()

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expected := "https://example.com/api/v1/resource?key1=value1&key2=value2"
	if url != expected {
		t.Errorf("expected %q, got %q", expected, url)
	}
}
