package main

func FindDiagonalWhichHigherMainDiagonal(array [][]int) [][]int {
	arrayWithResult := [][]int{}

	for indexI := 0; indexI < len(array)-1; indexI++ {
		row := array[indexI]
		resultRow := []int{}
		for indexJ := indexI + 1; indexJ < len(row); indexJ++ {
			resultRow = append(resultRow, row[indexJ])
		}

		arrayWithResult = append(arrayWithResult, resultRow)
	}

	return arrayWithResult
}
