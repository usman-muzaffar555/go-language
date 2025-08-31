package tests

import (
	"fmt"
	"strings"
	"testing"
)

func TestUppercase(t *testing.T) {
	fmt.Println("Running TestUppercase")
	got := strings.ToUpper("golang")
	want := "GOLANG"
	if got != want {
		t.Errorf("expected %s, got %s", want, got)
	}
}

func TestContains(t *testing.T) {
	fmt.Println("Running TestContains")
	if !strings.Contains("postgresql", "sql") {
		t.Errorf("expected string to contain 'sql'")
	}
}

