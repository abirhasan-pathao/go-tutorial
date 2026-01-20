package main

import (
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

type Image struct{}

func (m Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (m Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, 256, 256)
}

func (m Image) At(x, y int) color.Color {
	c := uint8((x + y) / 2)
	return color.RGBA{c, c, 255, 255}
}

func main() {
	m := Image{}
	pic.ShowImage(m)
}

// copy the code and replace the code here https://go.dev/tour/methods/24 to see the actual image
