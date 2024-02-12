package main

type Direction int

const (
	Z Direction = iota + 1
	N
	S
)

type Argument struct {
	Direction Direction
}

func DrawingResult(array [][]int, argument Argument) []int {
	if argument.Direction == Z {
		return DrawZ(array)
	}

	if argument.Direction == N {
		return DrawN(array)
	}

	if argument.Direction == S {
		return DrawS(array)
	}

	return []int{}
}

func DrawZ(array [][]int) []int {
	lenOfArray := len(array)
	result := []int{}

	for index := 0; index < lenOfArray-1; index++ {
		result = append(result, array[0][index])
	}

	for index := 0; index < lenOfArray; index++ {
		left := index
		right := lenOfArray - 1 - index

		result = append(result, array[left][right])
	}

	for index := 1; index < lenOfArray; index++ {
		result = append(result, array[lenOfArray-1][index])
	}

	return result
}

func DrawN(array [][]int) []int {
	lenOfArray := len(array)
	result := []int{}

	for index := lenOfArray - 1; index >= 0; index-- {
		result = append(result, array[index][0])
	}

	for index := 1; index < lenOfArray-1; index++ {
		result = append(result, array[index][index])
	}

	for index := lenOfArray - 1; index >= 0; index-- {
		result = append(result, array[index][lenOfArray-1])
	}

	return result
}

func DrawS(array [][]int) []int {
	rows := len(array)
	cols := len(array[0])

	result := make([]int, 0, rows*cols)

	for indexI := 0; indexI < rows; indexI++ {
		if indexI%2 == 0 {
			for indexJ := cols - 1; indexJ >= 0; indexJ-- {
				result = append(result, array[indexI][indexJ])
			}
		} else {
			for indexL := 0; indexL < cols; indexL++ {
				result = append(result, array[indexI][indexL])
			}
		}
	}

	return result
}
