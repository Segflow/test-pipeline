package main

import "testing"

func TestFact(t *testing.T) {

	r := fact(-1)
	if r != -1 {
		t.Errorf("Fact(-1) should be -1, got %d", r)
	}

	r = fact(0)
	if r != 1 {
		t.Errorf("Fact(0) should be 1, got %d", r)
	}

	r = fact(1)
	if r != 1 {
		t.Errorf("Fact(1) should be 1, got %d", r)
	}

	r = fact(5)
	if r != 120 {
		t.Errorf("Fact(5) should be 120, got %d", r)
	}
}
