package main

import (
	"fmt"
	"net/http"
)

// создаём чтобы при обработке хендлера чтобы могли пользоваться полями структуры
type Handler struct {
	Name string
}

// по сути обычный хэндлер, но кастомизируем за счёт структур
func (h *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Name:", h.Name, "URL:", r.URL.String())
}

func main() {
	testHandler := Handler{"test"}
	http.HandleFunc("/test/", testHandler.ServeHTTP)

	rootHandler := Handler{"root"}
	http.HandleFunc("/", rootHandler.ServeHTTP)

	fmt.Println("starting server at :8080")
	http.ListenAndServe(":8080", nil)
}
