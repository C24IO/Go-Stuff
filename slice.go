package main 
import "fmt"

// Return the biggest value in a slice of ints
func Max(slice []int) int {
	max := slice[0]

	for index := 1; index < len(slice); index++ {
		if slice[index] > max {
			max = slice[index]
		}
	}
	return max
}

func main() {

	//Declare three arrays of different sizes, to test the function Max. 
	A1 := []int {1,2,3,4,5,6,7,87,9,66,66,66,6666}
	A2 := [4]int {1,2,3,4}
	A3 := [1]int {1}

	//Declare a slice of ints 

	var slice []int

	slice = A1[:]
	fmt.Println("The biggest value of A1 is", Max(slice))
	slice = A2[:]
	fmt.Println("The biggest value of A2 is", Max(slice))
	slice = A3[:]
	fmt.Println("The biggest value of A3 is", Max(slice))

}