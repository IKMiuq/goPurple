package main

import (
	"fmt"
	"math/rand"
	"sync"
)

func randCh(Ch chan int) {
	Ch <- rand.Intn(101)
}

func quaCh(num int, Ch chan int) {
	fmt.Printf("Number %d\n", num)
	Ch <- (num * num)
}

func closeWhenDone(wg *sync.WaitGroup, ch chan int) {
	go func() {
		wg.Wait()
		close(ch)
	}()
}

func main() {
	var waitGroup sync.WaitGroup
	var waitGroup1 sync.WaitGroup
	randChannel := make(chan int)
	quaChannel := make(chan int)

	for i := 0; i < 5; i++ {
		waitGroup.Add(1)
		go func() {
			defer waitGroup.Done()
			randCh(randChannel)
		}()
	}
	closeWhenDone(&waitGroup, randChannel)

	for num := range randChannel {
		waitGroup1.Add(1)
		go func(num int) {
			defer waitGroup1.Done()
			quaCh(num, quaChannel)
		}(num)
	}
	closeWhenDone(&waitGroup1, quaChannel)

	for num := range quaChannel {
		fmt.Printf("Number*Number %d\n", num)
	}
}
