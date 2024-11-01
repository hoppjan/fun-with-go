package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	if len(os.Args) <= 1 {
		log.Fatal("Usage: goat path/to/file.go path/to/another.txt")
	}

	for i := 1; i < len(os.Args); i++ {
		path := os.Args[i]

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
}
