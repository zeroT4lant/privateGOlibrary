package main

import (
	"fmt"
)

//Интерфейсы
//Интерфейс используется для определения набора методов, которые должен реализовывать какой-либо тип данных
//Интерфейс может хранить в себе любые типы и значения, удолетворяющие ему

// В отличии от пустых структур
// var emptyInterface interface{} - может хранить совершенно любые значения
// var emptyStruct struct{} - используется для создания переменных без данных

// 4) Интерфейсы
type Foo struct{}

func (f *Foo) A() {}
func (f *Foo) B() {}
func (f *Foo) C() {}

type AB interface {
	A()
	B()
}

type BC interface {
	B()
	C()
}

// Основной main
func main() {
	var any interface{}

	any = "foobar"

	if s, ok := any.(string); ok {
		println("this is a string:", s)
	}

	var chislo interface{}
	chislo = "aboba"
	Type_Switch(chislo)
	Type_Switch(34)

	var f AB = &Foo{}
	y := f.(BC) // сработает ли такой type-assertion?
	//y.A()       // а этот вызов?
	_ = y

	//---------------------
	//Утверждение типов - Type Assertion
	var a interface{} = 10
	n, ok := a.(int)
	fmt.Println(n, ok)

	var b interface{} = 11

	if m, ok := b.(string); ok {
		fmt.Println(m)
	} else {
		// У переменной 'b' тип не стринг
		fmt.Println("B is not a string")
	}
}

// Type switch - способ проверить какого типа значение хранится в интерфейсе
func Type_Switch(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("Twice %v is %v\n", v, v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("I don't know about type %T!\n", v)
	}
}
