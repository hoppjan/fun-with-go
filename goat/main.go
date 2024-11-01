package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	path := "main.go"

	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}

	content, err := os.ReadFile(file.Name())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s", content)
}
