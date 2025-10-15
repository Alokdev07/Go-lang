package main

import (
	"fmt"
	"reflect"
	"time"
)

func main() {
	// simple switch
	i := 5
	switch i {
	case 1:
		fmt.Println("one")
	case 2:
		fmt.Println("two")
	case 3:
		fmt.Println("three")
	case 4:
		fmt.Println("four")
	default:
		fmt.Print("other\n")
	}

	// multiple condition switch

	switch time.Now().Weekday(){
	case time.Saturday,time.Sunday:
		fmt.Println("holiday")
	default:
		fmt.Println("it's work day")
	}

    //type switch
	whoAmI := func(i any) {
		switch i := i.(type) {
		case int:
			fmt.Println("it's an integer")
		case string:
			fmt.Println("it's a string")
		default:
			fmt.Println("other", reflect.TypeOf(i))
		}
	}

	// Call the function with sample arguments
	whoAmI(42)
	whoAmI("hello")
	whoAmI(3.14)
}