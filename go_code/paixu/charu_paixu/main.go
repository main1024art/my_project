package main

import "fmt"

func main() {
	arr := [6]int{98, 9, 55, 12, 34, 45}
	charu(&arr)
}
func charu(arr *[6]int) {
	for j := 1; j < len(arr); j++ {
		v := arr[j]
		i := j - 1
		for i >= 0 && arr[i] < v {
			arr[i+1] = arr[i]
			arr[i] = v
			i--
		}
	}
	fmt.Println(*arr)
}
