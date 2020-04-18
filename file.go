package main

import (
	"path"
	"strings"
)

//IsGIF checks if provided file is a GIF
func IsGIF(filePath string) bool {
	return strings.Compare(path.Ext(filePath), ".gif") == 0
}
