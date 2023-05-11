package main

import "fmt"

type employee struct {
	firstname string
	lastName  string
}

type hr struct {
	employee
	designation string
}

func main() {

	h := hr{
		employee: employee{
			firstname: "tushar",
			lastName:  "sutar",
		},
		designation: "hr",
	}

	fmt.Println(h)

}
