package main

import (
	"fmt"
	"math"
)

type vertex struct {
	X, Y float64
}

func (v vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *vertex) Scale(f float64) {
	v.X = v.X * f
	v.Y = v.Y * f
}

func main() {
	v := vertex{3, 4}
	v.Scale(10)
	fmt.Println(v.Abs())
}
