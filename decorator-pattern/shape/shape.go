package shape

import "fmt"

type Shape interface {
	Draw()
}

type Circle struct {
}

func (c Circle) Draw() {
	fmt.Println("Shape: Circle")
}

type Rectangle struct {
}

func (r Rectangle) Draw() {
	fmt.Println("Shape: Rectangle ")
}
