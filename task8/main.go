package main

func FindElementsWhichAreAtTheBottomOfTheMatrix(array [][]int) []int {
	arrayWithResult := []int{}
	lenOfArray := len(array)

	for indexI := lenOfArray - 1; indexI >= lenOfArray/2; indexI-- {
		row := array[indexI]

		if indexI == lenOfArray-1 {
			for indexJ := 0; indexJ < len(row); indexJ++ {
				arrayWithResult = append(arrayWithResult, row[indexJ])
			}
		}

	}

	return arrayWithResult
}
