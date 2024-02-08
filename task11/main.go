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
