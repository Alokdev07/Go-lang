package main

import "fmt"

func sum(nums ...int) int{
	total := 0
	for _,val := range nums{
		total = val+total
	}
	return total
}

func main() {
	fmt.Println(1,2,3,4,5)
	result := sum(1,2,3,4)
	fmt.Println(result)
}