package main

import "fmt"

func main() {
	const N = 10
	years := [N]int{8, 9, 10, 11, 12, 13, 14, 15, 16, 17}
	y := [N]float64{68.27, 69.29, 70.44, 71.02, 71.15, 71.37, 71.38, 71.68, 71.98, 71.76}
	var sumX, sumY, sumXY, sumXSquare, sumSquareX float64 = 0, 0, 0, 0, 0

	for i := 0; i < N; i++ {
		sumXY += float64(years[i]) * y[i]
		sumX += float64(years[i])
		sumY += y[i]
		sumXSquare += float64(years[i]) * float64(years[i])
	}
	sumSquareX = sumX * sumX

	a := (N*sumXY - sumX*sumY) / (N*sumXSquare - sumSquareX)
	b := (sumY - a*sumX) / N
	fmt.Printf("a: %v\n", a)
	fmt.Printf("b: %v\n", b)
}
