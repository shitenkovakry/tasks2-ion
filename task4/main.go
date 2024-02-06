package main

func FindSideDiagonal(array [][]int) []int {
	arrayWithResult := []int{}

	for index := 0; index < len(array); index++ {
		row := array[index]

		if index < len(row) {
			value := row[len(row)-index-1]

			arrayWithResult = append(arrayWithResult, value)
		} else {
			arrayWithResult = append(arrayWithResult, 0)
		}
	}

	return arrayWithResult
}
