package main

import "testing"

func TestSum(t *testing.T) {
	result := Sum(1, 1)
	if result != 2 {
		t.Errorf("Sum is fail, 1 + 1 must be 2 but receive %d", result)
	}
}

func TestMutiple(t *testing.T) {
	result := Mutiple(2, 2)
	if result != 4 {
		t.Errorf("Sum is fail, 2 * 2 must be 4 but receive %d", result)
	}
}
