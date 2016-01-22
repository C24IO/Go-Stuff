package main 
import . "fmt"

type Text string 

func main() {

	var name Text = "Ellie"



	Printf("name = %v stored at %v\n", name, &name)

	printthem()
	printAddress()

	var x = 10

	if x > 10 {
		Println("x is greater than 10")
	} else {
		Println("x is less than 10 ")
	}

	if x:= returnBig(); x > 10 {
		Println("x is greater than 10")
	} else {
		Println("x is less than 10")
	}


	for index := 0; index < 10 ; index++ {
		x += index
		Printf(" %v ", x)
	}

	//for {} - this is an infinite loop 

}

const(
	x = iota
	y = iota 
	z = iota
	a
	)

func printthem() {

	Println("Values - ", x, y, z, a)

}

func printAddress() {

	Printf("We got a variable - %v\n", new(Text))

}

func returnBig() int {
	return 20
}

/*func funcname(input1 type1, input2 type2, input3 type3) (output1 type1, output2 type2) {
	
return value1, value2

}*/

