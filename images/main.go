package main

import (
	"image"
	"image/color"
	"golang.org/x/tour/pic"
)

type Image struct{
	width, height int
}

func (img Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (img Image) Bounds() image.Rectangle {
	return image.Rectangle {
		image.Point { 0, 0 },
		image.Point{
			img.width,
			img.height,
		},
	}
}

func (img Image) At(x, y int) color.Color {
	v := x ^ y
	return color.RGBA{uint8(v),uint8(v), 255, 255}
}

func main() {
	m := Image{256, 256}
	pic.ShowImage(m)
}
