package main

import "fmt"

type Square struct {
	side int
}

func (s Square) calculateArea() int {
	return s.side * s.side
}

func (s Square) calculatePerimeter() int {
	return 4 * s.side
}

type SquareUI struct {
	color string
}

func (su SquareUI) draw() {
	fmt.Println("Code for drawing a square of color ", su.color)
}

func (su *SquareUI) rotate() {
	fmt.Println("Code for rotating a square of color ", su.color)
}

// Now the overall cohesion increase as we separate out the functions
