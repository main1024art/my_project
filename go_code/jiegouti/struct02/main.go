package main
import "fmt"
type Circle struct {
	Nradius float64
}
func (c Circle) test() float64{
	return (3.14 * c.Nradius * c.Nradius)
}
func main() {
	var a Circle
	var res float64
	a.Nradius = 3.0
	res = a.test()
	fmt.Println(res)
}