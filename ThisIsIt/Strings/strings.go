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
	sliceOfStr := strings.Split(str, "")
	fmt.Println(sliceOfStr)

	//склеивает строки из слайса в одну строку, с разделителем во втором аргументе
	fmt.Println(strings.Join([]string{"hello", "bot"}, "-"))

	//stringBuilder
	sb := strings.Builder{}

	sb.WriteString("yo")
	sb.WriteString("-")
	sb.WriteString("nigga")

	str2 := sb.String()
	fmt.Println(str2)
}
