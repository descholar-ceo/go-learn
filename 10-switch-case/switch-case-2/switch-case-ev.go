package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("When is Saturday?")
	today := time.Now().Weekday()
	switch time.Saturday {
	case today + 0:
		fmt.Println("It's today")
	case today + 1:
		fmt.Println("It's in one days")
	case today + 2:
		fmt.Println("It's in two days")
	default:
		fmt.Println("It's far from today")
	}
}
