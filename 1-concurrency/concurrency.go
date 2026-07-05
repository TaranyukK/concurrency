package concurrency

import "math/rand/v2"

func CreateSlice(out chan<- int) {
	numbers := []int{}
	for range 10 {
		numbers = append(numbers, rand.IntN(101))
	}
	for _, n := range numbers {
		out <- n
	}
	close(out)
}

func PowNumber(in <-chan int, out chan<- int) {
	for number := range in {
		out <- number * number
	}
	close(out)
}
