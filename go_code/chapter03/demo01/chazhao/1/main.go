package main
import (
	"fmt"
	"time"
	"math/rand"
)
func main() {
	var char [10]int
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len(char); i++ {
		char[i] = rand.Intn(100) + 1
	}
	fmt.Println(char)
	
	//倒序
	for i := 0; i < len(char) / 2; i++ {
		t := 0
		t = char[i]
		char[i] = char[len(char) - i -1]
		char[len(char) - i -1] = t
	}
	fmt.Println(char)
	
	//平均值
	sum := 0
	for i := 0; i < len(char); i++ {
		sum += char[i]
	}
	fmt.Println(sum / len(char))

	//最大值最小值
	max := 0
	a := 0
	for i := 0; i < len(char); i++ {
		if max < char[i]{
			max = char[i]
			a = i
		}
	}
	fmt.Println(max,a)
	var min int
	if char[0] < char[1]{
		min = char[0]
	} 
	for i := 0; i < len(char) - 1; i++ {
		if min >= char[i]{
			min = char[i]
			a = i
		} 
	}
	fmt.Println(min,a)

	//查找55
	b := 55
	for i := 0; i < len(char); i++ {
		if char[i] == b {
			fmt.Println("有55")
		}
    }
	fmt.Println("没有55")
}
