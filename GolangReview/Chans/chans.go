package main

import "fmt"

// type chan struct {
// 	mx sync.mutex
// 	buffer []T
// 	readers []Goroutines
// 	writers []Goroutines
//   }

//Про аксиомы каналов посмотри

func main() {

	ch := make(chan int)

	go func() {
		for i := 0; i < 5; i++ {
			ch <- i
		}

	}()

	for n := range ch {
		fmt.Println(n)
	}
}
