// Изменяем программу lissajous так, чтобы она генерировала изображения
// разных цветов, добавляя в палитру p a l e t t e больше значений, а затем
// выводя их путем изменения третьего аргумента функции SetColorIndex
// некоторым нетривиальным способом.
package main

import (
	"image"
	"image/color"
	"image/gif"
	"io"
	"math"
	"math/rand"
	"os"
	"time"
)

var palette = []color.Color{
	color.Black,
	color.RGBA{0, 255, 0, 255},
	color.RGBA{0, 51, 255, 255},
	color.RGBA{205, 0, 0, 255},
	color.RGBA{255, 153, 0, 255},
	color.RGBA{153, 0, 255, 255},
}

func main() {
	lissajous(os.Stdout)
}

func lissajous(out io.Writer) {
	const (
		cycles  = 5
		res     = 0.001
		size    = 100
		nframes = 64
		delay   = 8
	)

	var colorIndex uint8 = 1
	var points int

	rand.Seed(time.Now().UTC().UnixNano())
	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, palette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			points++
			if points%500 == 0 {
				colorIndex++
				if colorIndex > uint8(len(palette)) {
					colorIndex = 1
				}
			}
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), colorIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
