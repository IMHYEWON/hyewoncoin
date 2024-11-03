package main

import "fmt"

/* Pointer */
func main() {
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
