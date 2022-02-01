package main_test

import (
	"testing"
)

func TestMultiplication(testingPointer *testing.T) {
	actual := 5 * 2
	expected := 10
	if actual != expected {
		testingPointer.Errorf("Expected '%v' but got '%v'", expected, actual)
	}
}

func TestAddition(t *testing.T) {
	actual := 5 - 3
	expected := 3
	if actual != expected {
		t.Errorf("Expected this '%v', but got this '%v'", expected, actual)
	}
}
