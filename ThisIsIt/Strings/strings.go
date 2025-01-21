package main

import (
	"fmt"
	"strings"
)

func main() {
	//СТРОКИ - ПАКЕТ strings
	str := "hello Almaz"
	//ищем определённую подстроку
	fmt.Println(strings.Contains(str, "Almaz"))

	//разделяет по определённому символу
	//Разделение по каждому символу
	sliceOfStr := strings.Split(str, "")
	fmt.Println(sliceOfStr)

	//склеивает строки из слайса в одну строку, с разделителем во втором аргументе
	fmt.Println(strings.Join([]string{"giga", "nigga"}, "-"))

	//stringBuilder, создаём объект структуры
	sb := strings.Builder{}

	sb.WriteString("yo")
	sb.WriteString("-")
	sb.WriteString("nigga")

	str2 := sb.String()
	fmt.Println(str2)
}
