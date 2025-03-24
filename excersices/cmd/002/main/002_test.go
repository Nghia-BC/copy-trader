package main

import "testing"

func TestFactorial(t *testing.T) {

	// Negative case
	expected := 0
	actual, err := calF(-1)
	if err == nil {
		t.Errorf("calF = %v, want %v", actual, err)
	}

	// input == 0
	expected = 1
	actual, err = calF(0)
	if err == nil && actual != uint64(expected) {
		t.Errorf("Expected '%v' but got '%v'", 1, actual)
	}

	expected = 40320
	actual, _ = calF(8)

	if actual != uint64(expected) {
		t.Errorf("expected = %d, actual = %d", expected, actual)
	}
}
