package main

import(
	"testing"
)

func TestIsGif(t *testing.T)  {
	var file = "foo.gif"
	if !IsGIF(file) {
		t.Error("The file is not a GIF")
	}
}