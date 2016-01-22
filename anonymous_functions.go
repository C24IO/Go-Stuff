package main 
import "fmt"

func main() {

	add1 := func(x int) int {
		return x+1
	}

	n := 6
	fmt.Println("n = ", n)
	n = add1(n)
	fmt.Println("add1(n) = ", n)
}