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
	}
}
