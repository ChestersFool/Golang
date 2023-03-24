package main

import (
	"fmt"
	"math"
)

func MatrixVectorMultiplication(matrix [4][4]float64, vec [4]float64) [4]float64 {
	var result [4]float64

	for i, rows := range matrix {
		for j, elem := range rows {
			result[i] += elem * vec[j]
		}
	}

	return result
}

func VectorSubstraction(vec1 [4]float64, vec2 [4]float64) [4]float64 {
	var result [4]float64

	for i, elem := range vec1 {
		result[i] = elem - vec2[i]
	}

	return result
}

func VectorScalar(vec1 [4]float64, vec2 [4]float64) float64 {
	var result float64

	for i, elem := range vec1 {
		result += elem * vec2[i]
	}

	return result
}

func VectorAddition(vec1 [4]float64, vec2 [4]float64) [4]float64 {
	var result [4]float64

	for i, elem := range vec1 {
		result[i] = elem + vec2[i]
	}

	return result
}

func VectorNumberMultiplication(vec [4]float64, num float64) [4]float64 {
	var result [4]float64

	for i, elem := range vec {
		result[i] = elem * num
	}

	return result
}

func VectorNorm(vec [4]float64) float64 {
	var result float64 = vec[0]

	for _, elem := range vec {
		if math.Abs(elem) > result {
			result = math.Abs(elem)
		}
	}

	return result
}

func main() {
	A := [4][4]float64{
		{0.0625, -0.0379, -0.0713, 0.0932},
		{-0.0379, 0.0322, 0.0419, -0.0746},
		{-0.0713, 0.0419, 0.1018, -0.1204},
		{0.0932, -0.0746, -0.1204, 0.2317},
	}
	b := [4]float64{0.1528, -0.0985, -0.2474, 0.3872}
	x_k := [4]float64{0.1528, -0.0985, -0.2474, 0.3872}
	var r_k, p_k, q_k [4]float64
	var alfa_k, beta_k, norm float64
	const EPS float64 = 0.0001

	for i := 0; i < 4; i++ {
		fmt.Println(A[i])
	}

	// A * x_0
	r_k = MatrixVectorMultiplication(A, x_k)
	fmt.Printf("MatrixVectorMultiplication(A, b): %v\n", r_k)

	//  b - Ax_0
	r_k = VectorSubstraction(b, r_k)
	fmt.Printf("r_k: %v\n", r_k)

	p_k = r_k

	for {
		// q_k = A * p_k
		q_k = MatrixVectorMultiplication(A, p_k)
		fmt.Printf("q_k: %v\n", q_k)

		// alfa_k = (r_k, p_k) / (q_k, p_k)
		alfa_k = VectorScalar(r_k, p_k) / VectorScalar(q_k, p_k)
		fmt.Printf("alfa_k: %v\n", alfa_k)

		// x_k+1 = x_k + alfaK * p_k
		x_k = VectorAddition(x_k, VectorNumberMultiplication(p_k, alfa_k))
		fmt.Printf("x_k: %v\n", x_k)

		// r_k+1 = r_k - alfaK * q_k
		r_k = VectorSubstraction(r_k, VectorNumberMultiplication(q_k, alfa_k))
		fmt.Printf("r_k: %v\n", r_k)

		norm = VectorNorm(r_k)
		fmt.Printf("norm: %v\n", norm)

		if norm <= EPS {
			fmt.Printf("EPPP, x_k: %v\n", x_k)
			return
		}

		beta_k = VectorScalar(r_k, q_k) / VectorScalar(p_k, q_k)
		fmt.Printf("beta_k: %v\n", beta_k)

		p_k = VectorSubstraction(r_k, VectorNumberMultiplication(p_k, beta_k))
		fmt.Printf("p_k: %v\n\n-----------NEXT STEP------------\n\n", p_k)
	}
}
