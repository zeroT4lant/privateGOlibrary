package main

import (
	"fmt"
	"sync"
)

//Паттерн fan-out/fan-in используется для распараллеливания и координации задач (горутин).
//Особенно полезен, когда одну трудоемкую задача можно разделить на более мелкие подзадачи.

func factorial(n int) int {
	if n == 0 {
		return 1
	}

	return n * factorial(n-1)
}

func main() {

	//10 чисел у которых будем вычислять факториал.
	nums := []int{5, 6, 7, 12, 11, 10, 2, 3, 4, 5}

	// создание канала для результатов.
	results := make(chan int)

	var wg sync.WaitGroup

	//fan-out Задача распределяется между несколькими горутинами.
	for _, n := range nums {
		wg.Add(1)
		go func(num int) {
			defer wg.Done()
			results <- factorial(num)
		}(n)
	}

	//После завершения всех горутин закрываем канал и обнуляем счётчик WaitGroup
	go func() {
		wg.Wait()
		close(results)
	}()

	// fan-in: сбор результатов из канала results.
	for result := range results {
		fmt.Println(result)
	}

}
