package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("GIF file not provided")
		return
	}

	filePath := os.Args[1]
	if !IsGIF(filePath) {
		fmt.Println("Provided file is not a GIF")
		return
	}

	fmt.Println(filePath)
}
