package main

import "fmt"

func main() {
	i, j := 12, 18
	p, x := &i, &j
	*p = *p + 22
	*x = 45
	fmt.Println(i, j)
}
