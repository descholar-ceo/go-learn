package main

import "fmt"

type myStru struct {
	fname string
	lname string
	email string
}

func main() {
	user := myStru{"descholar", "emma", "descholar@nezago.com"}
	fmt.Println("First name: ", user.fname)
	fmt.Println("Last name: ", user.lname)
	fmt.Println("Email: ", user.email)
}
