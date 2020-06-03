package main

import "fmt"

func main() {
	x := 45
	fmt.Println(x)
	x = 56 + 67
	fmt.Println(x)
	type niche int
	var z niche
	x = int(z)
	fmt.Println(x)
	s := fmt.Sprintf("%T\n%T", x, z)
	fmt.Println(s)
}
