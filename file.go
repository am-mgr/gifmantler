package main

import (
	"os"
	"path"
	"path/filepath"
	"strings"
)

//IsGIF checks if provided file is a GIF
func IsGIF(filePath string) bool {
	return strings.Compare(path.Ext(filePath), ".gif") == 0
}

//GetOutputPath gets location to drop unpacked images
func GetOutputPath(filePath string) string {
	absPath := filepath.Dir(filePath)
	outPath := path.Join(absPath, "gif_out")
	if _, err := os.Stat(outPath); os.IsNotExist(err) {
		os.Mkdir(outPath, os.ModeDir)
	}
	return outPath
}
