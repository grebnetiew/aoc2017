package main

import "testing"

func TestCoordsum(t *testing.T) {
	if c := coordsum(1); c != 0 {
		t.Errorf("1. test failed, expected 0, got %d", c)
	}
	if c := coordsum(12); c != 3 {
		t.Errorf("2. test failed, expected 3, got %d", c)
	}
	if c := coordsum(23); c != 2 {
		t.Errorf("3. test failed, expected 2, got %d", c)
	}
	if c := coordsum(1024); c != 31 {
		t.Errorf("4. test failed, expected 31, got %d", c)
	}
}