package main

import (
	"fmt"
	"image/color"
	"math"
)

// Point xxx
type Point struct{ X, Y float64 }

func (p Point) distance(q Point) float64 {
	xdelta := p.X - q.X
	ydelta := p.Y - q.Y
	return math.Sqrt(xdelta*xdelta + ydelta*ydelta)
}

// ColoredPoint xxx
type ColoredPoint struct {
	*Point
	color.RGBA
}

func main() {
	var cp ColoredPoint
	cp.X = 1
	fmt.Println(cp.distance(Point{1, 5}))
	fmt.Println(cp)
	cp.Point.X = 2
	fmt.Println(cp)
}
