package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"unicode/utf8"
)

// gotoilet counts stuff in a file, like wc
func main() {
	if len(os.Args) <= 2 {
		log.Fatalln(
			"Usage: gotoilet [OPTION] path/to/file..txt\n",
			"Options:\n",
			"-w\tcount words\n",
			"-l\tcount lines\n",
			"-r\tcount runes\n",
			"-b\tcount bytes",
		)
	}

	option, path := os.Args[1], os.Args[2]

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

	switch option {
	case "-w":
		words := len(strings.Fields(string(content)))
		fmt.Println(words)
	case "-l":
		lines := strings.Count(string(content), "\n")
		fmt.Println(lines)
	case "-r":
		runes := utf8.RuneCount(content)
		fmt.Println(runes)
	case "-b":
		length := len(content)
		fmt.Println(length)
	default:
		log.Fatalf("unknown option: %s\n", option)
	}
}
