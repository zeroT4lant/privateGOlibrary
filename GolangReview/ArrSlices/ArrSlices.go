package main

import "fmt"

//Слайсы передаются "по ссылке"
//(фактически будет передана копия структуры slice со своими len и cap, НО УКАЗАТЕЛЬ НА МАССИВ ARRAY БУДЕТ ТОТ-ЖЕ САМЫЙ)
//Массивы с элементами одного типа, но разной длины - это разные типы.

// !!!!!!!! слайс - это структура
// Пока аппендом не превысли ёмкость первичного слайса, то новый слайс ссылается на старый
func XXX() {
	list := make([]int, 4, 4)

	list2 := append(list, 1) //Новый слайс, так как превысил cap
	//То есть больше не привязан к list
	//list2{0 0 0 0 1}

	list[0] = 5  // 5 0 0 0
	list2[0] = 9 // 9 0 0 0 1

	fmt.Println(list)
	fmt.Println(list2)
}

func XXX2() {
	list := make([]int, 4, 5)

	list2 := append(list, 1) //Ссылается на старый слайс, пока append'ом не превысили cap и не создали новый слайс.
	//То есть больше не привязан к list
	//list2{0 0 0 0 1}

	list[0] = 5  // 5 0 0 0
	list2[0] = 9 // 9 0 0 0 1

	//list2 ссылается на list, так как аппендом не превысили лимит и не создали новый слайс
	fmt.Println(list)  // 9 0 0 0
	fmt.Println(list2) // 9 0 0 0 1
}

// --------------------
func HSkills() {
	nums := make([]int, 1, 2)
	fmt.Println(nums) // 0

	//Не добавится элемент
	appendSlice(nums, 1024) // 0

	//Элемент не добавился, так что нет индекса 1, числа нет.
	mutateSlice(nums, 1, 512)
	fmt.Println(nums) //nums 0
}

func appendSlice(sl []int, val int) {
	sl = append(sl, val)
}

func mutateSlice(sl []int, idx, val int) {
	sl[idx] = val
}

//--------------------

//Слайс - динамический массив, который может храанить элементы одного типа.
//Слайс - ссылка на массив, который содержит элементы слайса.
//Слайс можно передавать в функции по значению.Но изменение элементов слайса в функции отразится на исходном.
//Так как они ссылаются на один и тот же массив.

//Массив - значение, слайс - ссылка. При передаче слайса в функцию/переприсваивании - происходит копирование его заголовка, но не элементов.
//Массивы можно сравнивать с другим массивом такого же размера и типа с помощью '=='

//Слайс можно сравнить только с nil
//struct slice {
//array unsafe.Pointer --- указатель на первый элемент массива,который содержащит элементы слайса.
//len int
//cap int
//}

//Append - принимает слайс и несколько элементов для вставки и возвращает новый слайс, с добавленными данными.
//Таким образом, если вместимости исходного слайса достаточно для добавления новых элементов,
//То она копирует их в свободное место массива, на который ссылается слайс.
//
//Если вместимость недостаточно, то функция append - выделяет новый массив большего размера, копирует все элементы исходного слайса и добавляет новые элементы
// и возвращает новый слайс, который ссылается на новый массив.
//Однако, если мы добавим элементы в подслайс с помощью функции append,
//то это может привести к перевыделению памяти и созданию нового массива,
//если вместимость подслайса недостаточна. В этом случае, подслайс будет ссылаться на новый массив, а исходный слайс - на старый.
////Стоит быть внимательным при добавлении элементов в слайс в функции с помощью append, так как мы изменяем копию структуры слайса.

//При размере слайса больше 1024 элементов, слайс увеличивается на четверть текущего размера.

// Через функцию можем поменять значение переменной.
// При помощи апппенда поменять если передавать срез arr[:]
// Либо не возвращать новое значение через return
func modify(arr []int) {
	//arr[0] = 10
	arr = append(arr, 10)
}

func main() {
	//a()
	//d()
	//e()

	x := []int{1, 2, 3, 4}
	x = append(x, 5)
	fmt.Println(len(x), cap(x))
	//поменяет значение аппендом, если передать срез x[:1]
	modify(x)
	fmt.Println(x)

	//---------------------------------------
	specSlice := []int{5, 1, 2}
	//можно поменять аппендом при помощи передачи слайса
	//ОБЬЯСНЯЕТСЯ ПОВЕДЕНИЕМ СЛАЙСА, СМОТРИ ЧУТЬ НИЖЕ
	modifySlice(specSlice[:1])
	fmt.Println(specSlice)

	aboba := []int{1, 5, 2, 6, 8}
	tryaska := aboba[:2]
	fmt.Println(cap(tryaska))
	//пока не превышает слайс aboba, то меняет его элементы. Попробуй убрать и посмотри что будет
	tryaska = append(tryaska, 228, 12333, 123, 555)
	fmt.Println(aboba, " Прикол с аппендом")
	fmt.Println(tryaska, " Прикол с слайсом")

	fmt.Println("----XXX----")
	XXX()

	fmt.Println("----XXX2----")
	XXX2()
}

// НЕ ЗАБЫВАЙ ПРО ПРИКОЛЫ АППЕНДА
// можем передавать сразу слайс в аргумент, потому что он хранит ССЫЛКУ
func modifySlice(nums []int) {
	//с copy можно создать копию слайса

	//append прерывает работу со слайсом переданным в аргументе, если мы привысили размер текущего слайса. Превосходя cap слайса, он увеличивается в размера в 2 раза и аллоцируется в новом участке памяти, перенося значения

	//до сюда приходит слайс с первым значениеm [5]

	//меняется первый элемент на 2, всё хорошо
	nums[0] = 2
	//слайс nums[2]

	//аппендим в конец слайса, увлечивая len & cap, теперь в слайсе nums[2,4]
	//если хочешь пользоваться аппендом, возвращай новый слайс в конце функции
	nums = append(nums, 4)

	//если передали срез(часть) от исходного слайса, то измениться и настоящий
}

// 1) Слайсы
func a() {
	x := []int{} //{0,1,2}
	x = append(x, 0)
	x = append(x, 1)
	x = append(x, 2) // 3 4 //x = {0,1,2}
	// y - копирует значения из X при аппенде и добавляет туда цифру 3
	y := append(x, 3) // x = {0,1,2} y = {0,1,2,3}
	//1)
	fmt.Println(x, y) //x = {0,1,2} y = {0,1,2,3}
	// z - копирует значения из X при аппенде и добавляет туда цифру 4
	z := append(x, 4) // {0,1,2,4}
	//2)
	//fmt.Println(y, z)
	//3)
	fmt.Println(x, y, z)
	//Разобраться почему меняется значение!!!!!!!!!

	s1 := []int{1, 2, 3}    // создаем первый слайс
	s2 := []int{4, 5, 6}    // создаем второй слайс
	s3 := append(s1, s2...) // добавляем второй слайс в конец первого слайса
	fmt.Println(s3)         // выводит [1 2 3 4 5 6]

}

// C помощью операции среза [:] создаем слайс sl.
// В этом случае не происходит копирования элементов массива arr, слайс sl просто ссылается на те же данные.
// Поэтому изменения в слайсе sl, также отразятся на массиве arr
func b() {
	arr := [5]int{1, 2, 3, 4, 5}
	sl := arr[:]
	sl[0] = 3
	fmt.Println(arr)
	fmt.Println(sl)
	//[1, 2, 3, 4, 5] -> [3, 2, 3, 4, 5]
}

func c() {
	arr := [5]int{1, 2, 3, 4, 5}
	sl := arr[1:]
	sl[0] = 3
	fmt.Println(arr)
	fmt.Println(sl)
	//arr [1, 2, 3, 4, 5] -> [1, 3, 3, 4, 5]
}

func d() {
	arr := [5]int{1, 2, 3, 4, 5}
	sl := arr[1:2]
	sl = append(sl, 4)
	fmt.Println(sl)
	fmt.Println(arr)
	//arr [1, 2, 3, 4, 5] -> [1, 2, 4, 4, 5] ; sl [2] -> [2, 4]
}

// Все привязаны к исходному массиву arr
func e() {
	arr := [5]int{1, 2, 3, 4, 5}
	sl := arr[1:2]       //2
	sl2 := append(sl, 4) //2,4
	sl2[0] = 8           //8,4
	fmt.Println(sl2)
	fmt.Println(sl)
	fmt.Println(arr)
	//arr[1, 2, 3, 4, 5] -> [1, 8, 4, 4, 5]
	//sl[2] -> sl[8]
	//sl2 [2, 4] -> [8, 4]
}

func f() {
	arr := [5]int{1, 2, 3, 4, 5}
	sl := arr[1:2]     //[2]
	sl = append(sl, 9) //[2 9]
	sl = append(sl, 8) //[2 9 8]
	sl = append(sl, 7) //[2 9 8 7]
	fmt.Println(arr)   // 1 2 9 8 7

	//Когда SL превысил ёмкость, то выделилась область под новый слайс X2
	//и туда скопировались все элементы
	//теперь SL не ссылается на ARR, то есть изменения больше не влияют на arr
	sl = append(sl, 6) // [2,9,8,7,6]
	fmt.Println(arr)   // 1 2 9 8 7

	sl[0] = 0
	sl[1] = 0
	fmt.Println(sl)  //[0 0 8 7 6]
	fmt.Println(arr) //[1 2 9 8 7]
}

func g() {
	var x []int
	x = append(x, 0) // 0
	x = append(x, 1) // 0 1
	x = append(x, 2) // 0 1 2

	//y и z - указывают на один и тот же слайс X
	//Когда Z модифицируется, изменения касаются и Y.Потому что разделяют одну область памяти X.
	y := append(x, 3) // y - 0 1 2 3
	z := append(x, 4) // z - 0 1 2 4
	//y - 0 1 2 4 / z - 0 1 2 4
	fmt.Println(y, z)
}
