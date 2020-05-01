package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {

	version := flag.Bool("version", false, "Prints tool version")
	isPng := flag.Bool("png", false, "Output as png images")
	isJpeg := flag.Bool("jpeg", false, "Output as jpeg images")
	jpegQuality := flag.Int("jpeg-quality", 100, "Quality of jpeg images ranging 1 to 100. Higher number, better quality")
	sourceGif := flag.String("gif", "", "The source GIF")

	flag.Parse()
	if *version {
		fmt.Println(Version)
		os.Exit(0)
	}

	if len(*sourceGif) == 0 {
		fmt.Println("GIF file not provided")
		os.Exit(1)
	}

	filePath := *sourceGif
	if !IsGIF(filePath) {
		fmt.Println("Provided file is not a GIF")
		os.Exit(1)
	}

	if (*jpegQuality < 1 || *jpegQuality > 100) && *isJpeg {
		fmt.Println("Jpeg quality set as", *jpegQuality, ". Allowed values range 1-100")
		os.Exit(1)
	}

	ProcessGIF(filePath, &Parameters{
		IsPNG:       *isPng,
		IsJPEG:      *isJpeg,
		JpegQuality: *jpegQuality,
	})

}
