package urls

import (
	"testing"

	"github.com/nelsonsaake/go/pty"
)

func TestParse(t *testing.T) {

	urlStr := "https://example.com:8080/api/v1/resource?key1=value1&key2=value2"

	url, err := Parse(urlStr)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	urlstr := url.String()

	expected := urlStr
	if urlstr != expected {
		t.Errorf("expected %q, got %q", expected, urlstr)
	}

	pty.Println(url.Scheme)
	pty.Println(url.Host)
	pty.Println(url.Hostname())
	pty.Println(url.Port())
	pty.Println(url.Path)
	pty.Println(url.Query())
}
