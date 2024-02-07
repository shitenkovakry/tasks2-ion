package main

func FindElementsWhichHigherSideDiagonal(array [][]int) [][]int {
	arrayWithResult := [][]int{}
	lenOfArray := len(array)

	for indexI := 0; indexI < lenOfArray-1; indexI++ {
		row := array[indexI]
		resultRow := []int{}

		for indexJ := 0; indexJ < len(row)-indexI-1; indexJ++ {
			resultRow = append(resultRow, row[indexJ])
		}

		arrayWithResult = append(arrayWithResult, resultRow)
	}

	return arrayWithResult
}
