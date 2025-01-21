package main

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"time"
)

// context может
// 1)Хранить значения
// 2)Сообщать о завершении
func main() {
	baseKnowledge()
}

func baseKnowledge() {
	ctx := context.Background()
	fmt.Println(ctx)

	todo := context.TODO()
	fmt.Println(todo)

	withValue := context.WithValue(ctx, "name", "vasya")
	fmt.Println(withValue.Value("name"))

	//создали контекст, который может завершаться
	//вручную отменяем
	withCancel, cancel := context.WithCancel(ctx)
	fmt.Println(withCancel.Err())
	//завершили его, вылезает сообщение в терминале "context canceled"
	cancel()
	fmt.Println(withCancel.Err())

	withDeadline, cancel := context.WithDeadline(ctx, time.Now().Add(time.Second*3))
	//как только функция завершиться, контекст будет завершён
	defer cancel()
	fmt.Println(withDeadline.Deadline())
	//ошибки пока нет, будет <nil>
	fmt.Println(withDeadline.Err())
	//по окончанию приходит значение из канала
	//подождали дедлайн
	fmt.Println(<-withDeadline.Done())

	//через две секунды канал отменяется при помощи cancel()
	withTimeout, cancel := context.WithTimeout(ctx, time.Second*2)
	defer cancel()
	//и так же приходит значение
	fmt.Println(<-withTimeout.Done())

}

func workerPool() {

	ctx, cancel := context.WithCancel(context.Background())
	//ctx, cancel := context.WithTimeout(context.Background(),time.Millisecond*20)

	defer cancel()

	wg := sync.WaitGroup{}
	numbersToProcces, processedNumbers := make(chan int, 5), make(chan int, 5)

	for i := 0; i < runtime.NumCPU(); i++ {
		//добавляем одну задачу
		wg.Add(1)
		go func() {
			//после завершения задачи, уменьшаем пул общих задачек
			defer wg.Done()
			worker(ctx, numbersToProcces, processedNumbers)

		}()
	}

	go func() {
		for i := 0; i < 1000; i++ {
			numbersToProcces <- i
		}
		close(numbersToProcces)
	}()
}

func worker(ctx context.Context, toProcces <-chan int, proccesed chan<- int) {
	for {
		select {
		case <-ctx.Done():
			return
		case value, ok := <-toProcces:
			if !ok {
				time.Sleep(time.Millisecond)
				proccesed <- value * value
			}
		}

	}
}

//context.Background()
// context.TODO()
// - **`context.WithCancel(parent)`**
//     - Создает контекст с возможностью отмены.
//     - Возвращает новый контекст и функцию отмены (`cancel()`), которую можно вызывать для завершения операций.
// - **`context.WithDeadline(parent, deadline time.Time)`**
//     - Создает контекст с заданным дедлайном. Контекст автоматически отменяется, когда дедлайн истекает.
// - **`context.WithTimeout(parent, timeout time.Duration)`**
//     - То же, что и `WithDeadline`, но дедлайн вычисляется как текущее время + `timeout`.
// - **`context.WithValue(parent, key, value)`**
//     - Создает контекст с пользовательскими данными (ключ-значение).
//     - Используется для передачи метаданных, таких как ID пользователя или токен.
// ### **Зачем нужны контексты?**

// 1. **Управление временем жизни операций:**

//     - Например, вы хотите прервать долгую операцию, если пользователь закрыл соединение или истек таймаут.
//     - Контексты помогают избежать "утечек" горутин, автоматически отменяя зависимые операции.
// - И передача значений
// Контексты — мощный инструмент для управления временем жизни операций и согласованной работы между горутинами. Без них можно использовать каналы и мьютексы, но это усложняет код. Для остановки HTTP-сервера рекомендуется использовать метод `Shutdown` в сочетании с `context`
