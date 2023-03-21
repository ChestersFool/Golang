package main

import "fmt"

type Point struct {
	X int
	Y int
}

func main() {
	pointsMap := map[string]Point{"2": {2, 3}}
	otherPointsMap := make(map[string]Point)
	pointsMap["a"] = Point{2, 4}
	point1 := Point{4, 5}

	fmt.Printf("pointsMap: %v\n", pointsMap)
	fmt.Printf("pointsMap[\"a\"]: %v\n", pointsMap["a"])

	otherPointsMap["1"] = point1
	fmt.Printf("otherPointsMap: %v\n", otherPointsMap)
}
