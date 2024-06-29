package main
import (
	"fmt"
	"math/rand"
	"time"
)

type Person struct {
	Name string
	Age int
	Address string
}
func shili(p Person, structChan chan Person) {
	for i := 0; i < 10; i++ {
		rand.Seed(time.Now().UnixNano())
	    randomInt := rand.Intn(100)
        p.Age = randomInt
		structChan<- p
		fmt.Println(randomInt)
	}
	close(structChan)

}
func bianli(i bool, structChan chan Person) {
    for {
		v, ok := <-structChan
		if !ok {
			break
		}
		fmt.Println(v)
	}
    i = true
}
func main() {
	var i bool
	var p Person
	structChan := make(chan Person, 10)
	go shili(p, structChan)
	go bianli(i, structChan)
    for {
		if i {
			break
		}
	}
	

}
