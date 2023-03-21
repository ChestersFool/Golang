package main

import (
	"fmt"
)

type Point struct {
	X int
	Y int
}

func (p Point) Foo() {
	fmt.Printf("p: %v\n", p)
}

func main() {
	// USE GO: FILL STRUCT
	first := Point{X: 0, Y: 0}
	first.Foo()
}
