package main

import "fmt"

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("When you double me, I become %v\n", v*2)
	case string:
		fmt.Printf("%q is %v bytes long\n", v, len(v))
	default:
		fmt.Printf("Oooh! Sorry, I don't know about that type named %T\n", v)
	}
}

func main() {
	do(2)
	do("Hello there! Welcome to the new wrld!")
	do(8.9)
}
