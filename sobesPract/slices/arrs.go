package main

import (
	"fmt"
)

func main() {
	arr := [4]int{5, 6, 1, 2}
	//когда arr-массив, arr[:]-слайс, будет меняться как и массив, ведь это ссылка на него

	aboba := arr[:2]

	aboba = append(aboba, 56)

	arr[0] = 228

	aboba[1] = 4
	fmt.Println(aboba)
	fmt.Println(arr)

	fmt.Println("-----------------------------------")

	one := make([]int, 0, 3)
	one = append(one, 2, 3)
	two := one
	fmt.Println("cap", cap(one), len(one))

	two[0] = 123
	fmt.Println(one, two)

	one = append(one, 777)

	fmt.Println(one, two)

	two[1] = 229
	fmt.Println(one, two)

	fmt.Println("-----------------------------------")
	ints := []int8{2}
	ints = append(ints, 6)
	fmt.Println("cap", cap(ints))
	fmt.Println("-----------------------------------")

	newArr := make([]int, 6, 10)
	fmt.Println(newArr, len(newArr), cap(newArr))
	subNewArr := newArr[2:3]
	subNewArr = append(subNewArr, 228)
	fmt.Println(newArr, len(newArr), cap(newArr))
	fmt.Println(subNewArr, len(subNewArr), cap(subNewArr))

	fmt.Println("-----------------------------------")
	test1 := []int{1, 2, 3, 4, 5}
	test1 = test1[:3]
	fmt.Println(test1)
	test2 := test1[3:]
	fmt.Println(test2)
	fmt.Println(test2[:2])
}
