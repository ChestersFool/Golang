package main

import (
	"fmt"
)

type Point struct {
	x int
	y int
}

func (p Point) Foo() {
	fmt.Printf("p: %v\n", p)
}

func main() {
	defer fmt.Println("I would be in end")
	fmt.Println("Hello World!")

	name := "Misha"
	// var (
	// 	name3 = "mis"
	// 	age = "10
	// )
	// var name1, age1 = "Misha", 4

	c := fmt.Sprintf("My name is %s", name)
	fmt.Printf("c: %v\n", c)
	fmt.Println(test(1, 2, 3))

	// USE GO: FILL STRUCT
	first := Point{x: 0, y: 0}
	first.Foo()
}

func test(x, y, z float64) (a, b, c float64) {
	a = x
	b = y
	c = z
	return
}
