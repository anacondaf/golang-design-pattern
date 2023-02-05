package main

import (
	"fmt"
	decorator "go-design-pattern/decorator-pattern/redShapeDecorator"
	"go-design-pattern/decorator-pattern/shape"
)

func main() {
	var circle = shape.Circle{}

	var redBorderCircle = decorator.NewRedShapeDecorator(shape.Circle{})
	var redBorderRectangle = decorator.NewRedShapeDecorator(shape.Rectangle{})

	var yellowBgCircle = decorator.NewYellowBackgroundShapeDecorator(shape.Circle{})

	fmt.Println("Circle with normal border")
	circle.Draw()

	fmt.Println("\nCircle of red border")
	redBorderCircle.Draw()

	fmt.Println("\nRectangle of red border")
	redBorderRectangle.Draw()

	fmt.Println("\nCircle of yellow background")
	yellowBgCircle.Draw()
}
