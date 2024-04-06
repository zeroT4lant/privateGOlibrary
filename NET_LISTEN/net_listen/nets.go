package main

import (
	"bufio"
	"fmt"
	"net"
)

func main() {
	//начинаем слушать по соединению
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	//когда соединение будет принято - обработаем его
	for {
		//возвращает имя-адресс подключения
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	name := conn.RemoteAddr().String()

	fmt.Println(name + "connected")
	//пишем в соединение что-то
	conn.Write([]byte("hello, " + name + " !"))

	defer conn.Close()

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		text := scanner.Text()
		if text == "Exit" {
			conn.Write([]byte("bye"))
			fmt.Println(name, "disconnected")
			break
		} else if text != "" {
			fmt.Println(name, "enters", text)
			conn.Write([]byte("you enter " + text))
		}
	}

}
