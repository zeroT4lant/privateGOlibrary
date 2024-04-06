package main

import (
	"fmt"
	"sync"
	"time"
)

// WaitGroup - механизм ожидания завершения группы задач
func main() {
	//если здесь напишем go и не поставим задержку, то горутина не успеет выполнится
	//withoutWait()

	//подсчитает всё, подождёт остальных и завершит выполнение
	withWait()
	//readWithMutex()
}

func withoutWait() {

	//если на печать поставим горутину, то что-то успеет вывестись, если ниже поставим задержку
	for i := 0; i < 10; i++ {
		go fmt.Println(i + 1)
	}

	//можем подождать выполнение счётчика
	//time.Sleep(time.Second)

	fmt.Println("exit")
}

func withWait() {
	//добавлять задачи строго вне горутины!!! ОБЪЯСНЕНИЕ НИЖЕ
	//ТАСКА МОЖЕТ НЕ УСПЕТЬ добавиться к выполнению в WaitGroup
	//И мы пропустим выполнение это таски, так как не добавили её в WG
	//1) создаём WaitGroup
	//2) добавляем 10 задач
	//3) создаём 10 горутин
	//4) wg.Wait() - переход в состояние ожидания выполнения горутин
	//5) идём дальше по программе

	//примитив синхронизации
	//помещаем задачи сюда
	wg := sync.WaitGroup{}
	//задачи можно заранее создать wg.Add(10)
	wg.Add(10)

	//добавили 10 задач
	for i := 0; i < 10; i++ {
		//на каждую итерацию добавляем задачи, которые надо подождать
		//wg.Add(1)

		//создаём 10 горутин - асинхронных действий
		go func(i int) {
			fmt.Println(i + 1)
			//отчитались о выполнении
			//завершаем выполнение задачи
			wg.Done()
		}(i)
	}

	//основная горутина ждёт, пока не выполнятся все задачи из waitGroup
	//состояние ожидания
	wg.Wait()
	fmt.Println("exit")
}

func wrongWait() {
	wg := sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		go func(i int) {
			//ТАСКА МОЖЕТ НЕ УСПЕТЬ добавиться к выполнению в WaitGroup
			//И мы пропустим выполнение это таски, так как не добавили её в WG
			wg.Add(1)
			fmt.Println(i + 1)
			//отчитались о выполнении
			wg.Done()
		}(i)

	}

	wg.Wait()
	fmt.Println("exit")
}

// выполняется одним потоком, но медленно
func writeWithoutConcurrent() {
	//без гонки данных
	start := time.Now()
	var counter int

	//увеличиваем значение счётчика
	for i := 0; i < 1000; i++ {
		time.Sleep(time.Nanosecond)
		counter++
	}

	fmt.Println(counter)
	fmt.Println(time.Now().Sub(start).Seconds())
}

// выполняется ооочень быстро, многими потоками, но теряются значения при записи
func writeWithoutMutex() {
	//Из-за гонки данных DataRace, горутины дважды могут инкерментировать одно число и пропустить другое
	//так как они не согласованы

	//так что во время выполнение некоторые операции счётчика пропускаются
	start := time.Now()
	var counter int
	wg := sync.WaitGroup{}

	wg.Add(1000)
	for i := 0; i < 1000; i++ {
		go func() {
			defer wg.Done()
			time.Sleep(time.Nanosecond)
			counter++
		}()
	}

	wg.Wait()
	fmt.Println(time.Now().Sub(start).Seconds())
	fmt.Println("exit")
}

// без DataRace, но выполняется чуть-чуть дольше
func writeWithMutex() {
	//Mutex и RWMutex - механизм получения исключительной блокировки
	start := time.Now()
	counter := 0
	wg := sync.WaitGroup{}
	mu := sync.Mutex{}

	wg.Add(1000)
	for i := 0; i < 1000; i++ {
		go func() {
			defer wg.Done()
			time.Sleep(time.Nanosecond)

			//действие закрепляется за одной горутиной
			//берёт ответственность за выполнение на себя, так что остальные не пытаются записать туда что-то
			mu.Lock()
			counter++
			mu.Unlock()
		}()
	}

	wg.Wait()
	fmt.Println(counter)
	fmt.Println(time.Now().Sub(start).Seconds())
}

func readWithMutex() {
	start := time.Now()
	var counter int
	wg := sync.WaitGroup{}
	mu := sync.Mutex{}

	wg.Add(100)

	for i := 0; i < 50; i++ {
		go func() {
			defer wg.Done()

			mu.Lock()
			time.Sleep(time.Nanosecond)
			_ = counter
			mu.Unlock()
		}()

		go func() {
			defer wg.Done()

			mu.Lock()
			time.Sleep(time.Nanosecond)
			counter++
			mu.Unlock()

		}()

	}
	wg.Wait()
	fmt.Println(counter)
	fmt.Println(time.Now().Sub(start).Seconds())
}

func readWithRWMutex() {
	counter := 0

	wg := sync.WaitGroup{}
	//можно ещё делать блокировку на чтение
	mu := sync.RWMutex{}

	wg.Add(100)

	for i := 0; i < 50; i++ {
		go func() {
			defer wg.Done()

			//сокращает время на чтение
			//любая горутина может читать счётчик
			mu.RLock()

			time.Sleep(time.Nanosecond)
			_ = counter

			mu.RUnlock()
		}()
		go func() {
			defer wg.Done()

			//эксклюзивная блокировка
			mu.Lock()

			time.Sleep(time.Nanosecond)
			counter++

			mu.Unlock()
		}()

	}

	wg.Wait()
	fmt.Println(counter)
}
