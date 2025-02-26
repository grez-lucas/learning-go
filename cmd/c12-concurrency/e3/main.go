package main

import (
	"fmt"
	"math"
)

func buildMap() map[int]float64 {
	sqrMap := make(map[int]float64, 100_000)
	for i := 0; i < 100_000; i++ {
		sqrMap[i] = math.Sqrt(float64(i))
	}
	return sqrMap
}

func main() {
	sqrMap := buildMap()

	for i := 1_000; i < len(sqrMap); i += 1_000 {
		fmt.Printf("Sqrt of %d is %f\n", i, sqrMap[i])
	}
}
