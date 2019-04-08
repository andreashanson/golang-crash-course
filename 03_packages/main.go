package main
import (
	"fmt"
	"math"
	"github.com/andreashanson/go_crash_course/03_packages/util"
)

func greeting(name string) string {
	return name
}

func main() {
	fmt.Println("ANDREAS")
	var isumber bool = math.IsNaN(34)
	var test_floor float64 = math.Floor(3.5)
	var test_ceil float64 = math.Ceil(3.5)
	var test_sqrt float64 = math.Sqrt(17)
	var round_down float64 = math.Floor(test_sqrt)
	fmt.Println(test_floor)
	fmt.Println(test_ceil)
	fmt.Println(test_sqrt)
	fmt.Println(round_down)
	fmt.Println(isumber)
	fmt.Println(util.Reverse("andreas"))
	fmt.Println(greeting("Anna"))
}
