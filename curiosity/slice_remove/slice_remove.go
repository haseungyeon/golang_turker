package main

// import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5, 6}
	slice = append(slice, 0)
	arr := []int{1, 2}
	slice = append(slice[:2], arr...)
}
