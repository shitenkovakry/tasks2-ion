package main

func FindElementsWhichAreInTheRightSideOfTheMatrix(array [][]int) []int {
	arrayWithResult := []int{}
	lenOfArray := len(array)

	for indexI := 0; indexI < lenOfArray; indexI++ {
		row := array[indexI]

		arrayWithResult = append(arrayWithResult, row[len(row)-1])
	}

	return arrayWithResult
}
