package main
import (
	"fmt"
	"sort"
)
func main() {
	map1 := make(map[int]int,10)
	map1[10] = 12
	map1[6] = 43
	map1[53] = 65
	map1[34] = 67
	map1[76] = 87

	fmt.Println(map1)
	var k []int
	for i, _ := range map1 {
		k = append(k, i)
	}
	sort.Ints(k)
	fmt.Println(k)
	for _, k1 := range k {
	fmt.Printf("map1[%v]=%v\n",k1,map1[k1])
	}
}