package main

import "fmt"

func main() {
	// Arrays
	var fruitArr [2] string
	
	//Assign values
	fruitArr[0] = "Apple"
	fruitArr[1] = "Orange"
	fmt.Println(fruitArr)

	// Declare and assign values
	fruitArr2 := [2] string {"Apple", "Orange"}
	fmt.Println(fruitArr2)


	fruitArr3 := [3] string {"Apple", "Orange", "Banana"}
	fmt.Println(fruitArr3)

	fruitArr4 := [] string {"Apple", "Orange", "Banana", "Pears"}
	fmt.Println(fruitArr4)

	fmt.Println(len("DDD"))
	fmt.Println(len(fruitArr4))

	// Slice
	// Starts on index 2 and stops before index 3
	fmt.Println(fruitArr4[2:3]) // returns banana
	fmt.Println(fruitArr4[1:3]) // returns Orange Banana

	// Get first to in a new array
	var firstTwo [] string = fruitArr4[0:2]
	fmt.Println(firstTwo)

}