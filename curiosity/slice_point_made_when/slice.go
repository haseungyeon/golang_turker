package main

import "fmt"

func main() {
	slice1 := []int{1, 2, 3, 4, 5}
	slice2 := make([]int, len(slice1))

	fmt.Printf("slice1의 주소 : %p", slice1)
	fmt.Printf("slice2의 주소 : %p", slice2)
}