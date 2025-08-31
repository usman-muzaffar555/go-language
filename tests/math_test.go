package tests

import (
	"fmt"
	"testing"
)

func TestAddition(t *testing.T) {
	fmt.Println("Running TestAddition")
	sum := 2 + 3
	if sum != 5 {
		t.Errorf("expected 5, got %d", sum)
	}
}

func TestMultiplication(t *testing.T) {
	fmt.Println("Running TestMultiplication")
	product := 4 * 3
	if product != 12 {
		t.Errorf("expected 12, got %d", product)
	}
}

