package main
import "fmt"
func test() bool {
	fmt.Println("test.....")
	return true
}
func main() {
	var i int = 10 
	// if i > 9 && test() {
	// 	fmt.Println("hello")
	// }
	if i > 9 || test() {
		fmt.Println("ok")
	}
}