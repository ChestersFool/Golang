package main

import (
	"fmt"
	"math"
)

func p_func(x float64) float64 {
	return -(x + 1.0)
}

func q_func(x float64) float64 {
	return -1.0
}

func r_func(x float64) float64 {
	return 2.0 / math.Pow(x+1.0, 3)
}

// variant 11
func main() {
	const N int = 10
	a, b := 0, 1
	x, h := float64(a), float64(b-a)/float64(N)
	var A [N][N]float64

	// Перший є і останній from task
	A[0][0] = 1
	A[N-1][N-1] = 0.5
	for i := 1; i < N-1; i++ {
		A[i][i-1] = 1 - (1/2.0)*h*p_func(x)
		A[i][i] = 2 - h*h*q_func(x)
		A[i][i+1] = 1 + (1/2.0)*h*q_func(x)

		x += h
	}
	for _, row := range A {
		fmt.Printf("row: %v\n", row)
	}

	// TODO стовпець вільних членівб порахувати СЛАР
}
