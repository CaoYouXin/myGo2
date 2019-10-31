package main

import "fmt"

type point struct {
	x, y int
}

type circle struct {
	center point
	radius int
}

type wheel struct {
	shape  circle
	spokes int
}

func main() {
	w := wheel{
		shape: circle{
			center: point{
				x: 6,
				y: 6,
			},
			radius: 6,
		},
		spokes: 36,
	}
	fmt.Printf("%#v\n", w)

	w.shape.center.x = 72
	fmt.Println(w)
}
