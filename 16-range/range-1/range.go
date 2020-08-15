package main

import "fmt"

func main() {
	pow := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for i, v := range pow {
		fmt.Printf("index is %d and value is %v\n", i, v)
	}
}
