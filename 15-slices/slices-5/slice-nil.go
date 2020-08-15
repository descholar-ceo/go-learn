package main

import "fmt"

func main() {
	var slice []int
	fmt.Println(slice, cap(slice), len(slice))
	if slice == nil {
		fmt.Println("This slice is nil!")
	}
}
