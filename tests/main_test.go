package tests

import (
	"fmt"
	"os"
	"testing"
)

// TestMain lets us do setup/teardown around all tests
func TestMain(m *testing.M) {
	fmt.Println("Setting up test environment...")

	// Run all tests
	exitVal := m.Run()

	fmt.Println("Tearing down test environment...")

	// Exit with the same status code
	os.Exit(exitVal)
}

