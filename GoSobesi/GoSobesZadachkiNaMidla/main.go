package main

import (
	"fmt"
	"sort"
)

func main() {
	//0 ZADACHA
	a5 := []int{1, 2, 3}
	b5 := a5
	b5 = append(b5, 4)  //b {1 2 3 4} len-4 cap-6 - !!!отдельный слайс от a теперь
	c5 := b5            //c {1 2 3 4} len-4 cap-6
	b5[0] = 0           //c {0 2 3 4} b {0 2 3 4} len-4 cap-6
	e5 := append(c5, 5) //e {0 2 3 4 5} len-5 cap-6
	b5[2] = 7           //b {0 2 7 4},

	fmt.Println(a5, b5, c5, e5)
	//a - 1 2 3
	//b - 0 2 7 4
	//c - 0 2 7 4
	//e - 0 2 7 4 5
	//--------------------------------

	one := []int{1, 2}
	two := one
	two[0] = 123
	fmt.Println(one, two) //{123 2} {123 2}
	one = append(one, 666)

	//в two осталось последнее изменение от one, до аппенда
	fmt.Println(one, two) //{123 2 666} //{123 2}

	//------
	test1 := []int{1, 2, 3, 4, 5}
	test1 = test1[:3] //1 2 3
	fmt.Println(test1)
	test2 := test1[3:]                  //[]+4 5                 // От 3 индекса и до последнего
	fmt.Println(len(test2), cap(test2)) // Остаются 2 элемента, к котором ниже получим доступ.
	fmt.Println(test2[:2])              //Смотрим на остальные индексы, которые не попали из ориг слайса

	//----------
	//СЛАЙСЫ И МАПЫ
	//----------
	//#1
	v := []int{3, 4, 1, 2, 5}
	ap(v) //ничего не меняется
	sr(v) //сортировка пройдёт
	fmt.Println(v)

	//----------
	//#2
	fmt.Println("###ЗАДАЧА_2###")
	var foo []int
	var bar []int

	foo = append(foo, 1) //1
	foo = append(foo, 2) //1 2
	foo = append(foo, 3) //foo {1 2 3} _ len-3 cap-4
	bar = append(foo, 4) //foo {1 2 3} bar {1 2 3 4}
	foo = append(foo, 5) //foo {1 2 3 5} bar {1 2 3 5}
	//bar ссылается на исходный слайс foo, так что до того момента
	//как он не перенеицилиазируется, то будет ссылаться на исходник - foo

	fmt.Println(foo, bar) //foo {1 2 3 5} bar {1 2 3 5}

	//----------
	//#3
	fmt.Println("###ЗАДАЧА_3###")
	c := []string{"A", "B", "D", "E"}
	b := c[1:2]         //B len-1 cap-3 {B}
	fmt.Println(cap(b)) // 3 - оставшийся объём
	b = append(b, "TT") //{B TT}
	fmt.Println(c)      // {A B TT E}
	fmt.Println(b)      // {B TT}

	//----------
	//#4
	fmt.Println("###ЗАДАЧА_4###")
	//ошибка из-за неицилизированной-нил мапы, в неё нельзя записывать и читать
	//var m map[string]int

	//Правильный вариант
	m := map[string]int{}
	testStroka := []string{"hello", "world", "from", "the", "best", "language", "in", "the", "world"}
	for _, word := range testStroka {
		m[word]++
	}
	for k, v := range m {
		fmt.Println(k, v)
	}
	//----------
	//#5
	fmt.Println("###ЗАДАЧА_5###")
	//функция не возвращает слайс и пытается заапендить
	mutate := func(a []int) {
		a[0] = 0
		a = append(a, 1)
		fmt.Println(a) //0 1 2 3 4 1
	}
	//слайс изменится, точнее первый элемент
	a := []int{1, 2, 3, 4}
	mutate(a)
	fmt.Println(a) // 0 2 3 4

	//----------
	//#6
	fmt.Println("###ЗАДАЧА_6###")
	fmt.Println("---Вариант 1---")
	sl1 := []int{1, 2, 3, 5}
	mod1(sl1)
	fmt.Println(sl1) //Выведет все пятёрки
	fmt.Println("---Вариант 2---")
	sl2 := make([]int, 4, 8) // {0 0 0 0 _ _ _ _}
	sl2[0] = 1               //{1 0 0 0 _ _ _ _}
	sl2[1] = 2               //{1 2 0 0 _ _ _ _}
	sl2[2] = 3               //{1 2 3 0 _ _ _ _}
	sl2[3] = 5               //{1 2 3 5 _ _ _ _}
	mod2(sl2)
	fmt.Println(sl2) //{5 5 5 5 5 5 5 5}
	fmt.Println("---Вариант 3---")
	// Если в cap остались пустые элементы, то спокойно меняет, даже с аппендом
	sl3 := make([]int, 4, 8)
	sl3[0] = 1
	sl3[1] = 2
	sl3[2] = 3
	sl3[3] = 5
	mod3(sl3)
	fmt.Println(sl3) //{5 5 5 5 _ _ _ _}
	fmt.Println("---Вариант 4---")
	sl4 := []int{1, 2, 3, 4, 5}
	//На этот раз cap весь заполнен, так что не меняет значения
	//после аппенда любые взаимодействия со слайсом аннулируются
	mod4(sl4)
	fmt.Println(sl4) //1 2 3 4 5
	//----------
	//#7
	fmt.Println("###ЗАДАЧА_7###")
	s := make([]int, 3, 8)
	m1 := make(map[int]int, 8)

	// add to slice - не добавится и не сможешь обратиться к этому элементу
	a1(s)
	//не можем поменять неициализированное значение
	//println(s[3]) //0 0 0

	// add to map - добавится и выведется
	b1(m1)
	println(m1[3]) // {33}
	//----------
	//#8
	fmt.Println("###ЗАДАЧА_8###")
	a2 := []int{1, 2}   // 1 2
	a2 = append(a2, 3)  //{1 2 3} len 3 cap 4
	b2 := append(a2, 4) //{1 2 3 4}
	c2 := append(a2, 5) //{1 2 3 5}

	fmt.Println(b2) //{1 2 3 5}
	fmt.Println(c2) //{1 2 3 5}
	//----------
	//#9
	fmt.Println("###ЗАДАЧА_9###")
	a3 := []int{1, 2}
	a3 = append(a3, 3) //1 2 3 len-3 cap-4
	a3 = append(a3, 7) //1 2 3 7 len&cap-4
	//При переполнениии аппенд превышает кап и переинициализирует слайс
	b3 := append(a3, 4) //1 2 3 7 4
	c3 := append(a3, 5) //1 2 3 7 5

	fmt.Println(b3) //1 2 3 7 4
	fmt.Println(c3) //1 2 3 7 5

}

func a1(s []int) {
	s = append(s, 37)
}

func b1(m map[int]int) {
	m[3] = 33
}

func mod4(a []int) {
	a = append(a, 125)
	for i := range a {
		a[i] = 5
	}
	fmt.Println(a) // 5 5 5 5 5 5
}

// Если в cap остались пустые элементы, то спокойно меняет
func mod3(a []int) {
	a = append(a, 125) // {1 2 3 5 125 _ _ _}
	for i := range a { //теперь len(5)=range
		a[i] = 5
	}
	fmt.Println(a) //{5 5 5 5 5}
}

func mod2(a []int) {
	for i := range a {
		a[i] = 5
	}
	fmt.Println(a)
}

func mod1(a []int) {
	//Меняет все значения в исходном слайсе
	for i := range a {
		a[i] = 5
	}
	fmt.Println(a)
}

func ap(arr []int) {
	arr = append(arr, 10)
}

func sr(arr []int) {
	sort.Ints(arr)
}
