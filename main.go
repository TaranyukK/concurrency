package main

import (
	"fmt"
	concur "go/concurrency/1-concurrency"
)

func main() {
	generatedNum := make(chan int)
	powedNum := make(chan int)
	go concur.CreateSlice(generatedNum)
	go concur.PowNumber(generatedNum, powedNum)

	for val := range powedNum {
		fmt.Println(val)
	}
}
