package main

import "fmt"

func deferFunc1() {
	defer fmt.Println("This is a defered statement")
	fmt.Println("This statement is not defered")
}

func main() {
	deferFunc1()
}
