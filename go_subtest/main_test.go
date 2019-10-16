package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSum(t *testing.T) {
	tests := []struct {
		a    int
		b    int
		want int
	}{
		{1, 1, 3},
		{1, 9, 10},
		{2, 3, 5},
	}
	for _, tc := range tests {
		name := fmt.Sprintf("Test %d + %d", tc.a, tc.b)
		t.Run(name, func(t *testing.T) {
			got := Sum(tc.a, tc.b)
			assert.Equal(t, tc.want, got, "they should be equal")
		})
	}

}

func TestMutiple(t *testing.T) {
	result := Mutiple(2, 2)
	if result != 4 {
		t.Errorf("Mutiple is fail, 2 * 2 must be 4 but receive %d", result)
	}
}
