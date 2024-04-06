package Errorsi


interface error{
	Error() string
}

type NewError struct {
	text string
}

func (e *NewError) Error() string {
	return e.text
}

func main()  {
	foo()
}

// Ошибка
func foo() {
	println(handle())


	//errors.New()
}



func handle() error {
	return &NewError{text: "ошибка"}
}
