package main

func FindCenterOfRow(row []int) []int {
	result := []int{}

	if len(row)%2 == 0 {
		centerIndex := len(row) / 2

		result = append(result, row[centerIndex-1])
		result = append(result, row[centerIndex])
	} else {
		centerIndex := len(row) / 2

		result = append(result, row[centerIndex])
	}

	return result
}

func Result(array [][]int) [][]int {
	arrayWithResult := [][]int{}
	lenOfArray := len(array)
	top := array[0]
	arrayWithResult = append(arrayWithResult, FindCenterOfRow(top)) // center of top

	if lenOfArray%2 == 0 {
		centerIndex := lenOfArray / 2

		centerMinusOne := array[centerIndex-1]
		arrayWithResult = append(arrayWithResult, []int{centerMinusOne[0], centerMinusOne[len(centerMinusOne)-1]})

		center := array[centerIndex]
		arrayWithResult = append(arrayWithResult, []int{center[0], center[len(center)-1]})
	} else {
		centerIndex := lenOfArray / 2

		center := array[centerIndex]
		arrayWithResult = append(arrayWithResult, []int{center[0], center[len(center)-1]})
	}

	bottom := array[lenOfArray-1]
	arrayWithResult = append(arrayWithResult, FindCenterOfRow(bottom)) // center of bottom

	return arrayWithResult
}
