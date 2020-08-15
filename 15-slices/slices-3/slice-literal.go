package main

import "fmt"

func main() {
	slice1 := []int{1, 2, 3, 4, 5}
	fmt.Println("The slice one is : ", slice1)
	slice2 := []struct {
		val1 string
		val2 int
	}{
		{"emma", 1815},
		{"descholar", 999},
		{"descholar-ceo", 2100},
	}
	fmt.Println("The slice two is : ", slice2)
}
