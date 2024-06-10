package main

import (
	"fmt"
	"os"
)

func getPassedArgs(minArgs int) []string {
	if len(os.Args) < minArgs {
		fmt.Printf("At least %v arguments are needed\n", minArgs)
		os.Exit(1)
	}
	var args []string
	args = append(args, os.Args[1:]...)
	return args
}

func findLngest(args []string) string {
	var longest string
	for _, arg := range args {
		if len(arg) > len(longest) {
			longest = arg
		}
	}
	return longest
}

func main() {
	if longest := findLngest(getPassedArgs(3)); len(longest) > 0 {
		fmt.Println("The longest word passed was: ", longest)
	} else {
		fmt.Println("There was an error")
		os.Exit(1)
	}
}
