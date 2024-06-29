package main
import (
	"fmt"
	"time"
	"math/rand"
)

func main() {
	var char [5]int
	rand.Seed(time.Now().UnixNano())
    for i := 0; i < len(char); i++ {
		char[i] = rand.Intn(100)
	}
	fmt.Println(char)
	t := 0
	for i := 0; i < len(char) / 2; i++ {
		t = char[len(char) - 1 - i]
		char[len(char) - 1 - i] = char[i]
		char[i] = t		
	}
	fmt.Println(char)

}