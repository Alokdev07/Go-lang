package main

import "fmt"

// numbered sequence of specific length
func main(){
	// add zero in initial value
	var number [4]int
	fmt.Println(len(number))
	number[0] = 1
	number[1] = 2
	number[2] = 3
	number[3] = 4

	fmt.Println(number)

	// add false value
	var vals[2] bool 
	fmt.Println(vals)

	// add empty string in initial value
	var names[3] string
	fmt.Println(names)

	// to declare it in single line
	new_number := [3]int{1,2,3}
	fmt.Println(new_number)
	
	//2d array
	new_numbers := [2][2]int{{1,2},{3,4}}
	fmt.Println(new_numbers)

	// - fixed size,that is predictable
	// - memory optimization
	// - constant time access
}