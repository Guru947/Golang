package main

import "testing"

func TestSum(t *testing.T) {

	x := sum(1, 3, 5)
	if x != 9 {
		t.Error("expected this", 8, "got", x)
	}
}
