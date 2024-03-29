package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Case1(t *testing.T) {
	array := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	expected := [][]int{
		{1, 2},
		{4},
	}

	actual := FindElementsWhichHigherSideDiagonal(array)

	assert.Equal(t, expected, actual)
}
