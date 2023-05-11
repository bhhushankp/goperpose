package main

import "fmt"

type student struct {
	id       int
	lastName string
	name     string
	marks    []int
}

type extstudent struct {
	name     string
	lastName string
}

type info interface {
	studentinfo()
}

func (s extstudent) studentinfo() {
	fmt.Println("we are inside the method", s.name, s.lastName)
}

func (s student) studentinfo() {
	fmt.Println("we are inside the method", s.id, s.name, s.marks, s.lastName)
}

func printdata(s student) {
	fmt.Println("we are inside the function", s.id, s.name, s.marks, s.lastName)
}

func main() {

	employee := student{
		id:   1,
		name: "tushar",
		marks: []int{
			56,
			67,
			78,
		},
	}

	s1 := extstudent{
		name:     "durga",
		lastName: "patil",
	}
	stu := []info{employee, s1}

	for _, v := range stu {
		v.studentinfo()

	}
	printdata(employee)

	emp := student{}
	emp.id = 23
	emp.name = "tushar"
	emp.marks = []int{23, 23}
	fmt.Println(emp)

}
