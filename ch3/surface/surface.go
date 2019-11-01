package main

import (
	"fmt"
	"log"
	"math"
	"net/http"
	"strconv"
)

func main() {
	http.HandleFunc("/", surface)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

const (
	angle = math.Pi / 6
	cells = 100
	zmin  = -0.21722891503668823
	zmax  = 0.9850673555377986
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

type params struct {
	width, height, xyrange, xyscale, zscale float64
}

var paramsOne params

func parse(str string, errOut *http.ResponseWriter) (res float64) {
	res, err := strconv.ParseFloat(str, 64)
	if err != nil {
		fmt.Fprintf(*errOut, "surface: %v\n", err)
	}
	return
}

func surface(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		fmt.Fprintf(w, "surface: %v\n", err)
		return
	}

	for p, v := range r.Form {
		switch p {
		case "w":
			paramsOne.width = parse(v[0], &w)
		case "h":
			paramsOne.height = parse(v[0], &w)
		case "range":
			paramsOne.xyrange = parse(v[0], &w)
		}
		paramsOne.xyscale = paramsOne.width / 2 / paramsOne.xyrange
		paramsOne.zscale = paramsOne.height * 0.4
	}

	w.Header().Set("Content-Type", "image/svg+xml")

	fmt.Fprintf(
		w,
		`<svg xmlns='http://www.w3.org/2000/svg'
			style='stroke:grey; fill:white; stroke-width:0.7'
			width='%g' height='%g'>`,
		paramsOne.width,
		paramsOne.height,
	)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			// if i > cells/2 && j > cells/2 {
			// 	continue
			// }
			ax, ay, az := cornor(i+1, j)
			bx, by, bz := cornor(i, j)
			cx, cy, cz := cornor(i, j+1)
			dx, dy, dz := cornor(i+1, j+1)

			avgZ := int(((az+bz+cz+dz)/4 - zmin) / (zmax - zmin) * 100)

			fmt.Fprintf(w, "<polygon points='%g,%g %g,%g %g,%g %g,%g' style='stroke: rgb(%d%%, 0%%, %d%%); fill: rgb(%d%%, 0%%, %d%%)' />", ax, ay, bx, by, cx, cy, dx, dy, avgZ, 100-avgZ, avgZ, 100-avgZ)
		}
	}
	fmt.Fprintln(w, "</svg>")
}

func cornor(i, j int) (sx, sy, z float64) {
	x := paramsOne.xyrange * (float64(i)/cells - 0.5)
	y := paramsOne.xyrange * (float64(j)/cells - 0.5)

	z = f(x, y)

	sx = paramsOne.width/2 + (x-y)*cos30*paramsOne.xyrange
	sy = paramsOne.height/2 + (x+y)*sin30*paramsOne.xyrange - z*paramsOne.zscale

	return
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y)
	return math.Sin(r) / r
}
