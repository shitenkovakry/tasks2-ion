package main

func FindElementsWhichLowerSideDiagonal(array [][]int) [][]int {
	arrayWithResult := [][]int{}
	lenOfArray := len(array)

	for indexI := 1; indexI < lenOfArray; indexI++ {
		row := array[indexI]
		resultRow := []int{}

		for indexJ := len(row) - 1; indexJ >= lenOfArray-indexI; indexJ-- {
			resultRow = append(resultRow, row[indexJ])
		}

		arrayWithResult = append(arrayWithResult, resultRow)
	}

	return arrayWithResult
}
