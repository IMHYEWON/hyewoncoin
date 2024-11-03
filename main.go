package main

import (
	"fmt"
)

const str string = "hi"

// should be specified the type of return value (int)
func plus(a int, b int) int {
	return a + b
}

func plus2(a, b int, name string) (int, string) {
	return a + b, name
}

func plus3(a ...int) int {
	var total int
	// compiler will ignore index, if you write _ (underscore)
	for _, item := range a {
		total += item
	}
	return total
}

func main() {
	// var name string = "hyewon"
	// compiler will infer the type of variable automatically
	name := "hyewon"
	age := 28

	// update variable
	name = "haewon"

	result := plus(1, 2)

	result2, name2 := plus2(1, 2, "hyewon")

	result3 := plus3(1, 2, 3, 4, 5, 6, 7, 8, 9, 10)

	fmt.Println(name, age)
	fmt.Println(result)
	fmt.Println(result2, name2)
	fmt.Println(result3)

	sentence := "Hi I'm Hyewon"
	for index, letter := range sentence {
		fmt.Println(index, letter)  // it will print byte code
		fmt.Println(string(letter)) // it will print string
	}

	x := 23214345
	fmt.Printf("%b\n", x) // print x in binary
	fmt.Printf("%o\n", x) // print x in octal
	fmt.Printf("%x\n", x) // print x in hexadecimal
	fmt.Printf("%U\n", x) // print x in unicode

	xAsBinary := fmt.Sprintf("%b", x) // return x in binary as string
	fmt.Println(x, xAsBinary)         // print x in decimal and binary

	/* Slices And Arrays */
	// Arrays in GO are fixed length
	// Slices are dynamic arrays

	// Array
	foods := [3]string{"apple", "banana", "cherry"}
	for _, food := range foods {
		fmt.Println(food)
	}

	for i := 0; i < len(foods); i++ {
		fmt.Println(foods[i])
	}

	// Slice
	// Slice doesn't have limit
	foodsSlice := []string{"apple", "banana", "cherry"}
	foodsSlice = append(foodsSlice, "donut")
	foodsSlice = append(foodsSlice, "egg", "fish")

	served := foodsSlice[1:3] // it will return a slice of foodsSlice from index 1 to 3
	fmt.Println(foodsSlice)
	fmt.Println(served)

}
