package main

import "fmt"

func main() {
	fmt.Println("Pointers")
	a := 5
	b := &a

	// b contains the memory address where a is stored in the memory.

	fmt.Println(a, b, &a)

	fmt.Printf("type a: %T type b: %T\n", a, b)


	c := 10
	d := c

	// Shows that they are the same but not the same they are stored on different places in the memory

	fmt.Println(&c, &d)

	// Use * to read val from memory address
	fmt.Println(*b)

	// So &var gets the memory address from the variable.
	// And *var of memory address gets the value of it.

	// Change value with pointer
	// Because b contains the memory address of where a is stored and *b is the value.
	// You can change *b and that will change the value in a. 

	*b = 10
	fmt.Println(a)



}