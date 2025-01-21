package main

import (
	"errors"
	"fmt"
)

type Payer interface {
	Pay(int) error
}

type Wallet struct {
	Cash int
}

func (w *Wallet) Pay(amount int) error {
	if w.Cash < amount {
		return errors.New("not enough money")
	}
	w.Cash -= amount
	return nil
}

func Buy(p Payer) {
	err := p.Pay(10)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Спасибо за покупку через %T\n\n", p)
}

func main() {
	myWallet1 := &Wallet{Cash: 100}
	myWallet2 := &Wallet{Cash: 8}
	Buy(myWallet1)
	Buy(myWallet2)
}

// type Payer interface {
// 	Pay(int) error
// }

// type Wallet struct {
// 	Cash int
// }

// func (w *Wallet) Pay(amount int) error {
// 	if w.Cash < amount {
// 		return errors.New("not enough money")
// 	}
// 	w.Cash -= amount
// 	return nil
// }

// func Buy(p Payer) {
// 	err := p.Pay(10)
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Printf("Спасибо за покупку через %T\n\n", p)
// }
