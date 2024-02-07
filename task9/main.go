package main

func FindElementsWhichAreInTheLeftSideOfTheMatrix(array [][]int) []int {
	arrayWithResult := []int{}
	lenOfArray := len(array)

	for indexI := 0; indexI < lenOfArray; indexI++ {
		row := array[indexI]

		arrayWithResult = append(arrayWithResult, row[0])
	}

	return arrayWithResult
}
