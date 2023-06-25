package main

import "fmt"

type interfaceA interface{}

type interfaceB interface{}

type A struct{}

func main() {
	var a interfaceA
	fmt.Printf("a의 주소: %p", &a)
	b, error := a.(*interfaceB)
	fmt.Printf("a의 주소: %p", &a)
	fmt.Printf("b: %p, error: %v", &b, &error)
}
