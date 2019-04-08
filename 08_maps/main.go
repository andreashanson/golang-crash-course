package main

import "fmt"

func main() {
	emails := make(map[string]string)

	// Assign kv
	emails["bob"] = "bob@gmail.com"
	emails["sharon"] = "sharon@gmail.com"
	fmt.Println(emails)
	fmt.Println(emails["bob"])

	emails2 := make(map[int]string)
	emails2[2] = "bob@gmail.com"
	emails2[3] = "sharon@gmail.com"
	
	fmt.Println(emails2)
	fmt.Println(emails2[3])

	fmt.Println(len(emails2))

	emails3 := make(map[int]string)
	emails3[2] = "bob@gmail.com"
	emails3[3] = "sharon@gmail.com"
	emails3[5] = "mike@gmail.com"

	
	fmt.Println(emails3[3])
	fmt.Println(emails3)
	fmt.Println(len(emails3))
	fmt.Println("Delete Mike")
	delete(emails3, 5)
	fmt.Println(emails3)
	fmt.Println(len(emails3))

	emails4 := map[string]string{"Andreas": "andreas@gmail.com", "Anna": "anna@gmail.com"}
	fmt.Println(emails4)
	fmt.Println(emails4["Andreas"])
	fmt.Println(emails4["Anna"])

	emails4["Trulsy"] = "trulsy@gmail.com"

	fmt.Println(emails4)
	fmt.Println(emails4["Trulsy"])
}