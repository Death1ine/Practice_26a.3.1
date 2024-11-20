package filters

// FilterNegativeNumbers фильтрует отрицательные числа
func FilterNegativeNumbers(input <-chan int) <-chan int {
	output := make(chan int)
	go func() {
		defer close(output)
		for num := range input {
			if num >= 0 {
				output <- num
			}
		}
	}()
	return output
}

// FilterNotMultipleOfThree фильтрует числа, не кратные 3 (исключая 0)
func FilterNotMultipleOfThree(input <-chan int) <-chan int {
	output := make(chan int)
	go func() {
		defer close(output)
		for num := range input {
			if num != 0 && num%3 == 0 {
				output <- num
			}
		}
	}()
	return output
}
