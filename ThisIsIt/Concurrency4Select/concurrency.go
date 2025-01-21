package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	bufferedChan := make(chan string, 3)
	bufferedChan <- "first"

	//анализирует и ищет неблокирующие операции и дефолт, если неблокриующих кэйсов не осталось
	//если несколько неблокирующих операций, то выполняет рандомную
	select {
	case str := <-bufferedChan:
		fmt.Println("read", str)
	case bufferedChan <- "second":
		fmt.Println("write", <-bufferedChan, <-bufferedChan)
	}

	unbufChan := make(chan int)

	//записываем значение через секунду
	go func() {
		time.Sleep(time.Second)
		unbufChan <- 1
	}()

	//первый приоритет неблокирующим функциям
	//второй приоритет блокирующим, ждём пока разблокируются
	//в зависимости от времени выполнения, если неблокирующая не успела выполниться раньше таймера, то вызовется кейс с таймером
	//default срабатывает почти мгновенно, если подходящих нет
	select {
	//case bufferedChan <- "third":
	//	fmt.Println("unblocking writting")
	//case val := <-unbufChan:
	//	fmt.Println("blocking reading", val)

	//ожидание
	//по истечению времени если другие кейсы не сработали, то выполним это
	case <-time.After(time.Millisecond * 500):
		fmt.Println("time's up")
		//default:
		//	fmt.Println("default case")
	}

	resultChan := make(chan int)
	timer := time.After(time.Second) //таймер вне цикла
	//выполняем количество операций сколько успеем за определённое время

	//если селект в цикле, то таймер чтобы не обновлялся, надо вынести снаружу
	go func() {
		defer close(resultChan)

		for i := 0; i < 1000; i++ {
			select {
			case <-timer:
				fmt.Println("times up")
				return
			default:
				time.Sleep(time.Nanosecond)
				resultChan <- 1
			}

		}
	}()

	//читаем всё, что успели записать
	for v := range resultChan {
		fmt.Println(v)
	}
}

func gracefulShutdown() {
	sigChan := make(chan os.Signal, 1)

	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	timer := time.After(10 * time.Second)

	select {
	case <-timer:
		fmt.Println("konec vremeni")
	case sig := <-sigChan:
		fmt.Println("Stopped by signal: ", sig)
		return
	}
}

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	ch1 := make(chan string)
// 	ch2 := make(chan string)

// 	go func() {
// 		ch1 <- "one"
// 	}()

// 	go func() {
// 		ch2 <- "two"
// 	}()

// 	for i := 0; i < 2; i++ {
// 		select {
// 		case one := <-ch1:
// 			fmt.Println(one)
// 		case two := <-ch2:
// 			fmt.Println(two)
// 		}
// 	}
// }
