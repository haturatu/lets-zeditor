package main

import "fmt"

var g *int

func first() {
	a := 1
	g = &a
}

func second() {
	b := 2
	_ = b
}
func main() {
	first()
	second()
	fmt.Print(*g, "\n")
	fmt.Print("Hello, Zed!\n")
}
