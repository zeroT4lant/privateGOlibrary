package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler)

	fmt.Println("starting server at :8080")
	//запускаем сервер по порту
	http.ListenAndServe(":8080", nil)
}

// w - куда будем записывать результат, r - пришедший запрос
func handler(w http.ResponseWriter, r *http.Request) {

	//записываем в w
	fmt.Fprintln(w, "Привет мир!")

	//аналог, записываем байтами, отображаться будет нормальной строкой
	w.Write([]byte("!!!"))
}
