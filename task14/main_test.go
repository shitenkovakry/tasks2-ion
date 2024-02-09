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
		{2},
		{4, 6},
		{8},
	}

	actual := Result(array)

	assert.Equal(t, expected, actual)
}

func Test_Case2(t *testing.T) {
	array := [][]int{
		{1, 5, 10, 14},
		{2, 6, 11, 15},
		{3, 7, 12, 16},
		{4, 8, 13, 17},
	}

	expected := [][]int{
		{5, 10},
		{2, 15},
		{3, 16},
		{8, 13},
	}

	actual := Result(array)

	assert.Equal(t, expected, actual)
}
