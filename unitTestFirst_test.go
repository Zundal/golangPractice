package main

import (
	"testing"
)

func add(a, b int) int {
	return a + b
}

func TestFirst(t *testing.T) {
	result := add(2, 3)
	if result != 5 {
		t.Errorf("add(2, 3) = %d; want 5", result)
	}
}
