package main

import "fmt"

func change_num(num int) {
	num = 5
	fmt.Println("in change num : ",num)
} // by value change

func change_num_ref(num *int){
	*num = 5
	fmt.Println("in change ref",*num)
}

func main() {
	num := 1
	change_num(num)


	fmt.Println("after change num in main",num)
	fmt.Println("memory address",&num)
	// there is no change in num because value passing so we need to parsing with reference
	change_num_ref(&num) // passing memory reference
	fmt.Println("after change num in main",num)

}