package main

import "fmt"

type loca struct {
	Lat, Long float64
}

func main() {
	m := make(map[string]loca)
	m["kigali butare"] = loca{2212, 234}
	fmt.Println(m["kigali butare"])
}
