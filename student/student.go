package student

import "fmt"

// if i want to export the struct, I have to capitalize the first letter of the struct
type Student struct {
	name string
	age  int
}

// it will not copy the struct, but it will use the real struct using pointer
func (s *Student) SetStudent(name string, age int) {
	s.name = name
	s.age = age
	fmt.Println("SetStudent : ", s)
}

// when do we use pointer?
// when we want to change the value of the struct, we have to use pointer
// if we don't use pointer, it will copy the struct and it will not change the real struct
func (s Student) ShowName() {
	fmt.Println("showName : ", s.name)
}
