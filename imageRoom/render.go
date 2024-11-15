package main

import (
	"image"
	"image/color"
	"image/png"
	"os"
)

func main() {
	width := 1000
	height := 1000
	img := image.NewRGBA(image.Rect(0, 0, width, height))

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			img.Set(x, y, color.RGBA{uint8(x), uint8(y), 255, 255})
		}
	}

	f, _ := os.Create("image.png")
	defer f.Close()
	png.Encode(f, img)
}
