package main

import "fmt"

type person struct {
	names string
	age   int
}

func talk(p person) {
	fmt.Printf("Hello %s, I noticed that you are %d Y.O", p.names, p.age)
}

func main() {
	me := person{"descholar", 30}
	talk(me)
}
