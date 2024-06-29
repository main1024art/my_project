package main
import "fmt"
type Person struct {
	Name string
}
func (p Person) speak() {
	var aaa [3][3]int = [3][3]int{{1,2,3},{4,5,6},{7,8,9}}
	var bbb [3][3]int
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			bbb[i][j] = aaa[j][i]
		}
		fmt.Println()
	}
	fmt.Println(bbb)
}
func main() {
	
	var p Person
	p.speak()
}