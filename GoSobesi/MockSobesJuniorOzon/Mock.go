package main

import (
	"fmt"
	"sync"
	"time"
)

//https://www.youtube.com/watch?v=5-rENjTvYeU

func main() {
	//!!!1 Задача!!!
	// v := "Hello 📌"
	// //индекс - i и руны - c(могут быть переведены в string при помощи string(c) )
	// for i, c := range v {
	// 	fmt.Printf("%d of '%s'\n", i, c)
	// }

	// //
	// emoji := []rune("cool📌")
	// for _, ch := range emoji {
	// 	fmt.Printf("%d of '%s'\n", ch, string(ch))
	// }

	//!!!2 Задача - APPEND копирует значение из первого аргумента и вставляет его!!!
	//arr1 = append(arr2, 9) - переназначаем значения из arr1, в arr2 и добавляем 9.

	// arr1 := []int{1, 2, 3, 4}
	// arr2 := []int{5, 6, 7, 8}
	// arr1 = append(arr2, 9)
	// fmt.Println("ABOBA: ", arr1)

	//s1 := make([]account, 0, 2)
	//s1 = append(s1, account{})  //{acc1}
	//s2 := append(s1, account{}) //{acc1,acc2}
	//
	//acc := &s2[0]
	//
	//acc.value = 100
	//fmt.Println(s1, s2) //s1{acc1=100}___s2{acc1=100,acc2=0}
	//
	//s1 = append(s2, account{}) // s1{acc1=100,acc2=0,acc3=0}больше не связан с s2, так как превысил cap
	//acc.value += 100           // s2{acc1=200,acc2=0}
	//fmt.Println(s1, s2)

	//!!!3 Задача
	//a := make([]int32, 0)
	////По одному элементу добавляет в слайс значения, увеличивая len и cap постепенно
	//a = append(a, []int32{1, 2, 3}...) // len-cap 1-1 -> 2-2 -> 3-4
	//fmt.Println(len(a), cap(a))
	//a = append(a, 4) // 4-4
	//
	//fmt.Println(len(a), cap(a))
	//
	//fmt.Println("ТЕПЕРЬ 'b' с int64")
	//b := make([]int64, 0)
	////Сразу добавляет 3 элемента. len и cap становятся 3-3
	//b = append(b, []int64{1, 2, 3}...) // len-cap 3-3
	//fmt.Println(len(b), cap(b))
	//b = append(b, 4) // cap - увеличивается в 2 раза. 3-3 -> 4-6
	//
	//fmt.Println(len(b), cap(b))

	//!!!4 Задача
	//groceries := map[string]string{
	//	"📌": "150 рублей",
	//	"😀": "0 рублей",
	//	"🤫": "750 рублей",
	//}
	//
	//for idx, item := range groceries {
	//	fmt.Println(idx, item)
	//}
	//
	////Сортируем мапу
	//fmt.Println(groceries)

	//!!!5 Задача
	//fiveTask()

	//!!!6 Задача
	//sixTask()

	//!!!7 Задача
	//sevenTask()

	//!!!8 Задача
	mergeChan()
}

// 5 TASK ------------
//type I interface {
//	Foo()
//}
//
//type S struct{}
//
//func (s *S) Foo() {
//	fmt.Println("foo")
//}
//
//func Build() I {
//	//res с конкретным типом *S Реализует интерфейс I
//	var res *S
//	return res
//}
//
//func fiveTask() {
//	i := Build()
//	//Выведется foo, потому что i имеет конкретный тип, следовательно не равен nil
//	if i != nil {
//		i.Foo()
//	} else {
//		fmt.Println("nil")
//	}
//}

// 6 TASK ------------
//type X struct {
//	V int
//}
//
//func (x X) S() {
//	fmt.Println(x.V)
//}

func sixTask() {
	//Здесь запоминает первое попавшееся значение
	//x := X{123}
	//Если хотим, чтобы выводило 456, то вызываем анонимную функцию и не передаём переменную(поэтому возьмёт последнее попавшееся значение)
	//В ином случае если укажем в анонимной функции переменную, то захватит первую попавшуюся переменную
	//defer func() {
	//	x.S()
	//}()
	//x.V = 456

	//accum
	a := accum()

	fmt.Println(a(1)) //1
	fmt.Println(a(1)) //2
}

func accum() func(int) int {
	sum := 0
	//Передаём функцию, которая вернёт число
	//sum увеличивается.
	return func(x int) int {
		sum += x
		return sum
	}
}

// 6 TASK ------------

// 7 TASK ------------

type SafeCounter struct {
	v  map[string]int
	mu sync.Mutex
}

func (c *SafeCounter) Inc(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.v[key]++
}

func (c *SafeCounter) Value(key string) int {
	return c.v[key]
}

func sevenTask() {
	c := SafeCounter{v: make(map[string]int), mu: sync.Mutex{}}
	for i := 0; i < 1000; i++ {
		go c.Inc("somekey")
	}

	time.Sleep(time.Second)
	fmt.Println(c.Value("somekey"))
}

// 7 TASK ------------

// 8 TASK ------------
func mergeChans(chs ...<-chan int) <-chan int {
	wg := &sync.WaitGroup{}
	resChan := make(chan int, 1)

	for _, ch := range chs {
		wg.Add(1)
		go func(ch <-chan int) {
			defer wg.Done()

			for v := range ch {
				resChan <- v
			}
		}(ch)
	}

	go func() {
		wg.Wait()
		close(resChan)
	}()

	return resChan
}

// func mergeChan(cs ...<-chan int) <-chan int {
// 	resChan := make(chan int, 1)

// 	wg := &sync.WaitGroup{}

// 	wg.Add(len(cs))

// 	for _, someChan := range cs {
// 		go func(ch <-chan int) {
// 			defer wg.Done()

// 			for val := range someChan {
// 				resChan <- val
// 			}
// 		}(someChan)
// 	}

// 	go func() {
// 		wg.Wait()
// 		close(resChan)
// 	}()

// 	return resChan
// }

// 8 TASK ------------

//type account struct {
//	value int
//}
