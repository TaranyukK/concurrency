package main

import (
	"fmt"
	concur "go/concurrency/1-concurrency"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	generatedNum := make(chan int)
	powedNum := make(chan int)
	wg.Go(func() {
		concur.CreateSlice(generatedNum)
	})
	for val := range generatedNum {
		wg.Go(func() {
			concur.PowNumber(val, powedNum)
		})
	}
	for val := range powedNum {
		fmt.Println(val)
	}
}
