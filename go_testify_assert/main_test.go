package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSum(t *testing.T) {
	result := Sum(1, 1)
	assert.Equal(t, 3, result, "they should be equal")
}

func TestMutiple(t *testing.T) {
	result := Mutiple(2, 2)
	assert.Equal(t, 4, result, "they should be equal")
}
