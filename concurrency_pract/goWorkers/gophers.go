package main

import "fmt"

func main() {
	//если на небуфферизованный канал подавать значение в горутине и пытаться сразу читать его
	//то вызовется дедлок
	//с небуфферизованными - РАБОТАТЬ ТОЛЬКО С ГОРУТИНОЙ
	//str := make(chan string)
	//нужно чтобы горутина сразу отдавала значение и принимала в другом потоке-горутине
	//go func() {
	//	str <- "text"
	//}()
	//fmt.Println(<-str)

	//**********************
	//ещё способ как работать с небуфферизованными
	//суть в том, что поток main и горутина, должны отработать одновременно, что бы произвести передачу-получение значения из канала
	//main встаёт в ожидание и ждёт операцию из горутины

	//УСПЕВАЕТ ЧЕРЕЗ РАЗ
	//str := make(chan string)
	//
	//go func(chanForReading <-chan string) {
	//	time.Sleep(time.Second * 1)
	//	fmt.Println(<-str)
	//}(str)
	//
	//str <- "text"

	//---------------------------------------------

	//в буфферизованном делаем как хотим
	//buf := make(chan string, 1)
	//buf <- "text"
	//fmt.Println(<-buf)

	//---------------------------------------------

	channel := make(chan int)
	go throwingStarts(channel)

	//for v := range channel {
	//	fmt.Println(v)
	//}
	//либо так - читаем пока не закончатся значения
	//for {
	//	mes, open := <-channel
	//	if !open {
	//		break
	//	}
	//	fmt.Println(mes)
	//}

	//----------------------------------
	//SELECT`s

	ninja1, ninja2 := make(chan string), make(chan string)

	go captainElect(ninja1, "Ninja 1")
	go captainElect(ninja2, "Ninja 2")

	//делает один выбор рандомно среди равноправных кейсов
	select {
	case message := <-ninja1:
		fmt.Println(message)
	case message := <-ninja2:
		fmt.Println(message)
	}

}

func captainElect(ninja chan string, message string) {
	ninja <- message
}

func throwingStarts(channel chan int) {
	go func(channel chan int) {
		for i := 0; i < 3; i++ {
			channel <- i + 3
		}
		close(channel)
	}(channel)
}
