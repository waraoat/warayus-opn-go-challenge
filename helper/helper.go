package helper

import "fmt"

func Contains(slice []string, val string) bool {
	for _, item := range slice {
		if item == val {
			return true
		}
	}
	return false
}

func ShowProgressBar(current, total int) {
	const width = 50
	percent := float64(current) / float64(total)
	numChars := int(percent * width)

	fmt.Printf("\r[")
	for i := 0; i < numChars; i++ {
		fmt.Print("=")
	}
	for i := numChars; i < width; i++ {
		fmt.Print(" ")
	}
	fmt.Printf("] (%d/%d)", current, total)
}