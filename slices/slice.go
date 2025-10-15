package main

import (
	"fmt"
	"slices"
)

// slice -> dynamic
// most used concept in go
// useful methods

func main(){
	// uninitialized slice is nil-> null
	var number []int
	fmt.Println(number == nil,len(number))

	var another_number = make([]int,0,10) // second argument is initial size and third parameter is capacity
	//capacity => maximum numbers of elements can fit
	fmt.Println(cap(another_number))
	another_number = append(another_number, 1)
	another_number = append(another_number, 2)
	another_number = append(another_number, 3)
	another_number = append(another_number, 4)
	another_number = append(another_number, 5)
	another_number = append(another_number, 6)
	another_number = append(another_number, 7)
	another_number = append(another_number, 8)
	another_number = append(another_number, 9)
	fmt.Println(another_number,cap(another_number))

	// copy function
	another_number = append(another_number, 12)
	var copy_number = make([]int,len(another_number))
	
	copy(copy_number,another_number)
	fmt.Println(copy_number,another_number)

	// slice operator
	var another_slice = []int{1,2,3}
	fmt.Println(another_slice[0:2]) // slice operator

	// slice package
	var slice_package = []int{1,2}
	var another_slice_package = []int{1,2}
	fmt.Println(slices.Equal(slice_package,another_slice_package))
	slices.Reverse(slice_package)
	fmt.Println(slice_package)
	slice_package = slices.Replace(slice_package, 0, 1,3)
	fmt.Println(slice_package)
}