package main

import (
	"fmt"
	"image/gif"
	"image/png"
	"os"
	"path"
	"strconv"
)

//ProcessGIF will unpack the frames to individual images
func ProcessGIF(filePath string) {
	file, err := os.Open(filePath)
	defer file.Close()
	if err != nil {
		fmt.Println("Error reading file", filePath)
	}
	gifRef, err := gif.DecodeAll(file)
	if err != nil {
		fmt.Println("Could not decode GIF")
	}
	for id, img := range gifRef.Image {
		outPath := path.Join(GetOutputPath(filePath), strconv.Itoa(id)) + ".png"
		fs, err := os.OpenFile(outPath, os.O_CREATE, 0766)
		if err != nil {
			fmt.Println(err)
		}
		defer fs.Close()
		png.Encode(fs, img)
	}
}
