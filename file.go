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

//OutputLocation gets location to drop unpacked images
func OutputLocation(filePath string) string {
	out := path.Join(filePath, "gif_out")
	if _, err := os.Stat(out); os.IsNotExist(err) {
		os.Mkdir(out, os.ModeDir)
	}
	return out
}

//GetRootDir gets the root directory
func GetRootDir(filePath string) string {
	return filepath.Dir(filePath)
}
