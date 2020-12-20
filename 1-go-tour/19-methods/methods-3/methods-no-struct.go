package main

import (
	"fmt"
	"math"
)

type mFloat float64

func (m mFloat) divideWith(divider int) mFloat {
	return (mFloat(divider) / m)
}

func main() {
	num := mFloat(math.Pi)
	fmt.Printf("The result is : %f\n", num.divideWith(4))
}
