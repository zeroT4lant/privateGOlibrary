package main

import (
	"fmt"
	"net/http"
	"sync"
)

func fetchURL(url string, resChan chan<- string) {
	resp, err := http.Get(url)
	if err != nil {
		resChan <- fmt.Sprintf("%s - not ok ", url)
	} else {
		if resp.StatusCode == 200 {
			resChan <- fmt.Sprintf("%s - is ok ", url)
		} else {
			resChan <- fmt.Sprintf("%s - return status: %s ", url, http.StatusText(resp.StatusCode))
		}
	}
}

func main() {
	var urls = []string{
		"https://google.com",
		"https://somesite.com",
		"https://ya.ru",
		"https://dzen.ru",
		"https://youtube.com",
		"http://non-existent.domain.tld",
	}

	resChan := make(chan string, 1)
	jobsChan := make(chan string, 1)

	//Записывает ссылки
	go func() {
		for _, v := range urls {
			jobsChan <- v
		}
		close(jobsChan)
	}()

	workers := 4
	var wg sync.WaitGroup

	//обрабатывает их
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for v := range jobsChan {
				fetchURL(v, resChan)
			}
		}()
	}

	go func() {
		wg.Wait()
		close(resChan)
	}()

	for v := range resChan {
		fmt.Println(v)
	}
}

//Один из вариантов - fan in/fan out
// var wg sync.WaitGroup
// resChan := make(chan string, 1)

// for _, url := range urls {
// 	wg.Add(1)
// 	go func(url string) {
// 		defer wg.Done()
// 		fetchURL(url, resChan)
// 	}(url)
// }
// go func() {
// 	wg.Wait()
// 	close(resChan)
// }()

// for v := range resChan {
// 	fmt.Println(v)
// }
