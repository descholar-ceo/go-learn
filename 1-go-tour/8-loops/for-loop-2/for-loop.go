package main

import "fmt"

func main() {
	i := 4
	for i < 10 {
		fmt.Println("The current value of i is:", i)
		i += i
	}
	fmt.Println("The final value of i is:", i)
}
