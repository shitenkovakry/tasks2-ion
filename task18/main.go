package main

func FindPerimetr(array [][]string) [][]string {
	arrayWithResult := [][]string{}
	lenOfArray := len(array)

	arrayWithResult = append(arrayWithResult, array[0])

	for indexI := 1; indexI < lenOfArray-1; indexI++ {
		row := array[indexI]
		first := row[0]
		last := row[len(row)-1]

		arrayWithResult = append(arrayWithResult, []string{first, last})
	}

	arrayWithResult = append(arrayWithResult, array[lenOfArray-1])

	return arrayWithResult
}
