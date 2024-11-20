package filters

import (
	"log"
)

// FilterNegativeNumbers фильтрует отрицательные числа
func FilterNegativeNumbers(input <-chan int) <-chan int {
	output := make(chan int)
	go func() {
		defer close(output)
		for num := range input {
			if num >= 0 {
				log.Printf("Число %d прошло фильтрацию на отрицательные числа\n", num)
				output <- num
			} else {
				log.Printf("Число %d исключено, так как оно отрицательное\n", num)
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
				log.Printf("Число %d прошло фильтрацию на кратность 3\n", num)
				output <- num
			} else {
				log.Printf("Число %d исключено, так как оно не кратно 3 или равно 0\n", num)
			}
		}
	}()
	return output
}
