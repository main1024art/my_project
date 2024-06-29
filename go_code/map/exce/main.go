package main
import (
	"fmt"
)
func modifyUser(user map[string]map[string]string, name string) {
	if user[name] != nil {
		user[name]["pwd"] = "888888"
	} else {
        user[name] = make(map[string]string, 2)
		user[name]["pwd"] = "888888"
		user[name]["nickname"] = "nick" + name
	}
	

}
func main() {
    user := make(map[string]map[string]string)
	modifyUser(user,"tom")
	fmt.Println(user )
}