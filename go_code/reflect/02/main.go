package main
import (
	"fmt"
	"reflect"
)
type Cal struct {
	Num1 int
	Num2 int
}
func (c Cal) Getsub(a int, b int) int {
	return a + b

}
func (c Cal) Bianli() {
		fmt.Println(c)
}
func aa(a interface{}) {
	val := reflect.ValueOf(a)
	// typ := reflect.TypeOf(a)
	n := val.NumMethod()
	fmt.Println(n)
	val.Method(0).Call(nil)
	var ppp []reflect.Value 
	ppp = append(ppp, reflect.ValueOf(8))
	ppp = append(ppp, reflect.ValueOf(3))
	res := val.Method(1).Call(ppp)
	fmt.Println(res[0].Int())

}
func main() {
	c := Cal{
		Num1 : 8,
		Num2 : 3,
	}
	aa(c)
}