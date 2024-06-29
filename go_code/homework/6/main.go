package main
import "fmt"
type visitor struct {
	name string 
	age int
}
func (v *visitor) say() int {
	if v.age > 18{
        return 20
	} else {
		return 0
	}
}
	

func main() {
	var d = visitor{
		name : "tom",
		age : 12,
	}
	fmt.Println(d.say())
}