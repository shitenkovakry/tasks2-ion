package main

func FindTheCartesianProductOfTwoMatrices(arrayOne [][]string, arrayTwo [][]string) [][]string {
	result := [][]string{}

	for _, rowOne := range arrayOne {
		for _, rowTwo := range arrayTwo {

			newRow := make([]string, len(rowOne)+len(rowTwo))
			copy(newRow[:len(rowOne)], rowOne)
			copy(newRow[len(rowOne):], rowTwo)
			result = append(result, newRow)
		}
	}

	return result
}
