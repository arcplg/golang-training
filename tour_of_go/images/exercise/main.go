package main

import (
	"fmt"
	"image"
	"image/color"
)

type Image struct {
	width, height int
}

func (img Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (img Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, img.width, img.height)
}

func (img Image) At(x, y int) color.Color {
	v := uint8((x + y) % 256)
	return color.RGBA{v, v, 255, 255}
}

func main() {
	m := Image{width: 100, height: 100}
	fmt.Println(m.Bounds())
	fmt.Println(m.At(0, 0))
}
