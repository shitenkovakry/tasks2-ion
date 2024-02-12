package main

func FindTheNaturalUnionOfTwoMatrices(arrayOne [][]string, arrayTwo [][]string) [][]string {
	result := [][]string{}

	for indexI := 0; indexI < len(arrayOne); indexI++ {
		rowOne := arrayOne[indexI]

		firstElementRowOne := rowOne[0]
		if found := FindRowBeginningWith(arrayTwo, firstElementRowOne); found != nil {
			resultRow := []string{}

			resultRow = append(resultRow, rowOne...)
			resultRow = append(resultRow, found...)
			result = append(result, resultRow)
		}
	}

	return result
}

func FindRowBeginningWith(array [][]string, findItem string) []string {
	for indexI := 0; indexI < len(array); indexI++ {
		row := array[indexI]
		firstElement := row[0]

		if firstElement == findItem {
			return row[1:]
		}
	}

	return nil
}
