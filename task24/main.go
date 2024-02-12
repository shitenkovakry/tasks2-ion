package main

func FindTheDifferenceOfTwoMatrices(arrayOne [][]string, arrayTwo [][]string) [][]string {
	difference := [][]string{}

	for _, row := range arrayOne {
		if !HasRow(arrayTwo, row) {
			difference = append(difference, row)
		}
	}

	return difference
}

func HasRow(array [][]string, target []string) bool {
	for _, row := range array {
		found := true

		for indexItem, item := range row {
			targetItem := target[indexItem]
			if targetItem != item {
				found = false
				break
			}
		}

		if found {
			return true
		}
	}

	return false
}
