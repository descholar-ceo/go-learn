package main

import "fmt"

func main() {
	myChan := make(chan int, 4)
	myChan <- 3
	myChan <- 4
	myChan <- 40
	myChan <- 400
	fmt.Println(<-myChan)
	fmt.Println(<-myChan)
	fmt.Println(<-myChan)
	fmt.Println(<-myChan)
}
