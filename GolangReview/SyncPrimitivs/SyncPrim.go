package main

import (
	"fmt"
	"sync"
)

// К примитивам относят
// Мьютексы RWМютексы
// WaitGroups
// Once
// Атомики - позволяет делать единственную операцию без вмешательства других горутин
func main() {
	//var wg sync.WaitGroup
	//var mu sync.RWMutex
	//
	//m := make(map[int]int)
	//for i := 0; i < 10; i++ {
	//	wg.Add(1)
	//	go func(k int) {
	//		defer wg.Done()
	//		mu.Lock()
	//		m[k] = k * 2
	//		mu.Unlock()
	//	}(i)
	//}
	//wg.Wait()
	//fmt.Println(m)

	wg := sync.WaitGroup{}

	wg.Add(5)
	for i := 0; i < 5; i++ {
		go func(i int) {
			defer wg.Done()
			fmt.Println(i)
		}(i)
	}

	wg.Wait()

}
