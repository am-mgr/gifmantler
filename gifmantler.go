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

//Parameters for processing GIF
type Parameters struct {
	IsPNG       bool
	IsJPEG      bool
	JpegQuality int
}

type ImageWriter = func(writer io.Writer, img image.Image)

//ProcessGIF will unpack the frames to individual images
func ProcessGIF(filePath string, params *Parameters) {
	gifFile, err := os.Open(filePath)
	if err != nil {
		fmt.Println("Error reading file", filePath)
		return
	}

	defer gifFile.Close()
	gifRef, err := gif.DecodeAll(gifFile)
	if err != nil {
		fmt.Println("Could not decode GIF")
		return
	}
	var wg sync.WaitGroup
	for id, img := range gifRef.Image {
		outPath := path.Join(GetOutputPath(filePath), strconv.Itoa(id+1))
		wg.Add(1)
		if params.IsPNG {
			go writeImage(outPath+".png", img, pngWriter(), &wg)
		}
		if params.IsJPEG {
			go writeImage(outPath+".jpeg", img, jpegWriter(params.JpegQuality), &wg)
		}

	}
	wg.Wait()
}

func writeImage(imgPath string, img image.Image, imgWriter ImageWriter, wg *sync.WaitGroup) {
	defer wg.Done()

	fs, err := os.OpenFile(imgPath, os.O_CREATE, 0766)
	if err != nil {
		fmt.Println(err)
	}
	defer fs.Close()
	imgWriter(fs, img)
}

func pngWriter() ImageWriter {
	return func(writer io.Writer, img image.Image) {
		png.Encode(writer, img)
	}
}

func jpegWriter(quality int) ImageWriter {
	return func(writer io.Writer, img image.Image) {
		jpeg.Encode(writer, img, &jpeg.Options{Quality: quality})
	}
}
