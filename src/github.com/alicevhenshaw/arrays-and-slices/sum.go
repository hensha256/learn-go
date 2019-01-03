package arrays

func Sum(numbers [5]int) (sum int) {
	sum = 0
	for _, number := range numbers {
		sum += number
	}
	return
}

func SumSlice(numbers []int) (sum int) {
	sum = 0
	for _, number := range numbers {
		sum += number
	}
	return
}

func SumAll(numbersSlices []([]int)) []int {
	numberOfSlices := len(numbersSlices)
	// creates a slice with a starting capacity of numberOfSlices
	// or you can define sums as a 'var' which allows you to edit its capacity
	sums := make([]int, numberOfSlices)
	
	for index, numberSlice := range numbersSlices {
		sums[index] = SumSlice(numberSlice)
	}
	return sums
}

func SumAllTails(numbersSlices []([]int)) []int {
	numberOfSlices := len(numbersSlices)
	sums := make([]int, numberOfSlices)
	
	for index, numberSlice := range numbersSlices {
		if len(numberSlice) == 0 {
			sums[index] = 0
		} else {
			// removes the first element of the slice
			tail := numberSlice[1:]
			sums[index] = SumSlice(tail)
		}
	}
	return sums
}