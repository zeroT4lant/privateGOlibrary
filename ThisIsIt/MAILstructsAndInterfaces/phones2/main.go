package main

import (
	"errors"
)

type Payer interface {
	Pay(int) error
}

type Wallet struct {
	Cash int
}

type Card struct {
	Number  string
	Balance int
	CVV     string
}

type ApplePay struct {
	AppleID int
	Balance int
}

func (w *Wallet) Pay(amount int) error {
	if w.Cash < amount {
		return errors.New("not enough money")
	}
	w.Cash -= amount
	return nil
}

func (c *Card) Pay(amount int) error {
	if c.Balance < amount {
		return errors.New("not enough money")
	}
	c.Balance -= amount
	return nil
}

func (a *ApplePay) Pay(amount int) error {
	if a.Balance < amount {
		return errors.New("not enough money")
	}
	a.Balance -= amount
	return nil
}

//-----TYPE ASSERTION----- Утверждение типа при помощи конструкции p,ok := in.(Type)
// func Buy(in interface{},amount int) {
// 	if p,ok := in.(Payer); !ok {
// 		fmt.Println("Не платёжное средство")
// 		return
// 	} else {
// 		p.Pay(amount)
// 	}
// }

//-----TYPE SWITCH----- Выбор типа с - switch case "switch p.(type)"
// func Buy(p Payer) {
// 	switch p.(type){
// 	case *Wallet:
// 		fmt.Println("Оплата наличными!")
// 	case *Card:
// 		card,ok := p.(*Card)
// 		if !ok {
// 			fmt.Println("Не удалось преобразовать к типу *Card")
// 		}
// 		fmt.Println("Вставляйте карту",card.Number)
// 	default:
// 		fmt.Println("Что-то новенькое!")
// 	}

// 	err := p.Pay(10)
// 	if err != nil {

// 	}
// }

func main() {

}
