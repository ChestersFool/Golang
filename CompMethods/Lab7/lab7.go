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

func task(a float64, b float64, N int) []float64 {
	x, h := a, (b-a)/float64(N)
	A := make([][]float64, N)
	for i := range A {
		A[i] = make([]float64, N)
	}
	free := make([]float64, N)

	// Перший є і останній from task
	A[0][0] = 1
	A[0][1] = 1 + (1/2.0)*h*p_func(a) // this from (3)

	A[N-1][N-1] = 0.5
	A[N-1][N-2] = 1 - (1/2.0)*h*p_func(b) // this from (1)

	free[0] = r_func(float64(a))

	x += h

	for i := 1; i < N-1; i++ {
		A[i][i-1] = 1 - (1/2.0)*h*p_func(x) // (1)
		A[i][i] = 2 - h*h*q_func(x)
		A[i][i+1] = 1 + (1/2.0)*h*p_func(x) // (3)

		free[i] = r_func(x)

		x += h
	}

	free[N-1] = r_func(x)

	for i, row := range A {
		fmt.Printf("row: %v || free: %v    i: %v\n", row, free[i], i)
	}

	// https://uk.wikipedia.org/wiki/%D0%9C%D0%B5%D1%82%D0%BE%D0%B4_%D0%BF%D1%80%D0%BE%D0%B3%D0%BE%D0%BD%D0%BA%D0%B8
	d_x, c_x, result_SLAR := make([]float64, N), make([]float64, N), make([]float64, N)

	c_x[0] = A[0][1] / A[0][0]
	d_x[0] = free[0] / A[0][0]

	for i := 1; i < N-1; i++ {
		c_x[i] = A[i][i+1] / (A[i][i] - A[i][i-1]*c_x[i-1])
		d_x[i] = (free[i] - A[i][i-1]*d_x[i-1]) / (A[i][i] - A[i][i-1]*c_x[i-1])
	}

	c_x[N-1] = 0
	d_x[N-1] = (free[N-1] - A[N-1][N-2]*d_x[N-2]) / (A[N-1][N-1] - A[N-1][N-2]*c_x[N-2])

	result_SLAR[N-1] = d_x[N-1]
	for i := N - 2; i >= 0; i-- {
		result_SLAR[i] = d_x[i] - c_x[i]*result_SLAR[i+1]
	}

	return result_SLAR
}

// variant 11
func main() {
	a, b, N := 0, 1, 10

	fmt.Printf("task(a, b, N): %v\n", task(float64(a), float64(b), N))
	fmt.Printf("task(a, b, N): %v\n", task(float64(a), float64(b), N * 2))

	// TODO стовпець вільних членівб порахувати СЛАР
}
