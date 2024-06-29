package main
import (
	"fmt"
	"time"
	"math/rand"
)
func erfenfa(arr *[10]int,zuo int,you int,find int) {
	if zuo > you{
		fmt.Println("meiyou")
		return
	}
	z := (zuo + you) / 2
	if find < (*arr)[z] {
		erfenfa(arr, zuo, z - 1, find)
	} else if  find > (*arr)[z] {
		erfenfa(arr, z + 1, you,find)
	} else {
		fmt.Println("找到了 下标为",z)
	}
}
func main() {
	var char [10]int
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len(char); i++ {
		char[i] = rand.Intn(100) + 1
	}
	fmt.Println(char)
    
	// 冒泡排序
	t := 0
	for i := 0; i < len(char) - 1; i++  { 
		for j := 0; j < len(char) - 1 - i; j++ {
			if char[j] > char [j + 1] {
				t = char[j]
				char[j] = char [j + 1]
				char [j + 1] = t
		    }
	    }
	}
	fmt.Println(char)
	erfenfa(&char,0,len(char),90)
}	
    