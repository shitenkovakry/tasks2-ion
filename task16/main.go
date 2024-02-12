package main

func Result(array [][]string) [][]string {
	result := [][]string{}

	firstResultRow := []string{}
	lastResultRow := []string{}

	for index := 1; index < len(array)-1; index++ {
		firstRow := array[0]
		firstResultRow = append(firstResultRow, firstRow[index])

		lastRow := array[len(array)-1]
		lastResultRow = append(lastResultRow, lastRow[index])
	}

	result = append(result, firstResultRow)

	for index := 1; index < len(array)-1; index++ {
		row := array[index]

		first := row[0]
		last := row[len(row)-1]

		resultRow := []string{first, last}

		result = append(result, resultRow)

	}

	result = append(result, lastResultRow)

	return result
}
