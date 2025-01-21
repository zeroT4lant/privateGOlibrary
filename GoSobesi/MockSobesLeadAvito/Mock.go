package main

import (
	"context"
	"fmt"
	"time"
)

//https://www.youtube.com/watch?v=aLGi0I3tJx4

// Задача 1
// a2 a3 будут смотреть на исходный a1, и
func a() {
	a1 := make([]int, 0, 10)
	a1 = append(a1, []int{1, 2, 3, 4, 5}...) //1 2 3 4 5
	a2 := append(a1, 6)                      // 1 2 3 4 5 6
	a3 := append(a1, 7)                      // 1 2 3 4 5 7

	fmt.Println(a1, a2, a3)
	//1 2 3 4 5
	// 1 2 3 4 5 7
	// 1 2 3 4 5 7
}

// Задача 2
// Верни из функции ошибку, не подключая доп.пакеты
func main() {
	//println(handle())

	b()

	//---------------------
	//ctx := context.Background()
	//res, err := getDiscountWithContext(ctx)
	//if err != nil {
	//	fmt.Printf("Что-то пошло не так", err.Error())
	//	return
	//}
	//
	//fmt.Printf("Ваша скидка %v", res)
}

// Должен реализовать метод "Error" этого интерфейса, который вернёт string
func handle() error {
	return &myError{text: "Ошибка"}
}

// Теперь тоже реализует интерфейс "error"
type myError struct {
	text string
}

// Реализовали метод error для получения ошибки в виде строки
// !!!!Чтобы в handle мог вернуть error - надо реализовать метод Error, который возвращает string - текст ошибки
func (m *myError) Error() string {
	return m.text
}

//Интерфейс ошибки
//type error interface {
//	Error() string
//}

// Задача 3
func b() {
	first := []int{10, 20, 30, 40}
	//Слайс указателей с интами
	second := make([]*int, len(first))
	//Переменные создаются и хранятся в какой-либо области памяти, пока операция выполняется
	//Лежат в какой-то области памяти i:=b010 v:=b102
	for i, v := range first {
		//Слайс из областей памяти, одни и те же значения будут
		//При итерации указатель смещается и теперь во всём слайсе указатель на 40
		second[i] = &v
	}
	//Разыменовываем элементы и выводим их
	//Весь слайс будет забит последними значениями. 40
	fmt.Println(*second[0], *second[1]) //40 40 40 40
	fmt.Println(second)
}

// Задача 4
func getDiscount() float64 {
	//discount, _ := http.Get("http:discounts.com/my")

	return 22.8
}

func getDiscountWithContext(ctx context.Context) (float64, error) {
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	ch := make(chan float64)

	go func() {
		ch <- getDiscount()
	}()

	select {
	case <-ctx.Done():
		return 0.0, ctx.Err()
	case res := <-ch:
		return res, nil
	}
}
