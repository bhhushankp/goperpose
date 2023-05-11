package main

import "fmt"

//accesable only through the package not outside the packge
var a int = 23

//accesable out side the package not outside the packge
var A int = 45

func main() {
	//accesable only through block or function
	b := 34
	fmt.Println(b)
	fmt.Println(a)
	fmt.Println(add())
}

func add() int {
	return a
}
