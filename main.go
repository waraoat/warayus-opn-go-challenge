package main

import (
	"fmt"
	"os"

	"tamboon/donation"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run main.go <filename>")
		return
	}

	filename := os.Args[1]

	fmt.Printf("performing donations...\n\n")
	
	donation.Process(filename)
}

