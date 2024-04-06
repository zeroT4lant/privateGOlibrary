package main

import "fmt"

func main() {
	str := "Hello Almaz"

	//УКАЗАТЕЛЬ - тип данных, который в качестве значения хранит адрес ячейки памяти значения, либо другого указателя
	//&x - адрес X
	//*int - указатель на int
	//может быть nil

	//поменяли значения переменной str, через указатель
	//звёздочка перед типом означает, теперь это указатель на какой тип
	var ptrStr = &str
	*ptrStr = "Hello chumaz"
	fmt.Println(str)

	num := 3
	//не работает как надо, передаём по значению аргумент
	square(num)
	fmt.Println(num)

	//передали указатель

	//ptrSquare(&num)
	//fmt.Println(num)

	//Используя синтаксис *p , мы обращаемся к значению, которое находится по адресу p(получается достучаться и изменить значение n)
	n := 4
	//p теперь имеет указатель на n
	p := &n
	//добираем до значения n, разыменовывая указатель (*p)
	*p = 228
	fmt.Println(n)

	n1 := 0
	p1 := &n1
	*p1 = 404
	fmt.Println(&n1, p1)
	fmt.Println(n1, *p1)
}

func square(n int) {
	n *= n
}

func ptrSquare(n *int) {
	*n *= *n
}

func hasWallet(balance *int) bool {
	return balance != nil
}
