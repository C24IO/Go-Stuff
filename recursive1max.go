
package main 

import "fmt"

func Max(slice []int) int {
	if len(slice) == 1{
		return slice[0] //There's only one element in the slice, return it ! 
	}

	middle := len(slice)/2

	m1 := Max(slice[:middle])
	m2 := Max(slice[middle:])

	if m1 > m2 {
		return m1
	}
	return m2 
}

func main() {
	s := []int {1,2,3,4,4,5,3,2,1,1,1,3,0,345,2,34,324,34,32432,423,4}
	fmt.Println("Max(s) =", Max(s))
}
