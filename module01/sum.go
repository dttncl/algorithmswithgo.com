package module01

// Sum will sum up all of the numbers passed
// in and return the result
func Sum(numbers []int) int {

	//version1
	//sum := 0
	//for _, val := range numbers {
	//	sum += val
	//}
	//return sum

	//version2
	if len(numbers) == 0 {
		return 0 // if len of array is 0, return 0
	}

	return numbers[0] + Sum(numbers[1:]) // else, take first element and add with Sum(remaining elements)
}
