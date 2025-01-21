package main

import "fmt"

type animal interface {
	walker
	runner
}

type bird interface {
	walker
	flyer
}

type flyer interface {
	fly()
}

type walker interface {
	walk()
}

type runner interface {
	run()
}

type Eagle struct {
}

func (e Eagle) fly() {
	fmt.Println("eagle fly")
}

func (e Eagle) walk() {
	fmt.Println("eagle walk")
}

type Cat struct {
}

func (c Cat) run() {
	fmt.Println("cat run")
}

func (c Cat) walk() {
	fmt.Println("cat walk")
}

func main() {

	// var c animal = &Cat{}
	// var e bird = &Eagle{}

	// c.walk()
	// c.run()

	// e.walk()
	// e.fly()

	var c animal = &Cat{}
	var e bird = &Eagle{}

	c.walk()
	c.run()

	e.fly()
	e.walk()

	//СТРУКТУРЫ
	//самый базовый синтаксис объявления структур
	//в первых скобках - поля структуры
	//во вторых - значения полей
	//type varName struct {}{}

	//конструктор структуры - в функции создали объект и присвоили его переменной
	emp1 := newEmployee("babidjon", "no", 23, 2300)
	fmt.Println(emp1)

	//метод структуры - перед названием в объявлении функции писали - метод какой структуры
	//для вызова нужен объект этой структуры
	//emp1.getInfo()
	emp1.getInfo()

	//метод сеттер - использует указатель на объект структуры
	//на этот раз ссылку на объект структуры передаём в аргумент
	setName1(&emp1, "ymedjon")
	fmt.Println(emp1.getInfo())

	//при создании функции в ресивере указали на объект какой структуры ссылаемся
	emp1.setName2("bedjon")
	fmt.Println(emp1.getInfo())
}

type employee struct {
	name   string
	sex    string
	age    int
	salary int
}

// конструктор структуры
func newEmployee(name string, sex string, age int, salary int) employee {
	return employee{
		name:   name,
		sex:    sex,
		age:    age,
		salary: salary,
	}
}

// сеттер структуры - использует указатель на структуру и меняет указанный объект
func setName1(e *employee, name string) {
	e.name = name
}

// ещё один способ создания сеттера структуры
// в ресивере ссылаемся на указатель объекта
func (e *employee) setName2(name string) {
	e.name = name
}

// метод структуры, если хотим получать какие-то данные из её объектов
// теперь метод гет можно вызвать только использовав объект структуры employee
// то пишем в начале из какой структуры - перед названием
func (e employee) getInfo() string {
	return fmt.Sprintf("Сотрудник %s\nВозраст: %d\nЗарплата%d\n", e.name, e.age, e.salary)
}

//------

// type animal interface {
// 	walker
// 	runner
// }

// type bird interface {
// 	walker
// 	flyer
// }

// type walker interface {
// 	walk()
// }

// type runner interface {
// 	run()
// }

// type flyer interface {
// 	fly()
// }

// type Cat struct {
// }

// func (c Cat) walk() {
// 	fmt.Println("cat is walking")
// }

// func (c Cat) run() {
// 	fmt.Println("cat is running")
// }

// type Eagle struct {
// }

// func (e Eagle) walk() {
// 	fmt.Println("eagle is walking")
// }

// func (e Eagle) fly() {
// 	fmt.Println("eagle is flying")
// }

// ```
// type iface struct {
// tab  *itab
// data unsafe.Pointer
// }
// ```

// `data` - это непосредственно наш f1 из примера. Вся информация по методам, типам и прочему скрывается в `tab` - интерфейсной таблицей (interface table).

// ```
// type itab struct {
// inter *interfacetype
// _type *_type
// hash  uint32 // copy of _type.hash. Used for type switches.
// _     [4]byte
// fun   [1]uintptr // variable sized. fun[0]==0 means _type does not implement inter.
// }
