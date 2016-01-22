package main 

import(
	"fmt"
	"math"
)

func MySqrt(f float64) (squareroot float64, ok bool){
	if f > 0 {
		squareroot, ok = math.Sqrt(f), true
	} else {
		squareroot, ok = 0, false
	}
	return squareroot, ok
}

func main() {
	for index := -2.0; index <= 10; index++ {
		squareroot, possible := MySqrt(index)
		if possible {
			fmt.Printf("The square root of %f is %f\n", index, squareroot)
		} else {
			fmt.Printf("Sorry, no square root for %f\n", index)
		}
	}

	var A [10]int
	var people [10]person

	third_name := perople[2].name 

	third_person := perople[2]
	thirdname := third_person.name



}