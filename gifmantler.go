package main

import (
	"fmt"
	"image"
	"image/gif"
	"image/png"
	"os"
	"path"
	"strconv"
	"sync"
)

//ProcessGIF will unpack the frames to individual images
func ProcessGIF(filePath string) {
	file, err := os.Open(filePath)
	defer file.Close()
	if err != nil {
		fmt.Println("Error reading file", filePath)
		return
	}
	gifRef, err := gif.DecodeAll(file)
	if err != nil {
		fmt.Println("Could not decode GIF")
		return
	}
	var wg sync.WaitGroup
	for id, img := range gifRef.Image {
		outPath := path.Join(GetOutputPath(filePath), strconv.Itoa(id+1)) + ".png"
		go writeImage(outPath, img, &wg)
	}
	wg.Wait()
}

func writeImage(imgPath string, img image.Image, wg *sync.WaitGroup) {
	wg.Add(1)
	defer wg.Done()

	fs, err := os.OpenFile(imgPath, os.O_CREATE, 0766)
	if err != nil {
		fmt.Println(err)
	}
	defer fs.Close()
	png.Encode(fs, img)
}
