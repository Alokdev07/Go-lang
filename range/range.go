package main
import "fmt"

func main(){
	nums := []int{1,2,3,4}
	
	/*
	 for i:=0;i<len(num);i++{
	 	fmt.Println(num[i])
	 }
	   -> classical for loop  */
	   
	   for index,num:= range nums{
		fmt.Println(num,index)
	   }

    user_map := map[string]string{"first_name":"alok","second_name":"bhuyan"}
	for key,value := range user_map{
		fmt.Println(key,value)
	}

	name := "Alok"
	// unicode point rune
	// starting byte of rune 
	// 300 -> 1 byte,2 byte
	for index,unicode := range name{
		fmt.Println(index,string(unicode))
	}
}