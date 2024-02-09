package main

func FindCornersOfRow(row []int) []int {
	result := []int{}

	if len(row)%2 == 0 {
		result = append(result, row[0], row[len(row)-1])
	} else {
		result = append(result, row[0], row[len(row)-1])
	}

	return result
}

func Result(array [][]int) [][]int {
	arrayWithResult := [][]int{}
	lenOfArray := len(array)
	top := array[0]
	arrayWithResult = append(arrayWithResult, FindCornersOfRow(top))

	bottom := array[lenOfArray-1]
	arrayWithResult = append(arrayWithResult, FindCornersOfRow(bottom))

	return arrayWithResult
}
