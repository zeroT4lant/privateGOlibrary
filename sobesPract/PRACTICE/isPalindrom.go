package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	fmt.Println(isPalidrom("anna"))
	fmt.Println(fibonacci(10))
	checkLink("https://www.google.com")

	p1 := Person{"John", 20}
	data, _ := json.Marshal(p1)
	fmt.Println(string(data))

	var decodedPerson Person

	json.Unmarshal(data, &decodedPerson)
	fmt.Println(decodedPerson)
}

func checkLink(link string) bool {
	resp, err := http.Get(link)
	if err != nil {
		fmt.Println("Ошибка при проверке ссылки: ", err)
		return false
	}

	if resp.StatusCode != http.StatusOK {
		fmt.Println("Ссылка недоступна")
		return false
	}

	fmt.Println("Ссылка доступна")
	return true
}
