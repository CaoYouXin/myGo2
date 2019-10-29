package lissajous

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

var (
	palette = []color.Color{color.RGBA{0x00, 0xFF, 0x00, 0xFF}, color.White, color.Black}
)

const (
	// cycles  = 5     // number of complete x oscillator revolutions
	res     = 0.001 // angular resolution
	size    = 100   // image canvas covers [-size..+size]
	nframes = 64    // number of animation frames
	delay   = 8     // delay between frames in 10ms units
)

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	Lissajous(os.Stdout, 5)
}

// Lissajous wow
func Lissajous(out io.Writer, cycles float64) {
	freq := rand.Float64() * 3.0
	phase := 0.0

	anim := gif.GIF{LoopCount: nframes}

	for i := 0; i < nframes; i++ {
		rect := image.Rect(0, 0, 2*size+1, 2*size+1)
		imag := image.NewPaletted(rect, palette)

		for t := 0.0; t < cycles*2*math.Pi; t += res {
			x := math.Sin(t)
			y := math.Sin(t*freq + phase)
			imag.SetColorIndex(size+int(x*size+0.5), size+int(y*size+0.5), uint8(rand.Intn(2)+1))
		}

		phase += 0.1
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, imag)
	}

	gif.EncodeAll(out, &anim)
}
