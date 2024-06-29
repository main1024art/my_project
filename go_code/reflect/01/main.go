package main
import (
	"fmt"
	"reflect"
)
func aa(n interface{}) {
	rtype := reflect.TypeOf(n)
	fmt.Println("rtype", rtype)
	rValue := reflect.ValueOf(n)
	rKind := rValue.Kind()
	fmt.Println("rKind", rKind)
	cc := rValue.Interface()
	zz := cc.(float64)
	fmt.Println(zz)

}
func main() {
	v := 1.2
	aa(v)

}