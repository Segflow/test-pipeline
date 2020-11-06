package main

import "testing"

func TestFactRecursive(t *testing.T) {
	r := factRecursive(-1)
	if r != -1 {
		t.Errorf("factRecursive(-1) should be -1, got %d", r)
	}

	r = factRecursive(0)
	if r != 1 {
		t.Errorf("factRecursive(0) should be 1, got %d", r)
	}

	r = factRecursive(1)
	if r != 1 {
		t.Errorf("factRecursive(1) should be 1, got %d", r)
	}

	r = factRecursive(5)
	if r != 120 {
		t.Errorf("factRecursive(5) should be 120, got %d", r)
	}
}

func TestFactIter(t *testing.T) {
	r := factIter(-1)
	if r != -1 {
		t.Errorf("factIter(-1) should be -1, got %d", r)
	}

	r = factIter(0)
	if r != 1 {
		t.Errorf("factIter(0) should be 1, got %d", r)
	}

	r = factIter(1)
	if r != 1 {
		t.Errorf("factIter(1) should be 1, got %d", r)
	}

	r = factIter(5)
	if r != 120 {
		t.Errorf("factIter(5) should be 120, got %d", r)
	}
}

func BenchmarkFactRecusive(b *testing.B) {
	for i := 0; i < b.N; i++ {
		factRecursive(500)
	}
}

func BenchmarkFactIterative(b *testing.B) {
	for i := 0; i < b.N; i++ {
		factIter(500)
	}
}

func BenchmarkFact(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fact(500)
	}
}
