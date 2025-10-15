package main
import "fmt"
const age int = 20
func main(){
	const name string = "golang"
	fmt.Println(name,age)
	const (
		port = 5000
		host = "localhost"
	) // this is called constant group
	fmt.Println(port,host)
}