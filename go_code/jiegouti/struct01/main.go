package main
import "fmt"
type Person struct {
	Name string
}
func (p Person) speak() {
	fmt.Println("haoren")
}
func (p Person) jisuan() {
	res := 0
	for i := 1; i <= 100; i++ {
		res += i
	}
	fmt.Println(res)

}

func main() {
	var p Person
	p.Name = "qqq"
	fmt.Println(p)
	p.speak()
    p.jisuan()
}