package main

import "fmt"

func main() {
	ibr1 := []int{1, 2, 3, 4, 5}
	ibr2 := []int{6, 7, 8, 9, 10}
	ch1 := make(chan int, 2)
	ch2 := make(chan int, 2)

	go func() {
		for _, v := range ibr1 {
			ch1 <- v
		}
		close(ch1)
	}()

	go func() {
		for _, v := range ibr2 {
			ch2 <- v
		}
		close(ch2)
	}()

	for i := 0; ; i++ {
		select {
		case v, ok := <-ch1:
			if ok {
				fmt.Println(v)
			}
		case v, ok := <-ch2:
			if ok {
				fmt.Println(v)
			}
		}
		if i >= 10 {
			break
		}
	}

	//go func() {
	//	for _, v := range ibr1 {
	//		ch1 <- v
	//	}
	//	close(ch1)
	//}()
	//
	//go func() {
	//	for _, v := range ibr2 {
	//		ch2 <- v
	//	}
	//	close(ch2)
	//}()
	//
	//for i := 1; ; i++ {
	//	select {
	//	case val, ok := <-ch1:
	//		if ok {
	//			fmt.Println("channel 1:", val)
	//		}
	//	case val, ok := <-ch2:
	//		if ok {
	//			fmt.Println("channel 2:", val)
	//		}
	//	}
	//	if i >= 10 {
	//		break
	//	}
	//}

}
