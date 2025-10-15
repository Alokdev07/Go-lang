package main

import "fmt"

// for -> only construct in go for looping

func main(){
	// i := 1
	// for i<= 3{
	// 	fmt.Println(i)
	// 	i = i+1
	// } // while loop in using for in go

	// infinite loop
	// for {
	// 	fmt.Println(1)
	// }

	// classic for loop

	for i := 0; i<3; i++{
		fmt.Println(i)
	}

	for i:= range 3{
		fmt.Println(i)
	} // another way of using loop

	for i:=0;i<5;i++{
		if i == 4{
			break
		}
		if i == 3{
			continue
		}
		fmt.Println(i)
	} // this is how we use break and continue in go

	

}