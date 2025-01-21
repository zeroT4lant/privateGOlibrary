package main

import (
	"fmt"
	"sync"
)

func merge(chs ...<-chan int) <-chan int {
	resChan := make(chan int, 1)

	wg := &sync.WaitGroup{}

	for _, ch := range chs {
		wg.Add(1)
		go func(ch <-chan int) {
			defer wg.Done()
			for v := range ch {
				resChan <- v
			}
		}(ch)
	}

	go func() {
		wg.Wait()
		close(resChan)
	}()

	return resChan
}

func main() {
	ch1, ch2, ch3 := make(chan int), make(chan int), make(chan int)

	go func() {
		ch1 <- 1
		ch1 <- 3
		close(ch1)
	}()

	go func() {
		ch2 <- 2
		ch2 <- 4
		close(ch2)
	}()

	go func() {
		ch3 <- 0
		ch3 <- 5
		close(ch3)
	}()

	for ch := range merge(ch1, ch2, ch3) {
		fmt.Println(ch)
	}
}
