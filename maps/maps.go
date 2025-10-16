package main

import (
	"fmt"
	"maps"
)

//maps -> hash,object,dict
func main(){
	// creating map
	var create_map = make(map[string]string)
	// adding elements
	create_map["name"] = "golang"
	create_map["area"] = "backend"
	//Imp : if key does not exists in the map then it returns zero value
	fmt.Println(create_map["name"],create_map["area"])
	fmt.Println(len(create_map))
	delete(create_map,"area")
	fmt.Println(create_map)

	another_map := map[string]string{"price":"40","phone":"3"}
	fmt.Println(another_map)

	value,ok := another_map["price"] // value is the value and ok is the return value if it is exist or not
	if ok{
		fmt.Println("all ok")
	}else{
		fmt.Println("not ok")
	}
	fmt.Println(value)
	fmt.Println(maps.Equal(create_map,another_map))
}