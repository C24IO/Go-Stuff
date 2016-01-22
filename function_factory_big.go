package main 
import "fmt"

func isOdd(integer int) bool{
    if integer%2 == 0 {
        return false
    }
    return true
}

func isBiggerThan4(integer int) bool{
    if integer > 4 {
        return true
    }
    return false
}


/* Filter factory in essence takes in a function and creates another one 
of a completely diffrent type. 

Its taking in the function isOdd - that is good at pointing out Odd values 

And spitting out a function that will be operating on a slice and splitting 
that slice up using the isOdd function into two - yes and no slices. 

*/

func filter_factory(f func(int) bool) (func (s []int) (yes, no []int)) {
	return func(s []int) (yes, no []int) {
		for _, value := range s{
			if f(value){
				yes = append(yes, value)
			} else {
				no = append(no, value)
			}
		}
		return // here the return is empty because Go is going to figure out what to return.
	}
}

func main(){
    s := []int {1, 2, 3, 4, 5, 7}
    fmt.Println("s = ", s)
    odd_even_function := filter_factory(isOdd)
    odd, even := odd_even_function(s)
    fmt.Println("odd = ", odd)
    fmt.Println("even = ", even)

    //separate those that are bigger than 4 and those that are not.
    //this time in a more compact style.
    bigger, smaller := filter_factory(isBiggerThan4)(s)
    fmt.Println("Bigger than 4: ", bigger)
    fmt.Println("Smaller than or equal to 4: ", smaller)
}