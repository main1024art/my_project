package main
import "fmt"
func main() {
	var i int = 20
	fmt.Println("i的地址是=",&i)
	var ptr *int = &i
	fmt.Println(ptr)
	*ptr = 10
	fmt.Println(i)
}