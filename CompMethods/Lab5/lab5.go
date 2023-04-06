package main

import (
	"fmt"
	"math"
)

func function(number float64) float64 {
	return math.Sin(number + 2*math.Sqrt(number))
}

func KFTrapeze(a float64, b float64, n int) float64 {
	var result float64
	var h float64 = (b - a) / float64(n)
	var x float64 = a

	for i := 0; i < n; i++ {
		result += (function(x) + function(x+h)) / 2 * h
		x += h
	}

	return result
}

// now  variant 13 // need 16
func main() {
	fmt.Println("Hello, World!")
	fmt.Println(KFTrapeze(0, 10, 100))
}
