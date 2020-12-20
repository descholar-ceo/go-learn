package main

import "fmt"

func printSlice(name string, slice []int) {
	fmt.Printf("The slice %s has length of %d and capacity of %d, and the slice itself is %v\n", name, len(slice), cap(slice), slice)
}

func main() {
	slice1 := make([]int, 5)
	slice2 := make([]int, 4, 6)
	slice3 := slice1[:3]
	slice4 := slice2[1:3]
	printSlice("slice1", slice1)
	printSlice("slice2", slice2)
	printSlice("slice3", slice3)
	printSlice("slice4", slice4)
}
