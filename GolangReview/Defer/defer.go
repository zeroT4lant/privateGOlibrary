package main

import "fmt"

// 3) Defer инициализируется в стеке, вызовы помещаются в стэк

type X struct {
	V int
}

func (x X) S() {
	fmt.Println(x.V)
}

// В памяти сохранится значение 123 при выводе. Так как напрямую указали число и оно сохранилось
func main() {
	//x := X{123}
	//defer x.S()
	//x.V = 456

	//--------------------
	//
	a()

	//b()
}

func a() {
	//В данном случае, анонимная функция, которая откладывается с помощью defer,
	//имеет доступ к переменной tmp из внешнего контекста.

	//Если ничего не указать, то число не постоянное и меняется.
	//Так что меняется на 202.
	tmp := 101
	fmt.Println(tmp)
	defer func() {
		fmt.Println(tmp)
	}()
	tmp = 202
	return
}

func b() {
	//Захватывает одно текущее значение и не меняет его.
	//Оставляет 101.
	tmp := 101
	fmt.Println(tmp)
	defer func(v int) {
		fmt.Println(v)
	}(tmp)
	tmp = 202
	return
}
