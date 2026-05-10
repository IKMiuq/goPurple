package main

import (
	"fmt"
)

func sumPart(arr []int, ch chan int) {
	summ := 0
	for _, num := range arr {
		summ += num
	}
	ch <- summ
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}
	numGoroutines := 3
	code := make(chan int, numGoroutines)
	partSize := len(arr) / numGoroutines
	for i := 0; i < numGoroutines; i++ {
		start := i * partSize
		end := start + partSize
		go sumPart(arr[start:end], code)
	}

	totalSum := 0
	for i := 0; i < numGoroutines; i++ {
		totalSum += <-code
	}

	fmt.Println("Total sum: ", totalSum)
}
