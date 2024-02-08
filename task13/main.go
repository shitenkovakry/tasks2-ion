package main

func FindElementsWhichAreInTheMiddleSideOfTheMatrix(array [][]int) [][]int {
	arrayWithResult := [][]int{}
	lenOfArray := len(array)

	if lenOfArray%2 == 0 {
		indexOfDivision := lenOfArray / 2

		arrayWithResult = append(arrayWithResult, array[indexOfDivision-1])
		arrayWithResult = append(arrayWithResult, array[indexOfDivision])
	} else {
		indexOfDivisionForOdd := lenOfArray / 2

		arrayWithResult = append(arrayWithResult, array[indexOfDivisionForOdd])
		arrayWithResult = append(arrayWithResult, []int{})
	}

	return arrayWithResult
}

func FindElementsWhichAreInTheMiddleOfVerticalLineSideOfTheMatrix(array [][]int) [][]int {
	arrayWithResult := [][]int{}
	lenOfArray := len(array)
	resultRow := []int{}
	resultRow2 := []int{}

	for indexI := 0; indexI < lenOfArray; indexI++ {
		row := array[indexI]
		lenOfRow := len(row)

		if lenOfRow == 0 {
			continue
		}

		if lenOfRow%2 == 0 {
			indexOfDivisionForEven := lenOfRow / 2

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

func Result(array [][]int) [][]int {
	middleResult := FindElementsWhichAreInTheMiddleSideOfTheMatrix(array)
	finalResult := FindElementsWhichAreInTheMiddleOfVerticalLineSideOfTheMatrix(middleResult)

	return finalResult
}
