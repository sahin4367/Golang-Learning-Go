package main

import (
	"fmt"
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

type MyImage struct{}

func (i MyImage) ColorModel() color.Model {
	return color.RGBAModel
}

func (i MyImage) Bounds() image.Rectangle {
	return image.Rect(0, 0, 256, 256)
}

func (i MyImage) At(x, y int) color.Color {
	// Bu düstur rəng keçidi (gradient) yaradır
	return color.RGBA{uint8(x), uint8(y), 255, 255}
}

type Image struct {
	W, H int
}

func (img Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (img Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, img.W, img.H)
}

func (img Image) At(x, y int) color.Color {
	// (x + y) / 2 düsturu fərqli bir vizual naxış yaradır
	v := uint8((x + y) / 2)
	return color.RGBA{v, v, 255, 255}
}

func main4() {
	m1 := MyImage{}
	fmt.Println("MyImage obyekti yaradıldı:", m1.Bounds())

	m2 := Image{256, 256}
	fmt.Println("Şəkil göstərilir...")
	pic.ShowImage(m2)
}
