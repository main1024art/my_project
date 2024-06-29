package main
import (
	"fmt"
)
func main() {
	var aaa [3][4]int 
	for i := 0; i < len(aaa); i++ {
		for j := 0; j < len(aaa[i]); j++ {
			fmt.Println("输入")
			fmt.Scanln(&aaa[i][j])
		}
	}

	for _, v := range aaa {
		for _, v2 := range v {
			fmt.Printf("%v\t",v2)
		}
        fmt.Println()
	}
	for i := 0; i < len(aaa); i++ {
		for j := 0; j < len(aaa[i]); j++ {
			if i == 0 || j == 0 || i == 2 || j == 3 {
				aaa[i][j] = 0
			}
		}

	}
	for _, v := range aaa {
		for _, v2 := range v {
			fmt.Printf("%v\t",v2)
		}
        fmt.Println()
	}
		
}