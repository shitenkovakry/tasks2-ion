package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Case1(t *testing.T) {
	array := [][]string{
		{"1", "2", "3"},
		{"4", "5", "6"},
		{"7", "8", "9"},
	}

	expected := [][]string{
		{"5"},
	}

	actual := FindInsidePerimetr(array)

	assert.Equal(t, expected, actual)
}

func Test_Case2(t *testing.T) {
	array := [][]string{
		{"q", "w", "e", "r", "t"},
		{"y", "u", "i", "o", "p"},
		{"a", "s", "d", "f", "g"},
		{"h", "j", "k", "l", "z"},
		{"x", "c", "v", "b", "n"},
	}

	expected := [][]string{
		{"u", "i", "o"},
		{"s", "d", "f"},
		{"j", "k", "l"},
	}

	actual := FindInsidePerimetr(array)

	assert.Equal(t, expected, actual)
}
