package main

import "fmt"

func main() {
	fmt.Println("Hello World")

	var arr []int
	arr = append(arr, 1, 2, 3)

	for i, elem := range arr {
		fmt.Printf("index: %d; elem: %d\n", i, elem)
	}
	for _, elem := range arr {
		fmt.Printf("elem: %d\n", elem)
	}

	for i := range arr {
		fmt.Printf("index: %d\n", i)
	}
}
