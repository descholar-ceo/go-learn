package main

import "fmt"

type person struct {
	names string
	age   int
}

func (p person) talk() {
	fmt.Printf("Hello %s I noticed that are %d Y.O\n", p.names, p.age)
}

func main() {
	me := person{"descholar", 30}
	me.talk()
}
