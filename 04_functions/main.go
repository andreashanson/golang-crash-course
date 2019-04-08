package main

import "fmt"

func greetings(name string) string {
	return name
}

func getSum(num1 int, num2 int) int {
	return num1 + num2
}

func main() {
	fmt.Println(greetings("Anna"))
	fmt.Println(getSum(2, 44))
}