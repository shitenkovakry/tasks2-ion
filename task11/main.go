package main

func FindElementsWhichAreInTheMiddleSideOfTheMatrix(array [][]int) [][]int {
	arrayWithResult := [][]int{}
	lenOfArray := len(array)

	if lenOfArray%2 == 0 {
		indexOfDivision := lenOfArray / 2

		for indexI := indexOfDivision; indexI < lenOfArray; indexI-- {
			if len(arrayWithResult) == 2 {
				break
			}

			row := array[indexI]
			resultRow := []int{}

			for indexJ := 0; indexJ < len(row); indexJ++ {
				resultRow = append(resultRow, row[indexJ])

			}

			arrayWithResult = append(arrayWithResult, resultRow)
		}

	} else {
		indexOfDivisionForOdd := lenOfArray / 2

		for indexI := indexOfDivisionForOdd; indexI < lenOfArray; indexI-- {
			if len(arrayWithResult) == 1 {
				arrayWithResult = append(arrayWithResult, []int{})
				break
			}

			row := array[indexI]
			resultRow := []int{}

			for indexJ := 0; indexJ < len(row); indexJ++ {
				resultRow = append(resultRow, row[indexJ])

			}

			arrayWithResult = append(arrayWithResult, resultRow)
		}
	}

	return arrayWithResult
}
