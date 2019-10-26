package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSum(t *testing.T) {
	result := Sum(1, 1)
	if result != 2 {
		t.Errorf("Sum is fail, 1 + 1 must be 2 but receive %d", result)
	}
}

func TestMutiple(t *testing.T) {
	result := Mutiple(2, 2)
	if result != 4 {
		t.Errorf("Mutiple is fail, 2 * 2 must be 4 but receive %d", result)
	}
}

func TestMutipleTestify(t *testing.T) {
	tests := []struct {
		Name     string
		Input1   int
		Input2   int
		Expected int
	}{
		{"case1", 1, 2, 2},
		{"case2", 2, 2, 4},
		{"case2", 4, 2, 8},
	}

	for _, tc := range tests {
		t.Run(tc.Name, func(t *testing.T) {
			result := Mutiple(tc.Input1, tc.Input2)
			assert.Equal(t, tc.Expected, result, "Phai bang nhau")
		})
	}

}
