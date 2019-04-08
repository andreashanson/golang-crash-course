package main
import "fmt"
// Range is for looping through arrays maps slices etc
func main() {
	fmt.Println("Range\nd")
	idsSlice := []int{33, 44, 55, 66}
	fmt.Println(idsSlice)
	fmt.Printf("%T\n", idsSlice)

	idsArray := [4]int{22, 77, 88, 99}
	fmt.Println(idsArray)
	fmt.Printf("%T\n", idsArray)

	// Loop through ids
	for i := 0; i <= len(idsSlice)-1; i++ {
		fmt.Println("ID")
		fmt.Println(idsSlice[i])
	}

	// Loop through ids with for range
	for index, id := range idsSlice {
		//fmt.Println(id)
		//fmt.Println(index)
		fmt.Printf("%d - ID: %d\n", index, id)
	}

	// Loop but not use the index.
	for _, id := range idsSlice {
		fmt.Printf("ID: %d\n", id)
	}

	sum := 0
	for _, id := range idsSlice {
		sum += id
	}
	fmt.Printf("Summan av alla ids är: %d", sum)
	fmt.Println("Summan av alla ids är:", sum)

	phoneNumbers := map[string]string{"anna": "9999", "andreas": "4444"}

	for key, value := range phoneNumbers {
		fmt.Printf("%s: %s\n", key, value)
	}

}