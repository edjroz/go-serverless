package serverless
import (
  "image"
  "github.com/fogleman/gg"
)

func signature(name string) image.Image {
   const S = 128

    m := image.NewRGBA(image.Rect(0, 0, 128, 128))
    dc := gg.NewContext(S, S)
    dc.SetRGB(1, 1, 1)
    dc.Clear()
    dc.SetRGB(0, 0, 0)

    dc.DrawImage(m, 0, 0)
    dc.DrawStringAnchored(name, S/2, S/2, 0.5, 0.5)
    dc.Clip()
    return dc.Image()
}
