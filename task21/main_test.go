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

	argument := Argument{
		Direction: Z,
	}

	expected := []int{1, 2, 3, 5, 7, 8, 9}
	actual := DrawingResult(array, argument)

	assert.Equal(t, expected, actual)
}

func Test_Case2(t *testing.T) {
	array := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	argument := Argument{
		Direction: N,
	}

	expected := []int{7, 4, 1, 5, 9, 6, 3}
	actual := DrawingResult(array, argument)

	assert.Equal(t, expected, actual)
}

func Test_Case3(t *testing.T) {
	array := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	argument := Argument{
		Direction: S,
	}

	expected := []int{3, 2, 1, 4, 5, 6, 9, 8, 7}
	actual := DrawingResult(array, argument)

	assert.Equal(t, expected, actual)
}
