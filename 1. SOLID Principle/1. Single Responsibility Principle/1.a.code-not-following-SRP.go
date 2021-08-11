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

func (s Square) draw() {
	fmt.Println("Code for drawing a square")
}

func (s Square) rotate() {
	fmt.Println("Code for rotating a square")
}

// Low cohesion functions calculateArea/calculatePerimeter and draw/rotate are belonged to same type i.e. Square
