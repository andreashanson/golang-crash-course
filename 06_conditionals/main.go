package main

import "fmt"

func main() {
	x := 5
	y := 5

	if x < y {
		fmt.Printf("%d is less than %d\n", x, y)
	} else if x == y {
		fmt.Println("EQUAL")
	} else {
		fmt.Println("")
		fmt.Printf("%d is more than %d\n", x, y)
	}

	//Swithc
	color := "blu"

	switch color {
	case "red":
		fmt.Println("color is red")
	case "blue":
		fmt.Println("color is blue")
	default:
		fmt.Println("Color is NOT blue or red")
	}
}