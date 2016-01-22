package main
import "fmt"

func Max(slice []int) int {

	max := slice[0]

	for index, value := range slice {
		if value > max {
			max = value
		}
	}
	return max
}

func main() {

	A1 := []int {1,2,3,4,5,6,4,3,3233,2,2,2,1}
	A2 := [4]int {1,2,3,4}
	A3 := [1]int {1}

	var slice []int 

	slice = A1[:]
	slice = A2[:]
	slice = A3[:] 

}