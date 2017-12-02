package main

import "testing"

func TestChecksum(t *testing.T) {
	in1 := []int{5, 9, 2, 8}
	in2 := []int{9, 4, 7, 3}
	in3 := []int{3, 8, 6, 5}

	if c := checksum(in1); c != 4 {
		t.Errorf("First test failed, expected 4, got %d", c)
	}
	if c := checksum(in2); c != 3 {
		t.Errorf("First test failed, expected 3, got %d", c)
	}
	if c := checksum(in3); c != 2 {
		t.Errorf("First test failed, expected 2, got %d", c)
	}
}