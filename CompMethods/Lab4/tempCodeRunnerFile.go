package main

import "fmt"

func main() {
	const N = 13
	years := [N]int{1000, 2000, 3000, 4000, 5000, 6000, 7000, 8000, 9000, 10000, 11000, 12000, 13000}
	y := [N]float64{3900, 5800, 7200, 5050, 5280, 7320, 11260, 12860, 16970, 17280, 21700, 21470, 23470}
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
	for i, elem := range y {
		fmt.Printf("%v: %v\n", i, elem)
	}
	fmt.Printf("a: %v\n", a)
	fmt.Printf("b: %v\n", b)
}
