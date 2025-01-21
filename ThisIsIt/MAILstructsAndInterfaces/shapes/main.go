package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float32
}

type Square struct {
	side float32
}

func (s Square) Area() float32 {
	return s.side * s.side
}

type Circle struct {
	radius float32
}

func (c Circle) Area() float32 {
	return c.radius * c.radius * math.Pi
}

func main() {
	circle := &Circle{
		radius: 3,
	}
	//fmt.Println(circle.Area())

	square := &Square{
		side: 2,
	}
	//fmt.Println(square.Area())

	//Circle и Square подходят под описание Shape.
	for _, shape := range []Shape{circle, square} {
		fmt.Println(shape.Area())
	}
}
