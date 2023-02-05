package decorator

import (
	"fmt"
	"go-design-pattern/decorator-pattern/shape"
)

type ShapeDecorator struct {
	shape shape.Shape
}

type RedBorderShapeDecorator struct {
	ShapeDecorator
}

func NewRedShapeDecorator(shape shape.Shape) *RedBorderShapeDecorator {
	return &RedBorderShapeDecorator{
		ShapeDecorator{
			shape,
		},
	}
}

func (r RedBorderShapeDecorator) Draw() {
	r.shape.Draw()
	r.drawRedBorder()
}

func (r RedBorderShapeDecorator) drawRedBorder() {
	fmt.Println("Border color: Red")
}

type YellowBackgroundShapeDecorator struct {
	ShapeDecorator
}

func NewYellowBackgroundShapeDecorator(shape shape.Shape) *YellowBackgroundShapeDecorator {
	return &YellowBackgroundShapeDecorator{
		ShapeDecorator{
			shape,
		},
	}
}

func (y YellowBackgroundShapeDecorator) Draw() {
	y.shape.Draw()
	y.drawYellowBackground()
}

func (y YellowBackgroundShapeDecorator) drawYellowBackground() {
	fmt.Println("Background color: Yellow")
}
