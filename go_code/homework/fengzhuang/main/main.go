package main
import (
      "fmt"
	  "go_code/homework/fengzhuang/model"
)

func main() {
	var p model.Account
    p.Set("12345444",35,"345656")
	fmt.Println(p.Get())
}