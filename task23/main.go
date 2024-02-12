package main

func FindTheIntersectionOfTwoMatrices(arrayOne [][]string, arrayTwo [][]string) [][]string {
	intersection := [][]string{}

	for _, row := range arrayTwo {
		if HasRow(arrayOne, row) {
			intersection = append(intersection, row)
		}
	}

	return intersection
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
