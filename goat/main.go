package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) <= 1 {
		log.Fatal("Usage: goat path/to/file.go")
	}

	path := os.Args[1]

	if _, err := os.Stat(path); os.IsNotExist(err) {
		log.Fatalf("File %s does not exist", path)
	}

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
