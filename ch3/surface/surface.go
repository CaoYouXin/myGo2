package main

import (
	"fmt"
	"log"
	"math"
	"net/http"
)

func main() {
	http.HandleFunc("/", surface)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

const (
	width, height = 2000, 1200
	cells         = 100
	xyrange       = 30.0
	xyscale       = width / 2 / xyrange
	zscale        = height * 0.4
	angle         = math.Pi / 6
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

// var minZ, maxZ float64

func surface(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/svg+xml")

	fmt.Fprintf(w, "<svg xmlns='http://www.w3.org/2000/svg' style='stroke:grey; fill:white; stroke-width:0.7' width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			// if i > cells/2 && j > cells/2 {
			// 	continue
			// }
			ax, ay, az := cornor(i+1, j)
			bx, by, bz := cornor(i, j)
			cx, cy, cz := cornor(i, j+1)
			dx, dy, dz := cornor(i+1, j+1)

			avgZ := int(((az+bz+cz+dz)/4 + 0.21722891503668823) / (0.9850673555377986 + 0.21722891503668823) * 100)

			fmt.Fprintf(w, "<polygon points='%g,%g %g,%g %g,%g %g,%g' style='stroke: rgb(%d%%, 0%%, %d%%); fill: rgb(%d%%, 0%%, %d%%)' />", ax, ay, bx, by, cx, cy, dx, dy, avgZ, 100-avgZ, avgZ, 100-avgZ)
		}
	}
	fmt.Fprintln(w, "</svg>")
	// fmt.Println(minZ, maxZ)
}

func cornor(i, j int) (float64, float64, float64) {
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	z := f(x, y)
	// if z < minZ {
	// 	minZ = z
	// }
	// if z > maxZ {
	// 	maxZ = z
	// }

	sx := width/2 + (x-y)*cos30*xyrange
	sy := height/2 + (x+y)*sin30*xyrange - z*zscale

	return sx, sy, z
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y)
	return math.Sin(r) / r
}
