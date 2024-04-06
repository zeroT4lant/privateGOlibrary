package main

import (
	"fmt"
	"net/http"
	"time"
)

// можно с горутиной запустить несколько серверов
func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler)

	//структура для настройки сервера!!!
	server := http.Server{
		Addr:         ":8080",
		Handler:      mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	fmt.Println("starting server at :8080")

	//нужные данные берутся из объекта структуры Server
	server.ListenAndServe()
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "any text and url: "+r.URL.String())
}
