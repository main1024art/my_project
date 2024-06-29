package main
import (
	"fmt"
)
func main() {
	i := 1
	
	var a float64
	for ; i <= 3; i++ {
		var s float64
		sum := 0.0
		for t := 1; t <= 5; t++ {
			fmt.Printf("输入%v成绩",t)
			fmt.Scanln(&s)
			sum += s
		}
		a += sum
		fmt.Println("平均值a=\n",sum / 5)
    }
	fmt.Println("所以平均值",a / float64(i * 5))
}