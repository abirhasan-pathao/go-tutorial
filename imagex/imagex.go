package main

import (
	"golang.org/x/tour/pic"
	"image"
	"image/color"
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
