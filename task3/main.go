package main

func FindDiagonalWhichLowerMainDiagonal(array [][]int) [][]int {
	arrayWithResult := [][]int{}

	for indexI := 1; indexI < len(array); indexI++ {
		row := array[indexI]
		resultRow := []int{}

		for indexJ := 0; indexJ < indexI; indexJ++ {
			resultRow = append(resultRow, row[indexJ])
		}

		arrayWithResult = append(arrayWithResult, resultRow)
	}

	return arrayWithResult
}
