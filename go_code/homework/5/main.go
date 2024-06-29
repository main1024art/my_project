package main
import "fmt"
type Dog struct {
	name string 
	age int
	weight float64
}
func (d *Dog) say() string {
	a := fmt.Sprintf("name=[%v],age=[%v],weight=[%v]",d.name, d.age, d.weight)
	return a
}
func main() {
	var d = Dog{
		name : "tom",
		age : 12,
		weight : 23.0,
	}
	fmt.Println(d.say())
}