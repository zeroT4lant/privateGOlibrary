package main

import (
	"fmt"
	"sync"
)

//TODO ДОБАВЬ ШТУКУ С ОШИБКОЙ

// Слияние слайсов - ПОМЕНЯТЬ И ПОВТОРИТЬ!!
func merge(nums1 []int, nums2 []int) {
	// if len(nums1) > len(nums2) {
	// 	counter := 0
	// 	for i := len(nums1); i < len(nums2); i++ {
	// 		nums1[i] = nums2[counter]
	// 		counter++
	// 	}
	// }

	//m,n - количество значений в слайсе
	// for i := m; i < len(nums1); i++ {
	// 	nums1[i] = nums2[i-m]
	// }

	for i := 0; i < len(nums1)-1; i++ {
		for j := 0; j < len(nums1)-i-1; j++ {
			if nums1[j] > nums1[j+1] {
				nums1[j], nums1[j+1] = nums1[j+1], nums1[j]
				//or
				// temp := nums1[j]
				// nums1[j] = nums1[j+1]
				// nums1[j+1] = temp
			}
		}
	}

	// create 2 slice we want to merge
	// intSlice1 := [5]int{2, 4, 6, 8, 10}
	// intSlice2 := [4]int{3, 5, 7, 9}

	// 	// create a destination slice
	// var mergeSlice [len(intSlice1) + len(intSlice2)]int

	// 	// copy all elements from slice 1 to destination slice
	// copy(mergeSlice[:], intSlice1[:])

	// 	// copy all elements from slice 2 to destination slice
	// copy(mergeSlice[len(intSlice1):], intSlice2[:])
	// fmt.Printf("%v\n", mergeSlice)

	// fmt.Println(nums1)
}

// Изоморфные строки - ПОВТОРИТЬ!!
func isIsomorphic(s string, t string) bool {
	// Создаем две карты (map) для отображения символов из s в t и наоборот
	sMap := make(map[byte]byte)
	tMap := make(map[byte]byte)

	// Проверяем каждый символ из s и t
	for i := 0; i < len(s); i++ {
		// Если символы уже были отображены, проверяем их соответствие
		if _, ok := sMap[s[i]]; ok {
			if sMap[s[i]] != t[i] {
				return false
			}
		} else {
			// Если символы не были отображены, добавляем их в карты
			sMap[s[i]] = t[i]
		}

		if _, ok := tMap[t[i]]; ok {
			if tMap[t[i]] != s[i] {
				return false
			}
		} else {
			tMap[t[i]] = s[i]
		}
	}

	// Если проверка пройдена успешно, возвращаем true
	return true
}

// ПЕРЕСЕЧЕНИЕ СЛАЙСОВ - через мапу делается
func intersection(a, b []int) []int {
	mapa := make(map[int]int)
	var res []int

	for _, v := range a {
		if _, ok := mapa[v]; !ok {
			mapa[v] = 1
		} else {
			mapa[v]++
		}
	}

	for _, v := range b {
		if _, ok := mapa[v]; ok {
			mapa[v] -= 1
			res = append(res, v)
		}
	}
	return res
}

// ОБЪЕДИНЕНИЕ КАНАЛОВ
func mergeChannels(cs ...<-chan int) <-chan int {
	mergedCh := make(chan int, 1)
	wg := &sync.WaitGroup{}
	wg.Add(len(cs))

	for _, ch := range cs {
		go func(ch <-chan int) {
			//Считаем выполненную работу с каналами
			defer wg.Done()
			for v := range ch {
				mergedCh <- v
			}
		}(ch)
	}

	//закрываем всё
	go func() {
		wg.Wait()
		close(mergedCh)
	}()

	return mergedCh
}

// Факториал
func factorial(n int) int {
	if n < 1 {
		return 1
	}

	return n * factorial(n-1)
}

// Фибоначи
func fibonacci(n int) int {
	if n < 2 {
		return n
	}

	return fibonacci(n-1) + fibonacci(n-2)
}

// ПАЛИНДРОМ
func isPalindrome(word string) bool {
	//Строка - массив байтов
	//Читаем байты - циферки
	//for i := 0; i < len(word); i++ {
	//	fmt.Println(word[i])
	//}

	//Читаем тоже как бы по байтам, но по сути - это руны
	//Которые мы можем преобразовать в строку string()
	//Здесь i - индекс, rune - руна(почти байт)
	//fmt.Println("Другое")
	//for i, rune := range word {
	//	fmt.Println(i, string(rune))
	//}

	//сравниваем байты напрямую
	for i := 0; i < len(word); i++ {
		if word[i] != word[len(word)-i-1] {
			return false
		}
	}

	return true
}

// WORKER POOL
func worker(id int, f func(int) int, jobs <-chan int, result chan<- int) {
	for j := range jobs {
		result <- f(j)
	}
}

func main() {
	fmt.Println(isPalindrome("aboba"))

	numJobs := 5

	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	multiplier := func(x int) int {
		return x * 10
	}

	//Обрабатывает внутри jobs и results
	for w := 1; w <= 3; w++ {
		go worker(w, multiplier, jobs, results)
	}

	//Цикл с получение задач и записей их в канал jobs
	//Получаем работу
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	//Выводим результаты
	for i := 1; i <= numJobs; i++ {
		fmt.Println(<-results)
	}
}

// Подсчёт гласных
func countVowels(word string) int {
	counter := 0
	for _, char := range word {
		switch char {
		case 'a', 'A', 'e', 'E', 'i', 'I', 'o', 'O', 'u', 'U':
			counter++
		}
	}
	return counter
}
