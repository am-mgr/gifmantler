package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {

	version := flag.Bool("version", false, "prints tool version")
	flag.Parse()
	if *version {
		fmt.Println(Version)
		os.Exit(0)
	}

	if len(os.Args) < 2 {
		fmt.Println("GIF file not provided")
		os.Exit(1)
	}

	filePath := os.Args[1]
	if !IsGIF(filePath) {
		fmt.Println("Provided file is not a GIF")
		os.Exit(1)
	}

	ProcessGIF(filePath)

}
