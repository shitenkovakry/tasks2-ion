package main

func FindElementsWhichareAtTheTopOfTheMatrix(array [][]int) []int {
	arrayWithResult := []int{}
	lenOfArray := len(array)

	for indexI := 0; indexI < lenOfArray; indexI++ {
		row := array[indexI]

		if indexI == 0 {
			for indexJ := 0; indexJ < len(row); indexJ++ {
				arrayWithResult = append(arrayWithResult, row[indexJ])
			}
		}

	}

	return arrayWithResult
}
