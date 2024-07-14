package main

import (
	"fmt"
	"math"
)

func sqrt(x float64) float64 {
    z := 1.0
    epsilon := 1e-10
    for {
        difference := math.Abs(z*z - x)
        if difference < epsilon {
            break
        }
        z -= (z*z - x) / (2*z)
        fmt.Printf("The new z value is: %f \n", z)
    }
    return z
}

func main()  {
	fmt.Printf("===>The final solution is: %f", sqrt(9))
}