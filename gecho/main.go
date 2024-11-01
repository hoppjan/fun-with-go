package main

import (
	"fmt"
	"os"
)

// gecho is a program that prints its arguments like the echo command
func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Print(os.Args[i], " ")
	}
}
