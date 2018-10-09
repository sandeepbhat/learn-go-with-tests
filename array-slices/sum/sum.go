package sum

// Sum - Return sum of given numbers
func Sum(numbers []int) int {
	var sumOfNumbers int

	for _, number := range numbers {
		sumOfNumbers += number
	}

	return sumOfNumbers
}

// SumAll - Return sum numbers in each of set
func SumAll(setOfNumbers ...[]int) []int {
	var allSums []int

	for _, numbers := range setOfNumbers {
		allSums = append(allSums, Sum(numbers))
	}

	return allSums
}

// SumOfAllTails - Return sum of tail as an array for a given set of numbers
func SumOfAllTails(setOfNumbers ...[]int) []int {
	var allSums []int

	for _, numbers := range setOfNumbers {
		if len(numbers) == 0 {
			allSums = append(allSums, 0)
		} else {
			allSums = append(allSums, Sum(numbers[1:]))
		}
	}

	return allSums
}
