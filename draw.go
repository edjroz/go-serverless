package main 

import (
	"image"
	"image/draw"
)

func drawImage(base, watermark image.Image) *image.RGBA {
	bounds := base.Bounds()
	signed := image.NewRGBA(bounds)

	draw.Draw(signed, bounds, base, image.ZP, draw.Src)
	draw.Draw(signed, watermark.Bounds(), watermark, image.ZP, draw.Over)

	return signed
}
