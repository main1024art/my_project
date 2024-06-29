package main
import (
	"fmt"
)

func sum(numChan chan int, resChan chan int, exitChat chan bool) {
	fmt.Println("aaaaa")	
	for v := range numChan{
		sum := 0
		for i := 1; i <= v; i++{
			sum += i 
		}
		resChan<- sum	
	}
	exitChat<- true
}

func bianli(resChan chan int) {

	for v := range resChan{
		fmt.Println("v=", v)
	}
	
	
}
func main() {
	var resChan chan int = make(chan int, 100)
	numChan := make(chan int, 100)
	exitChat := make(chan bool, 1)
	for i := 1; i <= 100; i++ {
		numChan <- i
	}
	close(numChan)
	bianli(numChan)
	go sum(numChan, resChan, exitChat)
	go sum(numChan, resChan, exitChat)
	go sum(numChan, resChan, exitChat)
	go sum(numChan, resChan, exitChat)
	fmt.Println("qqqq")
	
	for {
		v, _ := <-exitChat
		if v {
		    break
	    }
	}
	fmt.Println("qqqq")
	bianli(resChan)
	

}