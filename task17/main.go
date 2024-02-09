package main

func FindInsidePerimetr(array [][]string) [][]string {
	arrayWithResult := [][]string{}
	lenOfArray := len(array)

	for indexI := 1; indexI < lenOfArray-1; indexI++ {
		row := array[indexI]
		resultRow := []string{}

		for indexJ := 1; indexJ < len(row)-1; indexJ++ {
			resultRow = append(resultRow, row[indexJ])
		}

		arrayWithResult = append(arrayWithResult, resultRow)
	}

	return arrayWithResult
}
