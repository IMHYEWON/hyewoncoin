package student

import "fmt"

// if i want to export the struct, I have to capitalize the first letter of the struct
type Student struct {
	name string
	age  int
}

func (s Student) SetStudent(name string, age int) {
	s.name = name
	s.age = age
	fmt.Println("SetStudent : ", s)
}
