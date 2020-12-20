package main

import "fmt"

func main() {
	fmt.Println("Enter number #1: ")
	var firstNum int
	fmt.Scanln(&firstNum)
	fmt.Println("Enter number #2: ")
	var secondNum int
	fmt.Scanln(&secondNum)
	fmt.Printf("The sum of %v and %v is %v\n", firstNum, secondNum, firstNum+secondNum)
}
