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
		{"c", "a", "d", "b", "c"},
		{"c", "h", "c", "b", "c"},
		{"g", "b", "d", "h", "a"},
	}

	actual := FindTheNaturalUnionOfTwoMatrices(arrayOne, arrayTwo)

	assert.Equal(t, expected, actual)
}
