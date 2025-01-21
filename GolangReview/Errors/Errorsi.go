package main

import "fmt"

type NewError struct {
	msg string
}

func (e *NewError) Error() string {
	return e.msg
}

func (e *NewError) handleError() error {
	return &NewError{
		e.Error(),
	}
}

func main() {
	newErr := &NewError{msg: "nigger error"}
	fmt.Println(newErr.handleError())
}
