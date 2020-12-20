package main

import "fmt"

func main() {
	fmt.Println("Starting counting...")
	for i := 0; i <= 10; i++ {
		defer fmt.Println(i)
	}
	fmt.Println("Done")
}
