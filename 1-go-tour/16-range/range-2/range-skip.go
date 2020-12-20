package main

import "fmt"

func main() {
	slice := make([]int, 6)
	fmt.Println(slice)
	for i := range slice {
		fmt.Printf("Element number %d is equal to %d\n", i+1, slice[i])
	}
	for _, v := range slice {
		fmt.Printf("The current element is equal to %d\n", v)
	}
}
