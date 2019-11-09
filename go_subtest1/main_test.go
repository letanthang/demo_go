package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type TestCase struct {
	Name   string
	Input1 int
	Input2 int
	Expect int
}

func TestSum(t *testing.T) {

	testCases := []TestCase{
		{"case 1: 1 + 1", 1, 1, 2},
		{"case 2: 1 + 2", 1, 2, 3},
		{"case 3: 2 + 2", 2, 2, 4},
	}

	// subtest
	for _, tc := range testCases {
		result := Sum(tc.Input1, tc.Input2)
		t.Run(tc.Name, func(t *testing.T) {
			assert.Equal(t, tc.Expect, result)
		})
	}

}

func TestMultiply(t *testing.T) {
	testCases := []TestCase{
		{"case 1: 1 x 1", 1, 1, 1},
		{"case 2: 1 x 2", 1, 2, 2},
		{"case 3: 2 x 2", 2, 2, 4},
	}

	// subtest
	for _, tc := range testCases {
		result := Multiply(tc.Input1, tc.Input2)
		t.Run(tc.Name, func(t *testing.T) {
			assert.Equal(t, tc.Expect, result)
		})
	}
}
