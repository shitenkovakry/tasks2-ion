package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Case2(t *testing.T) {
	arrayOne := [][]string{
		{"a3", "b1"},
		{"a2", "b3"},
		{"a1", "b2"},
	}

	arrayTwo := [][]string{
		{"a3", "b3"},
		{"a2", "b3"},
		{"a1", "b2"},
		{"a1", "b1"},
	}

	expected := [][]string{
		{"a3", "b1"},
		{"a2", "b3"},
		{"a1", "b2"},
		{"a3", "b3"},
		{"a1", "b1"},
	}

	actual := CombineTwoMatrixIntoOne(arrayOne, arrayTwo)

	assert.Equal(t, expected, actual)
}
