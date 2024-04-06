package main

import (
	"fmt"
	"time"
)

// в канале используется mutex
func main() {
	nums := []int8{2}
	nums = append(nums, 12)
	fmt.Println(cap(nums))
}

// почти не нужен
// размер буфера nil
func nilChannel() {
	//получаем интовые значение в канал
	//в ниловый канал нельзя ни записать ни прочитать значения
	//вызывается дедлок
	var nilChannel chan int
	fmt.Println(nilChannel)

	//запись в канал
	//nilChannel <- 2

	//чтение из канала
	//value := <-nilChannel

	//close(nilChannel)
}

// Размер буфера в начале 0
// на каждое записывающее действие, должно сразу же происходит чтение и наоборот
// горутина блокируется пока не произойдёт соответствующая операция чтения/записи
// пропускная способность равна ОДНОМУ
// ПИШЕМ И СРАЗУ ЧИТАЕМ

// блокируется до того момента пока не проведём обратную операцию. На запись-чтение и на чтение-должна придти запись
func unbufferedChannel() {
	//не задаём ему len,cap
	unbufferedChannel := make(chan int)

	//направленные каналы
	//unbufferedChannel1 := make(chan <- int)
	//unbufferedChannel2 := make( <- chan int)

	//unbufferedChannel <- 2

	//value := <-unbufferedChannel

	//запись значений
	//Объяснение - пытаемся сначала прочитать значения, которое придёт только через секунду из горутины
	//2)Произойдёт запись
	go func(chanForWritting chan<- int) {
		time.Sleep(time.Second)
		unbufferedChannel <- 1
	}(unbufferedChannel)

	//1)горутина встанет в очередь на чтение
	value := <-unbufferedChannel
	fmt.Println(value)

	//----------------

	//чтение значений
	//2)Как значение запишется, то мы сможем его прочитать
	go func(chanForReading <-chan int) {
		time.Sleep(time.Second)
		value := <-unbufferedChannel
		fmt.Println(value)
	}(unbufferedChannel)

	//1)Становимся в очередь на запись
	unbufferedChannel <- 2

}

// сами указываем размер буфера
func bufferedChannel() {
	bufferedChannel := make(chan int, 2)

	//можем заполнять пока не заполним всю вместимость
	//если превысить вместимость, вылезет дедлок
	//дальше придётся читать значения и освобождать канал
	bufferedChannel <- 2
	bufferedChannel <- 3

	//достаём из канала, не можем достать больше чем лежит в канале
	fmt.Println(<-bufferedChannel)
	fmt.Println(<-bufferedChannel)
}

func forRange() {
	//ФУФЛО СПОСОБ
	//когда прочитает все значения из слайса, будет дефолтные нули вытаскивать
	bufferedChannel := make(chan int, 3)

	numbers := []int{1, 5, 3, 6}

	go func() {
		for _, num := range numbers {
			bufferedChannel <- num
		}
		close(bufferedChannel)
	}()

	for {
		v, ok := <-bufferedChannel
		fmt.Println(v, ok)

		if !ok {
			break
		}
	}

	//СПОСОБ С FOR RANGE
	bufferedChannel2 := make(chan int, 3)

	//записываем значения
	go func() {
		for _, num := range numbers {
			bufferedChannel2 <- num
		}
		//не забываем закрыть, когда дочитали все значения
		close(bufferedChannel2)
	}()

	//удобно читаем значения из буфера
	for v := range bufferedChannel2 {
		fmt.Println("buffered", v)
	}

	//СПОСОБ ДЛЯ НЕБУФИРИЗИРОВАННОГО КАНАЛА
	//пропускная способность равна ОДНОМУ
	//ПИШЕМ И СРАЗУ ЧИТАЕМ
	unbufferedChannel := make(chan int)
	go func() {
		for v := range numbers {
			unbufferedChannel <- v
		}
		//не забываем закрыть, когда дочитали всё
		close(unbufferedChannel)
	}()

	//удобно читаем значения из канала
	for v := range unbufferedChannel {
		fmt.Println("buffered", v)
	}
}
