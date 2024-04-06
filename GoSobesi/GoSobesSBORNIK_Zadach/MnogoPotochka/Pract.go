package main

import (
	"fmt"
	"time"
)

func main() {
	//ch := make(chan int, 3)
	//wg := &sync.WaitGroup{}
	//wg.Add(3)
	//for i := 0; i < 3; i++ {
	//	go func(v int) {
	//		defer wg.Done()
	//		ch <- v * v
	//	}(i)
	//}
	//
	//wg.Wait()
	//close(ch)
	//
	//var sum int
	//for v := range ch {
	//	sum += v
	//}
	//fmt.Printf("result: %d\n", sum)
	//
	//foo1()

	x := make(map[int]int, 1)
	go func() { x[1] = 2 }()
	go func() { x[1] = 7 }()
	go func() { x[1] = 10 }()
	time.Sleep(100 * time.Millisecond)
	fmt.Println("x[1] =", x[1])

}

func foo1() {
	ch := make(chan bool, 2)
	ch <- true
	for v := range ch {
		go func() {
			fmt.Println(v)
		}()
	}

	ch <- true
	time.Sleep(time.Second * 5)
}
