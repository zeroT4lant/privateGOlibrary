package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

//https://www.youtube.com/watch?v=MmIdk_0HhVI

func main() {
	//1 Задача
	//Берёт не из своего скоупа переменную, так что передаём параметр в аргумент
	//1.22 Go
	wg := &sync.WaitGroup{}

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(i int) {
			fmt.Println(i)
			defer wg.Done()
		}(i)
	}
	wg.Wait()

	//2 Задача
	urls := make(chan string)

	//Запускаем функцию для параллельного скачивания контента
	go func() {
		urls <- "http://example.com"
		urls <- "http://example.org"
		urls <- "http://example.net"
		close(urls)
	}()

	//Запускаем параллельное скачивание с макс кол-вом воркеров
	result := ParallelDownload(context.Background(), urls, 3)

	//Выводим результаты
	for url, content := range result {
		fmt.Printf("Сайт: %s\nЗаголовок: %sДата: %s\nСодержание: %s\n\n",
			url, content.Title)
	}
}

// Имитация долгой работы, загрузки контента сайта
func DownloadSiteContent(ctx context.Context, url string) SiteContent {
	rand.Seed(time.Now().UnixNano())
	sleepTime := time.Duration(rand.Intn(6)+5) * time.Second

	time.Sleep(sleepTime)

	return SiteContent{
		Title:   "Заголовок сайта " + url,
		Date:    time.Now(),
		Content: "Cодержание сайта " + url,
	}
}

type SiteContent struct {
	Title   string
	Date    time.Time
	Content string
}

// Функция для параллельного скачивания контента сайтов
// Горутина для получения данных
// Канал отдаёт данные
func ParallelDownload(ctx context.Context, urls <-chan string, numWorkers int) map[string]SiteContent {
	//TODO
	mu := &sync.Mutex{}
	wg := &sync.WaitGroup{}

	mappa := make(map[string]SiteContent, len(urls))

	// Запускаем воркеров в отдельных горутинах

	for i := 0; i < numWorkers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			for {
				select {
				case <-ctx.Done():
					break
				case url, ok := <-urls:
					if !ok {
						break
					}

					content := DownloadSiteContent(ctx, url)
					//Записываем в мапу
					mu.Lock()
					mappa[url] = content
					mu.Unlock()
				}

			}
		}()
	}

	wg.Wait()
	return mappa
}

// 3 Задача
// AAAAAAABBBCCXYZDDDEEEAAABBBBBBBBBBBBBBB
// Подсчитать совпадения
func RLE(in string) string {
	s := []byte(in + " ")
	res := []byte{}
	cnt := 1

	for i := 0; i < len(s)-1; i++ {
		if s[i] == s[i+1] {
			cnt++
			continue
		} else if cnt == 1 {
			res = append(res, s[i])
		} else {
			res = append(res, s[i], byte(cnt))
			cnt = 1
		}
	}
	return string(res)
}
