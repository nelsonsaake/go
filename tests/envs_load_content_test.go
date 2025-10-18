package tests

import (
	"os"
	"testing"

	"github.com/nelsonsaake/go/envs"
)

func TestLoadContent(t *testing.T) {
	// Clean up any existing env vars
	os.Unsetenv("TEST_VAR1")
	os.Unsetenv("TEST_VAR2")
	os.Unsetenv("TEST_VAR3")

	// Test loading single content
	content1 := `
TEST_VAR1=value1
TEST_VAR2=value2
`
	err := envs.LoadContent(content1)
	if err != nil {
		t.Fatalf("Failed to load content: %v", err)
	}

	// Verify variables were set
	if got := os.Getenv("TEST_VAR1"); got != "value1" {
		t.Errorf("Expected TEST_VAR1=value1, got %s", got)
	}
	if got := os.Getenv("TEST_VAR2"); got != "value2" {
		t.Errorf("Expected TEST_VAR2=value2, got %s", got)
	}
}

func TestLoadContentMultiple(t *testing.T) {
	// Clean up
	os.Unsetenv("TEST_A")
	os.Unsetenv("TEST_B")
	os.Unsetenv("TEST_C")

	content1 := "TEST_A=valueA"
	content2 := `TEST_B=valueB
TEST_C=valueC`

	err := envs.LoadContent(content1, content2)
	if err != nil {
		t.Fatalf("Failed to load multiple contents: %v", err)
	}

	// Verify all variables were set
	if got := os.Getenv("TEST_A"); got != "valueA" {
		t.Errorf("Expected TEST_A=valueA, got %s", got)
	}
	if got := os.Getenv("TEST_B"); got != "valueB" {
		t.Errorf("Expected TEST_B=valueB, got %s", got)
	}
	if got := os.Getenv("TEST_C"); got != "valueC" {
		t.Errorf("Expected TEST_C=valueC, got %s", got)
	}
}

func TestLoadContentDeduplication(t *testing.T) {
	// Clean up
	os.Unsetenv("DEDUP_TEST")

	content := "DEDUP_TEST=original"

	// Load same content twice
	err := envs.LoadContent(content, content)
	if err != nil {
		t.Fatalf("Failed to load content: %v", err)
	}

	if got := os.Getenv("DEDUP_TEST"); got != "original" {
		t.Errorf("Expected DEDUP_TEST=original, got %s", got)
	}
}

func TestLoadContentEmpty(t *testing.T) {
	// Test with empty content
	err := envs.LoadContent("", "   ", "\n\n")
	if err != nil {
		t.Errorf("LoadContent should handle empty content gracefully, got error: %v", err)
	}
}
