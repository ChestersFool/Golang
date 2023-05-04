package main

import "fmt"

func func_u(u, v float64) float64 {
	return (0.45 - 0.04*v) * u
}

func func_v(u, v float64) float64 {
	return (2.31 - 0.01*v + 0.5*u) * v
}

func main() {
	var u, v, a, b float64 = 5, 10, 0, 20
	var n float64 = 200
	var k1_u, k1_v, k2_u, k2_v float64
	
	h := (b - a) / n

	for i := 0; i < int(n); i++ {
		k1_u = func_u(u, v)
		k1_v = func_v(u, v)

		k2_u = func_u(u+h * k1_u/2.0, v+h*k1_v/2.0)
		k2_v = func_v(u+h * k1_u/2.0, v+h*k1_v/2.0)

		u = u + h * k2_u
		v = v + h * k2_v

		fmt.Printf("u: %v; i: %v\n", u, i)
		fmt.Printf("v: %v; i: %v\n\n", v, i)
	}
}


