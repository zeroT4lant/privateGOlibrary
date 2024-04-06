package main

import (
	"fmt"
	"net/http"
)

func main() {
	//без слеша и параметров, обрабатывает только page, а всё что идёт после - нет
	http.HandleFunc("/page", func(w http.ResponseWriter, r *http.Request) {
		//напечатаем наш путь из запроса
		fmt.Fprintln(w, "single page", r.URL.String())
	})

	//в конце пути слэш, обрабатывает и последующие пути вроде /pages/aboba
	// и даже pages/joka/boka
	http.HandleFunc("/pages/", func(w http.ResponseWriter, r *http.Request) {
		//напечатаем наш путь из запроса
		fmt.Fprintln(w, "multipages:", r.URL.String())
	})

	//обработает любой запрос
	http.HandleFunc("/", handler3)

	http.ListenAndServe(":8080", nil)
}

func handler3(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Третий хэндлер, дефолт случай"))
}
