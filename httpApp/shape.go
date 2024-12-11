package main

import "fmt"

type triangle struct {
	height float64
	base   float64
}

type square struct {
	sideLength float64
}

type shape interface {
	getArea() float64
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

func main() {
	//for square
	sq := square{sideLength: 5.6}
	printArea(sq)
	tri := triangle{height: 12.0, base: 14.0}
	printArea(tri)
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}