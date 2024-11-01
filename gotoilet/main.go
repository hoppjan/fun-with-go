package main

import (
	"fmt"
	"log"
	"os"
)

// gotoilet counts characters in a file, like wc
func main() {
	if len(os.Args) <= 2 {
		log.Fatalln(
			"Usage: gotoilet [OPTION] path/to/file..txt\n",
			"Options:\n",
			"-c\tcount characters, including line breaks",
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
	case "-c":
		length := len(content)
		fmt.Println(length)
	default:
		break
	}
}
