package main

import (
	"fmt"
	"sync"
)

//10 конкурентных запросов
//вывести в консоль 10 статусов

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}
	var wg sync.WaitGroup
	code := make(chan int)
	for i := 1; i < 4; i++ {
		wg.Add(1)
		go func() {
			getMath(arr, i, code)
			wg.Done()
		}()
	}
	go func() {
		wg.Wait()
		close(code)
	}()
	var summ int = 0
	for res := range code {
		summ += res
	}
	fmt.Printf("Resut summ:= %d\n", summ)
}

func getMath(arr []int, num int, codeCh chan int) {
	var summ int = 0
	for i := (num - 1) * 4; i < num*4; i++ {
		summ += arr[i]
	}
	codeCh <- summ
}
