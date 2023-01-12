package main

import "image"

func OverlayLogo(dst *image.Paletted, logo image.Image) *image.NRGBA {
	res := image.NewNRGBA(dst.Rect)

	for x := 0; x < dst.Bounds().Max.X; x++ {
		for y := 0; y < dst.Bounds().Max.Y; y++ {
			res.Set(x, y, dst.At(x, y))
		}
	}

	offsetX := dst.Bounds().Max.X/2 - logo.Bounds().Max.X/2
	offsetY := dst.Bounds().Max.Y/2 - logo.Bounds().Max.Y/2

	for x := 0; x < logo.Bounds().Max.X; x++ {
		for y := 0; y < logo.Bounds().Max.Y; y++ {
			if _, _, _, alpha := logo.At(x, y).RGBA(); alpha > uint32(200) {
				res.Set(x+offsetX, y+offsetY, logo.At(x, y))
			}
		}
	}

	return res
}
