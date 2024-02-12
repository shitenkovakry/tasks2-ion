package main

func CombineTwoMatrixIntoOne(arrayOne [][]string, arrayTwo [][]string) [][]string {
	combined := [][]string{}

	for _, row := range arrayOne {
		combined = append(combined, row)
	}

	for _, row := range arrayTwo {
		if !HasRow(combined, row) {
			combined = append(combined, row)
		}
	}

	return combined
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
