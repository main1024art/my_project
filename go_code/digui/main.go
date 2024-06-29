package main
import "fmt"
func f(n int) int {
	if n == 10 {
		return 1
	} else  {
		return (2 * f(n + 1) + 2)
	}
}
func main() {
	fmt.Println(f(1))
 
}