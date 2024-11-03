package main

import "fmt"

// Go doesnt' have class or object, but it has struct
type person struct {
	name string
	age  int
}

// this is a method of the person struct
// it's similar to method of Object Oriented Programming
func (p person) sayHello() {
	fmt.Printf("Hello, my name is %s and I'm %d", p.name, p.age)
}

func (p person) sayKoreanAge() {
	fmt.Printf("Hello, my name is %s and I'm %d in Korean age", p.name, p.age+2)
}

/* Pointer & Structs */
func main() {
	hyewon := person{"hyewon", 28}
	fmt.Println("hyewon's name : ", hyewon.name)
	fmt.Println("hyewon's age : ", hyewon.age)
	hyewon.sayHello()
	hyewon.sayKoreanAge()

	a := 2
	b := a // copy the value of a to b
	a = 12
	fmt.Println(b)      // the result will be 2
	fmt.Println(&a, &b) // this will print the memory address of a

	// then, how can we the real value of a?
	// we will gonna use real data of a from the memory, not the copy of it
	// the type of c is *int, which means it is a pointer to an integer
	c := &a            // c, a would be in the same place in the memory
	fmt.Println(c, &a) // this will print the memory address of a
	fmt.Println(*c)    // this will print the value of a
}
