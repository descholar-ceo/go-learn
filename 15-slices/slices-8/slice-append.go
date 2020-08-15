package main

import "fmt"

func printSlice(name string, slice []int) {
	fmt.Printf("The slice %s has length of %d and capacity of %d, and the slice itself is %v\n", name, len(slice), cap(slice), slice)
}
func main() {

}
