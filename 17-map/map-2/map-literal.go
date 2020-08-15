package main

import "fmt"

type loca struct{ Lat, Long float64 }

func main() {
	myH := map[string]loca{
		"nezago": loca{100, 102},
		"home":   loca{101, 111},
		"office": loca{129, 144},
	}

	fmt.Println(myH)
}
