package tests

import (
	"testing"

	"github.com/nelsonsaake/go/axios"
	"github.com/nelsonsaake/go/pty"
)

func TestAxios(t *testing.T) {

	t.Log("building client ...")

	client := axios.
		NewClientBuilder().
		WithBaseUrl("{BaseURL}/{Version}").
		WithHeader("content-type", "application/json").
		WithHeader("accept", "application/json").
		WithEnvironment("local", EnvLocal).
		WithSelectedEnvironment("local").
		Build()

	t.Log("build client successful ...")

	res, err := client.Get("/ping")
	if err != nil {
		t.Fatal(err)
	}

	t.Log("client", pty.JSON(res))
}
