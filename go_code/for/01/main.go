package main
import "fmt"
func main() {
	var f int = 10
	for i := 1; i <= f; i++ {
		// for q := 1; q <= 2 * i - 1; q++ {
		// 	fmt.Printf(" ")
		// }
		for t := 1; t <= f / 4 - 1; t++ {
			fmt.Printf("*")
    
		}
		fmt.Println()
	}
}