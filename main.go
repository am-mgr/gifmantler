package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {

	version := flag.Bool("version", false, "prints tool version")
	isPng := flag.Bool("png", false, "output as png images")
	isJpeg := flag.Bool("jpeg", false, "output as jpeg images")
	jpegQuality := flag.Int("jpeg-quality", 100, "quality of jpeg images ranging 1 to 100")

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

	ProcessGIF(filePath, &Parameters{
		IsPNG:       *isPng,
		IsJPEG:      *isJpeg,
		JpegQuality: *jpegQuality,
	})

}
