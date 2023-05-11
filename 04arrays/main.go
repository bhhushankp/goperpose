package main

import "fmt"

func main() {
	grades1 := 45
	grades2 := 56
	grades3 := 89
	fmt.Printf(" Grades : %v ,%v, %v \n", grades1, grades2, grades3)

	//now array
	grades := [3]int{45, 56, 89}
	fmt.Printf("Grades :%v", grades)

	//
	var student [3]string
	fmt.Println(student)

	a := []int{23, 3}
	b := a
	fmt.Println(a)
	fmt.Println(b)
	b[0] = 2
	fmt.Println(a)
	fmt.Println(b)
}
