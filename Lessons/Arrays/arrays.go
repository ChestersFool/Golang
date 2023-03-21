package main

import (
	"math/rand"
	"fmt"
)

func main() {
	var arr [2]int
	for i := 0; i < len(arr); i++ {
		arr[i] = rand.Int() % 10
	}
	fmt.Printf("arr: %v\n", arr)

	slIce := make([]int, 5)
	slIce = append(slIce, 5, 7)
	fmt.Printf("slIce: %v\nlen: %v\ncap: %v\n", slIce, len(slIce), cap(slIce))

	names := [4]string{
		"Misha",
		"Alice",
		"Maks",
		"Teemo",
	}

	var names_slice []string = names[:2] // 0:2 :4 : 2:
	fmt.Printf("names_slice: %v\n", names_slice)
	names[0] = "Natali"
	fmt.Printf("names_slice: %v\n", names_slice)


}