package main

import "testing"

func TestOps(t *testing.T) {
	in2 := "12/34"
	in3 := "123/456/789"
	in4 := "1234/5678/9abc/defg"

	if s := hflip(in2); s != "21/43" {
		t.Errorf("H expected 21/43, got %s", s)
	}
	if s := hflip(in3); s != "321/654/987" {
		t.Errorf("H expected 321/654/987, got %s", s)
	}
	if s := vflip(in2); s != "34/12" {
		t.Errorf("V expected 34/12, got %s", s)
	}
	if s := vflip(in3); s != "789/456/123" {
		t.Errorf("V expected 789/456/123, got %s", s)
	}
	if s := rot(in2); s != "31/42" {
		t.Errorf("R expected 31/42, got %s", s)
	}
	if s := rot(in3); s != "741/852/963" {
		t.Errorf("R expected 741/852/963, got %s", s)
	}
	s1, s2, s3, s4 := split(in4)
	if s1 != "12/56" {
		t.Errorf("S expected 12/56, got %s", s1)
	}
	if s2 != "34/78" {
		t.Errorf("S expected 34/78, got %s", s2)
	}
	if s3 != "9a/de" {
		t.Errorf("S expected 9a/de, got %s", s3)
	}
	if s4 != "bc/fg" {
		t.Errorf("S expected bc/fg, got %s", s4)
	}
}
