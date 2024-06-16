package main

import "fmt"

func main() {
	counter := 4
	x := decrement(counter)
	fmt.Println(x())
	fmt.Println(x())
	fmt.Println(x())
	fmt.Println(x())
}
