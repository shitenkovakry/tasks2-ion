package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_case1(t *testing.T) {
	// Create a 5x5 array of ints
	arr := [][]int{
		{1, 2, 3, 4, 5},
		{6, 7, 8, 9, 10},
		{11, 12, 13, 14, 15},
		{16, 17, 18, 19, 20},
		{21, 22, 23, 24, 25},
	}

	spiralArgument := SpiralArgument{
		Direction: ClockwiseFrom00,
	}

	expected := []int{
		1, 2, 3, 4, // [0, 0] -> [0, 4]
		5, 10, 15, 20, // [0, 4] -> [4, 4]
		25, 24, 23, 22, // [4, 4] -> [4, 0]
		21, 16, 11, // [4, 0] -> [1, 0]

		6, 7, 8, // [1, 0] -> [1, 3]
		9, 14, // [1, 3] -> [3, 3]
		19, 18, // [3, 3] -> [3, 1]
		17, // [3, 1] -> [2, 1]

		12, // [2, 1] -> [2, 2]
		13, // [2, 2] -> [2, 2]
	}

	actual := SpiralOrder(arr, spiralArgument)
	assert.Equal(t, expected, actual)
}

func Test_case2(t *testing.T) {
	arr := [][]int{
		{1, 2, 3, 4, 5, 6, 7, 8},
		{9, 10, 11, 12, 13, 14, 15, 16},
		{17, 18, 19, 20, 21, 22, 23, 24},
		{25, 26, 27, 28, 29, 30, 31, 32},
		{33, 34, 35, 36, 37, 38, 39, 40},
		{41, 42, 43, 44, 45, 46, 47, 48},
		{49, 50, 51, 52, 53, 54, 55, 56},
		{57, 58, 59, 60, 61, 62, 63, 64},
	}

	spiralArgument := SpiralArgument{
		Direction: ClockwiseFrom00,
	}

	expected := []int{
		1, 2, 3, 4, 5, 6, 7, // [0, 0] -> [0, 7]
		8, 16, 24, 32, 40, 48, 56, // [0, 7] -> [7, 7]
		64, 63, 62, 61, 60, 59, 58, // [7, 7] -> [7, 0]
		57, 49, 41, 33, 25, 17, // [7, 0] -> [1, 0]
		9, 10, 11, 12, 13, 14, // [1, 0] -> [1, 6]
		15, 23, 31, 39, 47, // [1, 6] -> [6, 6]
		55, 54, 53, 52, 51, // [6, 6] -> [6, 1]
		50, 42, 34, 26, // [6, 1] -> [2, 1]
		18, 19, 20, 21, // [2, 1] -> [2, 5]
		22, 30, 38, // [2, 5] -> [5, 5]
		46, 45, 44, // [5, 5] -> [5, 2]
		43, 35, // [5, 2] -> [3, 2]
		27, 28, // [3, 2] -> [3, 4]
		29, // [3, 4] -> [4, 4]
		37, // [4, 4] -> [4, 3]
		36, // [4, 3] -> [4, 3]
	}

	actual := SpiralOrder(arr, spiralArgument)
	assert.Equal(t, expected, actual)
}

/*
// [0, 0] -> [0, 4]
// [0, 4] -> [4, 4]
// [4, 4] -> [4, 0]
// [4, 0] -> [1, 0]
// [1, 0] -> [1, 3]
// [1, 3] -> [3, 3]
// [3, 3] -> [3, 1]
// [3, 1] -> [2, 1]
// [2, 1] -> [2, 2]
// [2, 2] -> [2, 2]



// [0, 0] -> [0, 7]
// [0, 7] -> [7, 7]
// [7, 7] -> [7, 0]
// [7, 0] -> [1, 0]

// [1, 0] -> [1, 6]
// [1, 6] -> [6, 6]
// [6, 6] -> [6, 1]
// [6, 1] -> [2, 1]

// [2, 1] -> [2, 5]
// [2, 5] -> [5, 5]
// [5, 5] -> [5, 2]
// [5, 2] -> [3, 2]

// [3, 2] -> [3, 4]
// [3, 4] -> [4, 4]
// [4, 4] -> [4, 3]
// [4, 3] -> [4, 3]


N = 4

left := 0
rigth := N


i := 0
left++ // 0
rigth-- // 4

i++ // 1
left++ // 1
rigth-- // 3

i++ // 2
left++ // 2
rigth-- // 2

0,   4,  1,   3,      2




N = 7

left := 0
rigth := N

i := 0
left++ // 0
rigth-- // 7

i++ // 1
left++ // 1
rigth-- // 6

i++ // 2
left++ // 2
rigth-- // 5

i++ // 3
left++ // 3
rigth-- // 4




0 7 1 6 2 5 3 4
*/

func Test_case3(t *testing.T) {
	arr := [][]int{
		{1, 2, 3, 4, 5, 6, 7, 8},
		{9, 10, 11, 12, 13, 14, 15, 16},
		{17, 18, 19, 20, 21, 22, 23, 24},
		{25, 26, 27, 28, 29, 30, 31, 32},
		{33, 34, 35, 36, 37, 38, 39, 40},
		{41, 42, 43, 44, 45, 46, 47, 48},
		{49, 50, 51, 52, 53, 54, 55, 56},
		{57, 58, 59, 60, 61, 62, 63, 64},
	}

	expected := []int{1, 2, 3, 4, 5, 6, 7}
	actual := Walk(arr, 0, 0, 0, 7)
	assert.Equal(t, expected, actual)
}

func Test_case4(t *testing.T) {
	arr := [][]int{
		{1, 2, 3, 4, 5, 6, 7, 8},
		{9, 10, 11, 12, 13, 14, 15, 16},
		{17, 18, 19, 20, 21, 22, 23, 24},
		{25, 26, 27, 28, 29, 30, 31, 32},
		{33, 34, 35, 36, 37, 38, 39, 40},
		{41, 42, 43, 44, 45, 46, 47, 48},
		{49, 50, 51, 52, 53, 54, 55, 56},
		{57, 58, 59, 60, 61, 62, 63, 64},
	}

	expected := []int{64, 56, 48, 40, 32, 24, 16}
	actual := Walk(arr, 7, 7, 0, 7)
	assert.Equal(t, expected, actual)
}

func Test_case5(t *testing.T) {
	arr := [][]int{
		{1, 2, 3, 4, 5, 6, 7, 8},
		{9, 10, 11, 12, 13, 14, 15, 16},
		{17, 18, 19, 20, 21, 22, 23, 24},
		{25, 26, 27, 28, 29, 30, 31, 32},
		{33, 34, 35, 36, 37, 38, 39, 40},
		{41, 42, 43, 44, 45, 46, 47, 48},
		{49, 50, 51, 52, 53, 54, 55, 56},
		{57, 58, 59, 60, 61, 62, 63, 64},
	}

	expected := []int{8, 16, 24, 32, 40, 48, 56}
	actual := Walk(arr, 0, 7, 7, 7)
	assert.Equal(t, expected, actual)
}

func Test_case6(t *testing.T) {
	arr := [][]int{
		{1, 2, 3, 4, 5, 6, 7, 8},
		{9, 10, 11, 12, 13, 14, 15, 16},
		{17, 18, 19, 20, 21, 22, 23, 24},
		{25, 26, 27, 28, 29, 30, 31, 32},
		{33, 34, 35, 36, 37, 38, 39, 40},
		{41, 42, 43, 44, 45, 46, 47, 48},
		{49, 50, 51, 52, 53, 54, 55, 56},
		{57, 58, 59, 60, 61, 62, 63, 64},
	}

	expected := []int{57, 49, 41, 33, 25, 17}
	actual := Walk(arr, 7, 0, 1, 0)
	assert.Equal(t, expected, actual)
}

func Test_case7(t *testing.T) {
	// Create a 5x5 array of ints
	arr := [][]int{
		{1, 2, 3, 4, 5},
		{6, 7, 8, 9, 10},
		{11, 12, 13, 14, 15},
		{16, 17, 18, 19, 20},
		{21, 22, 23, 24, 25},
	}

	spiralArgument := SpiralArgument{
		Direction: CounterClockwiseFrom00,
	}

	expected := []int{
		1, 6, 11, 16, // [0, 0] -> [4, 0]
		21, 22, 23, 24, // [4, 0] -> [4, 4]
		25, 20, 15, 10, // [4, 4] -> [0, 4]
		5, 4, 3, // [0, 4] -> [0, 1]
		2, 7, 12, // [0, 1] -> [3, 1]
		17, 18, // [3, 1] -> [3, 3]
		19, 14, // [3, 3] -> [1, 3]
		9,  // [1, 3] -> [1, 2]
		8,  // [1, 2] -> [2, 2]
		13, // [2, 2] -> [2, 2]
	}

	actual := SpiralOrder(arr, spiralArgument)
	assert.Equal(t, expected, actual)
}

func Test_case8(t *testing.T) {
	arr := [][]int{
		{1, 2, 3, 4, 5, 6, 7, 8},
		{9, 10, 11, 12, 13, 14, 15, 16},
		{17, 18, 19, 20, 21, 22, 23, 24},
		{25, 26, 27, 28, 29, 30, 31, 32},
		{33, 34, 35, 36, 37, 38, 39, 40},
		{41, 42, 43, 44, 45, 46, 47, 48},
		{49, 50, 51, 52, 53, 54, 55, 56},
		{57, 58, 59, 60, 61, 62, 63, 64},
	}

	spiralArgument := SpiralArgument{
		Direction: CounterClockwiseFrom00,
	}

	expected := []int{
		1, 9, 17, 25, 33, 41, 49, // [0, 0] -> [7, 0]
		57, 58, 59, 60, 61, 62, 63, // [7, 0] -> [7, 7]
		64, 56, 48, 40, 32, 24, 16, // [7, 7] -> [0, 7]
		8, 7, 6, 5, 4, 3, // [0, 7] -> [0, 1]
		2, 10, 18, 26, 34, 42, // [0, 1] -> [6, 1]
		50, 51, 52, 53, 54, // [6, 1] -> [6, 6]
		55, 47, 39, 31, 23, // [6, 6] -> [1, 6]
		15, 14, 13, 12, // [1, 6] -> [1, 2]
		11, 19, 27, 35, // [1, 2] -> [5, 2]
		43, 44, 45, // [5, 2] -> [5, 5]
		46, 38, 30, // [5, 5] -> [2, 5]
		22, 21, // [2, 5] -> [2, 3]
		20, 28, // [2, 3] -> [4, 3]
		36, // [4, 3] -> [4, 4]
		37, // [4, 4] -> [3, 4]
		29, // [3, 4] -> [3, 4]
	}

	actual := SpiralOrder(arr, spiralArgument)
	assert.Equal(t, expected, actual)
}

/*
// [0, 0] -> [4, 0]
// [4, 0] -> [4, 4]
// [4, 4] -> [0, 4]
// [0, 4] -> [0, 1]
// [0, 1] -> [3, 1]
// [3, 1] -> [3, 3]
// [3, 3] -> [1, 3]
// [1, 3] -> [1, 2]
// [1, 2] -> [2, 2]
// [2, 2] -> [2, 2]

left:= 0
right := len(array) -1

i := 0 //0
left++ // 0
right -- // 4

i := 1 // 1
left++ // 1
right-- // 3

i := 2
left++ // 2
right-- // 2

0 4 1 3 2


// [0, 0] -> [7, 0]
// [7, 0] -> [7, 7]
// [7, 7] -> [0, 7]
// [0, 7] -> [0, 1]
// [0, 1] -> [6, 1]
// [6, 1] -> [6, 6]
// [6, 6] -> [1, 6]
// [1, 6] -> [1, 2]
// [1, 2] -> [5, 2]
// [5, 2] -> [5, 5]
// [5, 5] -> [2, 5]
// [2, 5] -> [2, 3]
// [2, 3] -> [4, 3]
// [4, 3] -> [4, 4]
// [4, 4] -> [3, 4]
// [3, 4] -> [3, 4]

*/

func Test_case9(t *testing.T) {
	// Create a 5x5 array of ints
	arr := [][]int{
		{1, 2, 3, 4, 5},
		{6, 7, 8, 9, 10},
		{11, 12, 13, 14, 15},
		{16, 17, 18, 19, 20},
		{21, 22, 23, 24, 25},
	}

	spiralArgument := SpiralArgument{
		Direction: ClockwiseFrom0N,
	}

	expected := []int{
		5, 10, 15, 20,
		25, 24, 23, 22,
		21, 16, 11, 6,
		1, 2, 3,
		4, 9, 14,
		19, 18,
		17, 12,
		7,
		8,
		13,
	}

	actual := SpiralOrder(arr, spiralArgument)
	assert.Equal(t, expected, actual)
}

func Test_case10(t *testing.T) {
	// Create a 5x5 array of ints
	arr := [][]int{
		{1, 2, 3, 4, 5},
		{6, 7, 8, 9, 10},
		{11, 12, 13, 14, 15},
		{16, 17, 18, 19, 20},
		{21, 22, 23, 24, 25},
	}

	spiralArgument := SpiralArgument{
		Direction: CounterClockwiseFrom0N,
	}

	expected := []int{
		5, 4, 3, 2,
		1, 6, 11, 16,
		21, 22, 23, 24,
		25, 20, 15,
		10, 9, 8,
		7, 12,
		17, 18,
		19,
		14,
		13,
	}

	actual := SpiralOrder(arr, spiralArgument)
	assert.Equal(t, expected, actual)
}

func Test_case11(t *testing.T) {
	// Create a 5x5 array of ints
	arr := [][]int{
		{1, 2, 3, 4, 5},
		{6, 7, 8, 9, 10},
		{11, 12, 13, 14, 15},
		{16, 17, 18, 19, 20},
		{21, 22, 23, 24, 25},
	}

	spiralArgument := SpiralArgument{
		Direction: ClockwiseFromN0,
	}

	expected := []int{
		21, 16, 11, 6,
		1, 2, 3, 4,
		5, 10, 15, 20,
		25, 24, 23,
		22, 17, 12,
		7, 8,
		9, 14,
		19,
		18,
		13,
	}

	actual := SpiralOrder(arr, spiralArgument)
	assert.Equal(t, expected, actual)
}

func Test_case12(t *testing.T) {
	// Create a 5x5 array of ints
	arr := [][]int{
		{1, 2, 3, 4, 5},
		{6, 7, 8, 9, 10},
		{11, 12, 13, 14, 15},
		{16, 17, 18, 19, 20},
		{21, 22, 23, 24, 25},
	}

	spiralArgument := SpiralArgument{
		Direction: CounterClockwiseFromN0,
	}

	expected := []int{
		21, 22, 23, 24,
		25, 20, 15, 10,
		5, 4, 3, 2,
		1, 6, 11,
		16, 17, 18,
		19, 14,
		9, 8,
		7,
		12,
		13,
	}

	actual := SpiralOrder(arr, spiralArgument)
	assert.Equal(t, expected, actual)
}

func Test_case13(t *testing.T) {
	// Create a 5x5 array of ints
	arr := [][]int{
		{1, 2, 3, 4, 5},
		{6, 7, 8, 9, 10},
		{11, 12, 13, 14, 15},
		{16, 17, 18, 19, 20},
		{21, 22, 23, 24, 25},
	}

	spiralArgument := SpiralArgument{
		Direction: ClockwiseFromNN,
	}

	expected := []int{
		25, 24, 23, 22,
		21, 16, 11, 6,
		1, 2, 3, 4,
		5, 10, 15,
		20, 19, 18,
		17, 12,
		7, 8,
		9,
		14,
		13,
	}

	actual := SpiralOrder(arr, spiralArgument)
	assert.Equal(t, expected, actual)
}

func Test_case14(t *testing.T) {
	// Create a 5x5 array of ints
	arr := [][]int{
		{1, 2, 3, 4, 5},
		{6, 7, 8, 9, 10},
		{11, 12, 13, 14, 15},
		{16, 17, 18, 19, 20},
		{21, 22, 23, 24, 25},
	}

	spiralArgument := SpiralArgument{
		Direction: ClockwiseFromNN,
	}

	expected := []int{
		25, 24, 23, 22,
		21, 16, 11, 6,
		1, 2, 3, 4,
		5, 10, 15,
		20, 19, 18,
		17, 12,
		7, 8,
		9,
		14,
		13,
	}

	actual := SpiralOrder(arr, spiralArgument)
	assert.Equal(t, expected, actual)
}

func Test_case15(t *testing.T) {
	// Create a 5x5 array of ints
	arr := [][]int{
		{1, 2, 3, 4, 5},
		{6, 7, 8, 9, 10},
		{11, 12, 13, 14, 15},
		{16, 17, 18, 19, 20},
		{21, 22, 23, 24, 25},
	}

	spiralArgument := SpiralArgument{
		Direction: CounterClockwiseFromNN,
	}

	expected := []int{
		25, 20, 15, 10,
		5, 4, 3, 2,
		1, 6, 11, 16,
		21, 22, 23,
		24, 19, 14,
		9, 8,
		7, 12,
		17,
		18,
		13,
	}

	actual := SpiralOrder(arr, spiralArgument)
	assert.Equal(t, expected, actual)
}
