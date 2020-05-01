package main

import (
	"fmt"
	"image"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io"
	"os"
	"path"
	"strconv"
	"sync"
)

//ProcessGIF will unpack the frames to individual images
func ProcessGIF(filePath string) {
	gifFile, err := os.Open(filePath)
	defer gifFile.Close()
	if err != nil {
		fmt.Println("Error reading file", filePath)
		return
	}
	gifRef, err := gif.DecodeAll(gifFile)
	if err != nil {
		fmt.Println("Could not decode GIF")
		return
	}
	var wg sync.WaitGroup
	for id, img := range gifRef.Image {
		outPath := path.Join(GetOutputPath(filePath), strconv.Itoa(id+1))
		wg.Add(2)
		go writeImage(outPath+".png", img, pngWriter(), &wg)
		go writeImage(outPath+".jpeg", img, jpegWriter(100), &wg)

	}
	wg.Wait()
}

func writeImage(imgPath string, img image.Image, imgWriter func(writer io.Writer, img image.Image), wg *sync.WaitGroup) {
	defer wg.Done()

	fs, err := os.OpenFile(imgPath, os.O_CREATE, 0766)
	if err != nil {
		fmt.Println(err)
	}
	defer fs.Close()
	imgWriter(fs, img)
}

func pngWriter() func(writer io.Writer, img image.Image) {
	return func(writer io.Writer, img image.Image) {
		png.Encode(writer, img)
	}
}

func jpegWriter(quality int) func(writer io.Writer, img image.Image) {
	return func(writer io.Writer, img image.Image) {
		jpeg.Encode(writer, img, &jpeg.Options{Quality: quality})
	}
}
