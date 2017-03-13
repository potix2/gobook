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

const (
	bgColorIndex   = 0
	lineColorIndex = 1
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	lissajous(os.Stdout)
}

func makePalette(numColor int) []color.Color {
	palette := make([]color.Color, numColor+1)
	palette[0] = color.Black
	for i := 1; i <= numColor/2; i++ {
		green := uint8((0xFF / numColor) * i / 2)
		palette[i] = color.RGBA{0x0, green, 0x0, 0xFF}
		palette[numColor-i+1] = color.RGBA{0x0, green, 0x0, 0xFF}
	}
	return palette
}

func lissajous(out io.Writer) {
	const (
		cycles  = 5
		res     = 0.001
		size    = 64
		nframes = 64
		delay   = 8
	)
	freq := rand.Float64() * 3.0
	anim := gif.GIF{LoopCount: nframes}
	phase := 0.0
	numColor := nframes
	myPalette := makePalette(numColor)
	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		img := image.NewPaletted(rect, myPalette)
		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5),
				uint8(i%numColor)+1)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
