package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	ID       int
	Username string
	Phone    string
}

var jsonStr = `{"id":"2", "username": "Ashot", "phone":"iphone"}`

var invalidJson = `[
	{"id": 17, "username": "iivan", "phone":0},
	{"id":"17", "address": "none", "company":"Mail.ru"}
]`

func main() {

	u1 := User{}

	//указываем в аргументе куда надо распаковать
	//1 аргумент - откуда достать, 2 куда положить
	err := json.Unmarshal([]byte(jsonStr), &u1)
	if err != nil {
		fmt.Errorf("pipec %w", err)
	}

	fmt.Println(u1)

	u1.Phone = "sasung"
	//тут сразу в переменную пакуем
	res, _ := json.Marshal(u1)
	fmt.Println(string(res))

	u2 := UserConstruct(1, "cumshot", "iphone")
	fmt.Println(u2)

	//для работы с неопределённым/неправильным JSON будем использовать пустой интерфейс
	data := []byte(invalidJson)

	//пустой интерфейс куда мы можем положить что захотим
	var user1 interface{}

	//складываем в него
	err = json.Unmarshal(data, &user1)
	fmt.Println(user1)

}

func UserConstruct(id int, username, phone string) User {
	return User{
		ID:       id,
		Username: username,
		Phone:    phone,
	}
}
