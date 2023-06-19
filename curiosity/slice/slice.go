package main

import "fmt"

func main() {
	var array [5]int = [5]int{1, 2, 3, 4, 5}
	slice1 := array[1:3:5]
	slice2 := append(slice1, 4, 5)

	fmt.Println(slice1)
	fmt.Println(slice2)
}
