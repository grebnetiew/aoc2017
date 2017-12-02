package main

import "testing"

func TestChecksum(t *testing.T) {
	in1 := []int{5, 1, 9, 5}
	in2 := []int{7, 5, 3}
	in3 := []int{2, 4, 6, 8}

	if c := checksum(in1); c != 8 {
		t.Errorf("First test failed, expected 8, got %d", c)
	}
	if c := checksum(in2); c != 4 {
		t.Errorf("First test failed, expected 4, got %d", c)
	}
	if c := checksum(in3); c != 6 {
		t.Errorf("First test failed, expected 6, got %d", c)
	}
}