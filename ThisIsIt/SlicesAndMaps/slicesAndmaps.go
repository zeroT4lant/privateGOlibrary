package main

import (
	"fmt"
	"sort"
)

type Person struct {
	Name string
	Age  int
}

// в слайсе хранится ссылка на первое значение элемента массива (области памяти)
// или просто ссылка на массив с определённым типом данных, так же в слайсе есть два атрибута len и cap
// пустой интерфейс interface{} в го означает пустой тип данных, под который подходит любой другой тип

func main() {
	//МАССИВЫ - примитивный тип данных, ПРОВЕРЬ!!!(поменять при помощи функции не получится, так как не по адрессу работает, а по значениям(скопирует значения, а не адресс))
	//
	//хранит в себе значение, а не ссылку на первое значение. т.е при переприсваивании копирует массив
	a := [4]int{1, 5, 4, 3}
	b := a
	a[0] = 342
	fmt.Println(a)
	fmt.Println(b)
	b[0] = 678
	fmt.Println(a)
	fmt.Println(b)
	//a и b - разные массивы !!!

	//слайсы разных размеров имеют разные типы!!!
	testSlice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	//длина - len будет 2 элемента в этом случае
	//!!!!А cap оставшаяся необрезанная длина слайса, от первого указанного элемента - до конца исходного слайса, то есть 7!!!!
	//Пятый индекс не включаем
	testArr := testSlice[2:5]
	fmt.Println("testArr = ", testArr)
	//вместо вставки обновляет пятый индекс
	testArr = append(testArr, 523)
	//cap testArr с выделенными значениями и оставшимся местом ((3,4,5),_,_,_,_)
	fmt.Println(len(testArr), cap(testArr), testArr)
	fmt.Println(len(testSlice), cap(testSlice), testSlice)

	//в структуре имеет ссылку на базовый массив , А ТОЧНЕЕ на его ПЕРВЫЙ ЭЛЕМЕНТ

	//ОБЯЗАТЕЛЬНО ПОСМОТРИ В ФУНКЦИЮ modify!!! Вспомни как работает аппенд
	//СЛАЙСЫ - в своей структуре имеет ссылку на область памяти со значениями, длину и вместимость
	//ЕСЛИ ПЕРЕПОЛНИТЬ слайс(append), то он аллоцируется в памяти в другом месте(создаст новую переменную - перезапишет её), и слайс передаст прежние значения, в новый слайс(сосуд)
	//при срезе с массива или слайса мы не создаем новый слайс, а ссылаемся на старый
	//для копирования по значению используем copy

	//сортировка слайса кастомного типа при помощи компоратора
	//Чтобы отсортировать срез, сохраняя при этом исходный порядок одинаковых элементов, используйте вместо этого sort.SliceStable.
	//ТАК ЖЕ ЕСТЬ МЕТОД isSorted !!!

	//Вариативные аргументы, троеточие перечисляет все входящие элементы
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := []int{6, 7, 8, 9, 10}
	fmt.Println(append(slice1, slice2...))

	//удаление из массива числа 5, крайний индекс справа не учитывается
	slice1 = slice1[:len(slice1)-1]
	fmt.Println(slice1, len(slice1), cap(slice1))

	people := []Person{
		{"Alice", 25},
		{"Bob", 30},
		{"Eve", 26},
	}
	fmt.Println("Before sorting:", people)
	sort.Slice(people, func(i, j int) bool {
		return people[i].Age < people[j].Age
	})

	fmt.Println("After sorting:", people)

	//уникальные значения из мапы
	//все неповторяющиеся значения хранятся в ключе, записавшись один раз, мапа будет пропускать одинаковые значения
	nums := []int64{55, 33, 22, 5, 3, 1, 6, 345, 223, 54}
	//сортировка по возрастанию, в аргументе передаётся слайс любого типа
	//вторым аргументом идёт функция, задача которой сравнить, правда ли что один элемент больше другого
	//под капотом sort.Slice используется быстрая сортировка
	sort.Slice(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})
	fmt.Println(nums)

	//MAPS
	//Мапы - тоже ссылочный тип !!! Изменяется как и оригинал, а не копируется.
	//При передачи карты функции или методу содержимое карты может измениться.
	//Такое поведение напоминает несколько срезов, что указывают на один и тот же базовый массив.
	sliceForMap := []int{1, 5, 3, 2, 2, 5, 223, 4, 6, 7, 7}
	mapNums := make(map[int]struct{}, 5)
	for _, el := range sliceForMap {
		//елемент слайса - el, будет как ключ у мапы
		//если не нашлось записи с таким ключом, то записываем
		//делаем из значения слайса - ключ, и записываем в него значение true
		if _, ok := mapNums[el]; !ok {
			mapNums[el] = struct{}{}
		}

		//или вместо этого просто делаем так
		//mapNums[el] = struct{}{}
	}

	fmt.Println(mapNums)
	//можно было бы поменять значение
	//mapNums[1] = false
	fmt.Println(mapNums)

	//подсчёт дубликатов значений
	sliceForMapString := []string{"a", "b", "a", "c", "d", "d", "a"}
	mapStrings := make(map[string]int, 5)
	for _, v := range sliceForMapString {
		if _, ok := mapStrings[v]; !ok {
			mapStrings[v] = 1
		} else {
			mapStrings[v]++
		}

		//или
		//mapStrings[v]++
	}

	fmt.Println("mapStrings: ", mapStrings)

	//поиск самого популярного слова в слайсе
	sussyBaka := []string{"aa", "b", "bb", "bb", "a", "c"}
	fmt.Println(MostPopularWord(sussyBaka))

	specSlice := []int{5, 1, 2}
	//можно поменять следующий элемент аппендом при помощи передачи слайса
	//modifySlice(specSlice[:1])
	//ОБЬЯСНЯЕТСЯ ПОВЕДЕНИЕМ СЛАЙСА, СМОТРИ ЧУТЬ НИЖЕ
	modifySlice(specSlice)
	fmt.Println("spec slice - ", specSlice)

	aboba := []int{1, 5, 2, 6, 8}
	tryaska := aboba[:2]
	fmt.Println(cap(tryaska))
	//пока не превышает слайс aboba, то меняет его элементы. Попробуй убрать и посмотри что будет
	tryaska = append(tryaska, 228, 12333, 123, 555)
	fmt.Println(aboba, " Прикол с аппендом")
	fmt.Println(tryaska, " Прикол с слайсом")
}

// НЕ ЗАБЫВАЙ ПРО ПРИКОЛЫ АППЕНДА
// можем передавать сразу слайс в аргумент, потому что он хранит ССЫЛКУ
func modifySlice(nums []int) {
	//!!!Если изменение до апппенда, то поменяется индекс
	//!!!Если после, то не меняется

	//с copy можно создать копию слайса

	//append прерывает работу со слайсом переданным в аргументе, если мы привысили размер текущего слайса. Превосходя cap слайса, он увеличивается в размера в 2 раза и аллоцируется в новом участке памяти, перенося значения

	//до сюда приходит слайс с первым значениеm [5]
	nums[0] = 2
	//меняется первый элемент на 2, всё хорошо

	//слайс nums[2]

	//аппендим в конец слайса, увлечивая len & cap, теперь в слайсе nums[2,4]
	//если хочешь пользоваться аппендом, возвращай новый слайс в конце функции
	nums = append(nums, 4)

	//если передали срез(часть) от исходного слайса, то измениться и настоящий
}

// для сопоставления значений мапы нам понадобится структура как слайс
// делаем из записей мапы - слайс структур
type kv struct {
	k string
	v int
}

func MostPopularWord(arr []string) string {
	kvStructSlice := make([]kv, 0, len(arr))
	mapForSlice := make(map[string]int, len(arr))

	//убираем все дубликаты, каждую повторную запись, пишем в мапу как значение для единственного ключа
	for _, v := range arr {
		//из значения слайса, делаем ключ для мапы --- newMap[v] = 1 --- присваиваем ему единичку
		if _, ok := mapForSlice[v]; !ok {
			mapForSlice[v] = 1
		} else {
			mapForSlice[v]++
		}
	}

	//сопоставляем значения мапы в слайс структур, чтобы потом могли отсортировать
	//используем kvStructSlice - как слайс
	for k, v := range mapForSlice {
		kvStructSlice = append(kvStructSlice, kv{k, v})
	}

	//сортируем
	sort.Slice(kvStructSlice, func(i, j int) bool {
		return kvStructSlice[i].v < kvStructSlice[j].v
	})
	return kvStructSlice[0].k
}

//func UniqueSortedUserIDs(userIDs []int64) []int64 {
//	for i := 0; i < len(userIDs); i++ {
//		for j := 0; j < len(userIDs)-1; j++ {
//			if userIDs[j] > userIDs[j+1] {
//				userIDs[j], userIDs[j+1] = userIDs[j+1], userIDs[j]
//			}
//		}
//	}
//	return userIDs
//}
