package main

import "fmt"

var z string

func main() {
	x := 42
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	fmt.Println(z)
	fmt.Printf("%T\n", z)
}
