package concurrency

import "math/rand/v2"

func CreateSlice(out chan<- int) {
	numbers := []int{}
	for range 10 {
		numbers = append(numbers, rand.IntN(100))
	}
	for _, n := range numbers {
		out <- n
	}
	close(out)
}

func PowNumber(number int, out chan<- int) {
	out <- number * number
	close(out)
}
