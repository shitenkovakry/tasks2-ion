package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Case1(t *testing.T) {
	arrayOne := [][]string{
		{"a", "b", "c"},
		{"c", "a", "d"},
		{"c", "h", "c"},
		{"g", "b", "d"},
	}

	arrayTwo := [][]string{
		{"g", "h", "a"},
		{"c", "b", "c"},
		{"h", "d", "d"},
	}

	expected := [][]string{
		{"a", "b", "c", "g", "h", "a"},
		{"a", "b", "c", "c", "b", "c"},
		{"a", "b", "c", "h", "d", "d"},
		{"c", "a", "d", "g", "h", "a"},
		{"c", "a", "d", "c", "b", "c"},
		{"c", "a", "d", "h", "d", "d"},
		{"c", "h", "c", "g", "h", "a"},
		{"c", "h", "c", "c", "b", "c"},
		{"c", "h", "c", "h", "d", "d"},
		{"g", "b", "d", "g", "h", "a"},
		{"g", "b", "d", "c", "b", "c"},
		{"g", "b", "d", "h", "d", "d"},
	}

	actual := FindTheCartesianProductOfTwoMatrices(arrayOne, arrayTwo)

	assert.Equal(t, expected, actual)
}
