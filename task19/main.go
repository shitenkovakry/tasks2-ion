package main

import "log"

type Direction int

const (
	Clockwise Direction = iota + 1
	CounterClockwise
)

type SpiralArgument struct {
	StartI    int
	StartJ    int
	Direction Direction
}

func SpiralOrder(array [][]int, spiralArgument SpiralArgument) []int {
	if spiralArgument.Direction == Clockwise {
		return SpiralClockwise(array)
	}

	if spiralArgument.Direction == CounterClockwise {
		return SpiralCounterClockwise(array)
	}

	return []int{}
}

func SpiralClockwise(array [][]int) []int {
	left := 0
	right := len(array) - 1
	result := []int{}

	for index := 0; index <= len(array)/2; index++ {
		if left > 0 {
			result = append(result, Walk(array, left, left-1, left, right)...)
		} else {
			result = append(result, Walk(array, left, left, left, right)...)
		}

		result = append(result, Walk(array, left, right, right, right)...)

		if right != left {
			result = append(result, Walk(array, right, right, right, left)...)
			result = append(result, Walk(array, right, left, left+1, left)...)
		}

		left++
		right--

		if left > right {
			break
		}
	}

	return result
}

func SpiralCounterClockwise(array [][]int) []int {
	left := 0
	right := len(array) - 1
	result := []int{}

	for index := 0; index <= len(array)/2; index++ {
		if left > 0 {
			result = append(result, Walk(array, left-1, left, right, left)...)
		} else {
			result = append(result, Walk(array, left, left, right, left)...)
		}

		result = append(result, Walk(array, right, left, right, right)...)

		if right != left {
			result = append(result, Walk(array, right, right, left, right)...)
			result = append(result, Walk(array, left, right, left, left+1)...)
		}

		left++
		right--

		if left > right {
			break
		}
	}

	return result
}

func Walk(array [][]int, x1, y1, x2, y2 int) []int {
	log.Printf("[%d, %d] -> [%d, %d]", x1, y1, x2, y2)
	result := []int{}

	if x1 == x2 && y1 < y2 {
		for index := y1; index < y2; index++ {
			result = append(result, array[x1][index])
		}
	}

	if x1 == x2 && y1 > y2 {
		for index := y1; index > y2; index-- {
			result = append(result, array[x1][index])
		}
	}

	if y1 == y2 && x1 > x2 {
		for index := x1; index > x2; index-- {
			result = append(result, array[index][y1])
		}
	}

	if y1 == y2 && x1 < x2 {
		for index := x1; index < x2; index++ {
			result = append(result, array[index][y1])
		}
	}

	if x1 == x2 && y1 == y2 {
		result = append(result, array[x1][y1])
	}

	return result
}
