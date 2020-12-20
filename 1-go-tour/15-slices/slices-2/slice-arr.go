package main

import "fmt"

func main() {
	arr := [5]string{
		"emma",
		"descholar",
		"fname",
		"lname",
		"nezago",
	}
	fmt.Println("The original array is : ", arr)
	a := arr[0:2]
	b := arr[0:5]
	fmt.Println("The slice a is : ", a)
	fmt.Println("The slice b is : ", b)
	b[1] = "mugirase"
	fmt.Println("The new slice a is : ", a)
	fmt.Println("The new slice b is : ", b)
	fmt.Println("The new array arr is : ", arr)
}
