package main
import (
	"fmt"
)
func fbn(n int) ([]uint64) {
	aaa := make([]uint64,n)
	for i := 0; i <= n - 1; i++ {
		if (i == 0 || i == 1) {
			aaa[i] = 1
		} else {
			aaa[i] = aaa[i - 1] + aaa[i - 2]
		}
	} 
	return aaa
}
func main() {
	bbb := fbn(10)
	fmt.Println(bbb)
}
