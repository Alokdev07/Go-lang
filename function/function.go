package main

import "fmt"

func add_two_integer(a , b int) int {
	return a + b
}

func get_languages()(string,string){
	return "golang","javascript"
}

func processIt(value int,fn func(a int) int) int{
	result := fn(value)
	return result
}

func another_function() func(a int) int{
	return func(a int) int{
		return  a
	}
}

func main() {
	var sum = add_two_integer(3, 5)
	fmt.Println(sum)
    first_language,second_language := get_languages()
	fmt.Println(first_language,second_language)
	fn := func(a int) int{
		return a*2
	}
	result := processIt(5,fn)
	fmt.Println(result)
	another_function()
}