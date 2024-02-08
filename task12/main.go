package main

func FindElementsWhichAreInTheMiddleOfVerticalLineSideOfTheMatrix(array [][]int) [][]int {
	arrayWithResult := [][]int{}
	lenOfArray := len(array)
	resultRow := []int{}
	resultRow2 := []int{}

	for indexI := 0; indexI < lenOfArray; indexI++ {
		row := array[indexI]
		lenOfRow := len(row)

		if lenOfRow%2 == 0 {
			indexOfDivisionForEven := lenOfRow / 2

			if len(arrayWithResult) == 2 {
				break
			}

			resultRow = append(resultRow, row[indexOfDivisionForEven-1])
			resultRow2 = append(resultRow2, row[indexOfDivisionForEven])
		} else {
			indexOfDivision := lenOfRow / 2

			resultRow = append(resultRow, row[indexOfDivision])
		}
	}

	arrayWithResult = append(arrayWithResult, resultRow)
	arrayWithResult = append(arrayWithResult, resultRow2)

	return arrayWithResult
}
