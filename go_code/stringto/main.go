package main
import (
	"fmt"
	"strconv"
)
func main() {
	var str string = "ture"
	var b bool
	b , _ = strconv.ParseBool(str)
	fmt.Printf("b type %T b=%v",b,b)
}