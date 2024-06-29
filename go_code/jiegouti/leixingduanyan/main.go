package main
import (
	"fmt"
)
type Usb interface {
	star()
	stop()
}
type Phone struct {
	name string
}
type Carmer struct {
	name string
}
type Computer struct {
}
func (p Phone) star() {
	fmt.Println("运行")
}
func (p Phone) stop() {
	fmt.Println("停止")
}
func (c Carmer) star() {
	fmt.Println("运行")
}
func (c Carmer) stop() {
	fmt.Println("停止")
}
func (p Phone) call() {
	fmt.Println("接听电话")
}
func (c Computer) working(usb Usb) {
	usb.star()
	if phone, ok := usb.(Phone); ok {
		phone.call()
	}
	usb.stop()
}
func main() {
	var a [3]Usb
	a[0] = Phone{"oppo"}
	a[1] = Phone{"vivo"}
	a[2] = Carmer{"sony"}
	var computer Computer
	for _, v := range a {
		computer.working(v)
		fmt.Println()
	}


}