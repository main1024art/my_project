package main 
import (
	"fmt"
	"strconv"
)
func main() {
	var i int =  99
	var a float64 = 12.3345
	var str string
	str = strconv.FormatInt(int64(i),10)
	fmt.Printf("str type %T str=%q",str,str)
	str = strconv.FormatFloat(a,'f',10,64)
	fmt.Printf("str type %T str=%q",str,str)
}