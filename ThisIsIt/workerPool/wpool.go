package main

import (
	"fmt"
	"time"
)

func main() {
	//Чтобы использовать наш воркер пул, нам нужно отправить им задание и получить результаты выполнения.
	//Для этого мы делаем 2 канала.
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	//Стартуем 3 воркера, первоначально заблокированных, т.к. еще нет заданий.
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)

	}

	//посылаем 5 заданий и закрываем канал
	for j := 1; j <= 5; j++ {
		jobs <- j

	}
	defer close(jobs)

	//числа над которыми было произведено действие выводятся
	for i := 1; i <= 5; i++ {
		fmt.Println(<-results)
	}

}

// Эти воркеры будут получать задания через канал jobs-отправляет(записывает) значения
// отсылать результаты в results - принимает(читает) значения
func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "started  job", j)
		time.Sleep(time.Second)
		fmt.Println("worker", id, "finished job", j)
		results <- j * 2
	}
}
