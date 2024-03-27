package main

import (
	"fmt"
	"os"

	"tamboon/donation"
)

var PUBLIC_KEY_OPN = os.Getenv("PUBLIC_KEY_OPN")
var SECRET_KEY_OPN = os.Getenv("SECRET_KEY_OPN")

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run main.go <filename>")
		return
	}

	filename := os.Args[1]

	fmt.Println("performing donations...")
	
	donation.Process(filename)
}

type Config struct {
	OpnPublicKey string
	OpnSecretKey string
}
