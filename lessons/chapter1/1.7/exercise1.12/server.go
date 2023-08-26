package main

import (
	"image"
	"image/color"
	"image/gif"
	"log"
	"math"
	"math/rand"
	"net/http"
	"strconv"

	//"os"
	"time"
)

var palette = color.Palette{
	color.RGBA{R: 00, G: 00, B: 00, A: 50},         // background: black
	color.RGBA{R: 0x00, G: 0xff, B: 0x00, A: 0xff}, // foreground: green
}

const (
	whiteIndex = 0
	blackIndex = 1
)

func main() {
	http.HandleFunc("/", lissajous)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

func lissajous(out http.ResponseWriter, r *http.Request) {
	var cycles float64 = 5
	cs := r.Header.Get("cycles")
	if cs != "" {
		cc, _ := strconv.Atoi(cs) // convert string to int
		if cc > 0 {
			cycles = float64(cc)
		}
	}

	const (
		res     = 0.001
		size    = 100
		nframes = 64
		delay   = 8
	)

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
			img.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), blackIndex)
		}
		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, img)
	}
	gif.EncodeAll(out, &anim)
}
