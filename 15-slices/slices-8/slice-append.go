package main

import "fmt"

func printSlice(name string, slice []int) {
	fmt.Printf("The slice %s has length of %d and capacity of %d, and the slice itself is %v\n", name, len(slice), cap(slice), slice)
}
func main() {
	var slice1 []int
	printSlice("slice1", slice1)
	slice1 = append(slice1, 0)
	printSlice("slice1", slice1)
	slice1 = append(slice1, 1)
	printSlice("slice1", slice1)
	slice1 = append(slice1, 2, 3, 4, 5)
	printSlice("slice1", slice1)
}
