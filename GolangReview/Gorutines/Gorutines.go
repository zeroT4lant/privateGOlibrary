package main

import (
	"context"
	"fmt"
	"time"
)

// Горутина - функция, которая может выполняться параллельно с другими функциями.
// Весит 2кб
// Лёгкие и проще переключаются в отличии от потоков
// Управляется планировщиком Go
// Общается с другими горутинами через каналы, которые обеспечивают синхронизацию и безопасность данных.
// GOMAXPROCS - это параметр, который определяет, сколько ядер ЦП используется для одновременного выполнения горутин.
// Остановить горутину можно с помощью контекста, канала или таймаута.
func main() {
	//for i := 0; i < 5; i++ {
	//	go func(int k) {
	//		fmt.Println(k)
	//	}(i)
	//}

	// создаем контекст с отменой
	ctx, cancel := context.WithCancel(context.Background())
	// запускаем горутину с этим контекстом
	go worker(ctx)
	// ждем 3 секунды
	time.Sleep(3 * time.Second)
	// отменяем контекст
	cancel()
	// ждем еще 2 секунды
	time.Sleep(2 * time.Second)

	// создаем канал для передачи сигнала остановки
	stop := make(chan bool)
	// запускаем горутину с этим каналом
	go worker2(stop)
	// ждем 3 секунды
	time.Sleep(3 * time.Second)
	// отправляем сигнал остановки в канал
	stop <- true
	// ждем еще 2 секунды
	time.Sleep(2 * time.Second)
}

func worker(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			// контекст отменен, выходим из горутины
			fmt.Println("worker stopped")
			return
		default:
			// продолжаем работать
			fmt.Println("worker working")
			time.Sleep(time.Second)
		}
	}
}

func worker2(stop chan bool) {
	for {
		select {
		case <-stop:
			// получили сигнал остановки, выходим из горутины
			fmt.Println("worker stopped")
			return
		default:
			// продолжаем работать
			fmt.Println("worker working")
			time.Sleep(time.Second)
		}
	}
}
