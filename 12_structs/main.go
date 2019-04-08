package main

import (
	"fmt"
	"strconv"
)


// Define Person struct

type Person struct {
	//firstName string
	//lastName string
	//city string
	//gender string
	//age int

	// Shorter way of doing above.
	firstName, lastName, city, gender string
	age int
}

func (p Person) greet() string {
	return "Hello, my name is " + p.firstName + " " + p.lastName + " and I am " + strconv.Itoa(p.age) + " years old"
}


func (p *Person) hasBirthday() {
	p.age++
}

func (p *Person) setFirstName(name string) {
	p.firstName = name
}

func (p *Person) setLastName(name string) {
	p.lastName = name
}


func (p *Person) getMarried(spouseLastName string) {
	if p.gender == "Male" {
		fmt.Println("Did not change because of Male")
		return
	} else {
		p.lastName = spouseLastName
	}
}

func main() {
	fmt.Println("Structs")



	// Init person using struct
	person1 := Person{firstName: "Andreas", lastName: "Hansson", city: "Copenhagen", gender: "Male", age: 36}
	fmt.Println(person1)

	// Alternative way of init a person using struct
	person2 := Person{"Anna", "Lindeblad", "Copenhagen", "Female", 37}
	fmt.Println(person2)
	fmt.Println(person2.firstName)

	person2.hasBirthday()
	fmt.Println(person2.greet())

	person1.setFirstName("Andreas")
	person1.setLastName("Lindeblad")
	fmt.Println(person1.greet())

	person2.getMarried("Hansson")
	fmt.Println(person2.greet())

	person1.getMarried("Lindeblad")
	fmt.Println(person1.greet())

}