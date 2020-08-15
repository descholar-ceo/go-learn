package types

import "fmt"

func main() {
	i := 12
	j := 12 + 2i
	x := 12.3
	y := "Some people are just funny!"
	z := 'c'
	fmt.Printf("The type of i is %T:\n", i)
	fmt.Printf("The type of j is %T:\n", j)
	fmt.Printf("The type of x is %T:\n", x)
	fmt.Printf("The type of y is %T:\n", y)
	fmt.Printf("The type of z is %T:\n", z)
}
