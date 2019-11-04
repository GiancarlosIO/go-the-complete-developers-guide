package main

import "fmt"

type shape interface {
	getArea() float64
}

type triangle struct {
	height float64
	base   float64
}

type square struct {
	sideLenght float64
}

func (t triangle) getArea() float64 {
	return 0.5 * t.height * t.base
}

func (s square) getArea() float64 {
	return s.sideLenght * s.sideLenght
}

func main() {
	s := square{
		sideLenght: 20,
	}
	t := triangle{
		height: 10,
		base:   25,
	}
	printArea(s)
	printArea(t)
}

func printArea(s shape) {
	fmt.Printf("%f\n", s.getArea())
}
