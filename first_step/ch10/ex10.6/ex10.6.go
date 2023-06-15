package main

import "fmt"

func getMyAge() int {
	return 22
}

func main() {
	variable := 3
	switch getMyAge(); variable {
	case 10:
		fmt.Println("Teenage")
	case 33:
		fmt.Println("Pair 33")
	default:
		fmt.Println("My age is", variable)
	}
}
