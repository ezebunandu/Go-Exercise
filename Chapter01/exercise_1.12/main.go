package main

import (
	"fmt"
	"time"
)

func main() {
	var count int
	fmt.Printf("Discount: %#v \n", count)
	var discount float64
	fmt.Printf("Discount: %#v\n", discount)
	var mapper map[string]int
	if mapper != nil {
		fmt.Println("The mapper is not nil")
	}
	fmt.Printf("Mapper: %#v \n", mapper)

	var ptr *int
	if ptr == nil {
		fmt.Println("It is nil")
	}
	fmt.Printf("ptr: %#v \n", ptr)
	var debug bool
	fmt.Printf("Debug: %#v \n", debug)
	var message string
	fmt.Printf("Message: %#v \n", message)
	var emails []string
	fmt.Printf("Emails: %#v \n", emails)
	var StartTime time.Time
	fmt.Printf("Start: %#v \n", StartTime)
}
