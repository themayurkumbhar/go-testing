package messages

import "testing"

func TestGreet(t *testing.T) {
	got := Greet("GoLanger")
	expect := "Hello, GoLanger!\n"

	if got != expect {
		t.Errorf("Expected %q but got %q", expect, got)
	}
}

func TestDepart(t *testing.T) {
	// here depart is avaiable as this test file is part of same package "messages"
	got := depart("GoLanger")
	expect := "Good bye GoLanger.\n"

	if got != expect {
		t.Errorf("Expected %q but got %q", expect, got)
	}

}
