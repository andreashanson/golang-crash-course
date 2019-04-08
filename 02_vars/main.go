package main
import "fmt"

func main() {
	var name string = "Andreas Hansson"
	var age int32 = 36
	name2, age2 := "Andreas", 36
	//fmt.Printf("Hello, Me\n" + name + "\n")
	fmt.Println(name2, age2)
	fmt.Printf("%T\n", name)
	fmt.Printf("%T\n", age)
	fmt.Printf("%T\n", name2)
	fmt.Printf("%T\n", age2)
}
