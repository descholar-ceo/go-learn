package main

import (
	"fmt"
	"math"
)

func power1next(x, y, z float64) float64 {
	if res := math.Pow(x, y); z < res {
		return res
	}
	return z
}

func main() {
	fmt.Println(
		power1next(2, 2, 3),
		power1next(3, 2, 10),
		power1next(3, 3, 20),
	)
}
