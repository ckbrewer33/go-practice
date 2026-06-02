package main

import "testing"

func TestAppend(t *testing.T) {
	a := ArrayList{}
	a.Append(1)

	if len(a.data) != 1 {
		t.Error("expected length of data array to be 1, but was", len(a.data))
	}
	if a.len != 1 {
		t.Error("expected struct to report length of 1, but got", a.len)
	}
	if a.data[0] != 1 {
		t.Error("expected 1 in position 0, but found", a.data[0])
	}
}
