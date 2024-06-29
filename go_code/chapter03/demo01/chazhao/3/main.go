package main
import (
	"fmt"
	"time"
	"math/rand"
	"math"
)
var max int
var min int
func zui(arr [10]int) (int,int){
	a := 0
	for i := 0; i < len(arr); i++ {
		if max < arr[i]{
			max = arr[i]
			a = i
		}
	}
	fmt.Println(max,a)
	if arr[0] < arr[1]{
		min = arr[0]
	} 
	for i := 0; i < len(arr) - 1; i++ {
		if min >= arr[i]{
			min = arr[i]
			a = i
		} 
	}
	fmt.Println(min,a)
	return max,min
}
var max1 float64
var min1 float64
func zui1(arr [10]float64) (float64,float64){
	a := 0
	for i := 0; i < len(arr); i++ {
		if max1 < arr[i]{
			max1 = arr[i]
			a = i
		}
	}
	fmt.Println(max1,a)
	if arr[0] < arr[1]{
		min1 = arr[0]
	} 
	for i := 0; i < len(arr) - 1; i++ {
		if min1 >= arr[i]{
			min1 = arr[i]
			a = i
		} 
	}
	fmt.Println(min1,a)
	return max1,min1
}
func main() {
	var char [10]int
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < len(char); i++ {
		char[i] = rand.Intn(100) + 1
	}
	fmt.Println(char)

	//最大值最小值
	zui(char)
    //平均值
    sum := 0
	for i := 0; i < len(char); i++ {
		sum += char[i]
	}
    p := (sum - min - max) / (len(char) - 2)
	fmt.Println(p)
	//差值
	var aaa [10]float64
	for i := 0; i < len(char); i++ {
		aaa[i] = math.Abs(float64(p - char[i]))
	}
	fmt.Println(aaa)
	zui1(aaa)
	
}