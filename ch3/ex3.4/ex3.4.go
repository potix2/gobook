package main

import (
	"fmt"
	"math"
	"net/http"
	"io"
	"log"
)

const (
	width, height = 600, 320
	cells         = 100
	xyrange       = 30.0
	xyscale       = width / 2 / xyrange
	zscale        = height * 0.4
	angle         = math.Pi / 6
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

func main() {
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "image/svg+xml")

	pointsList, minZ, maxZ := buildPoints()
	renderSVG(w, pointsList, minZ, maxZ)
}

func buildPoints() ([][]float64, float64, float64) {
	pointsList := [][]float64{}
	minZ := 0.0
	maxZ := 0.0
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay, az := corner(i+1, j)
			bx, by, bz := corner(i, j)
			cx, cy, cz := corner(i, j+1)
			dx, dy, dz := corner(i+1, j+1)
			if math.IsNaN(ax) || math.IsNaN(ay) ||
				math.IsNaN(bx) || math.IsNaN(by) ||
				math.IsNaN(cx) || math.IsNaN(cy) ||
				math.IsNaN(dx) || math.IsNaN(dy) {
				continue
			}
			avgZ := (az + bz + cz + dz) / 4.0

			points := make([]float64, 9)
			points[0] = ax
			points[1] = ay
			points[2] = bx
			points[3] = by
			points[4] = cx
			points[5] = cy
			points[6] = dx
			points[7] = dy
			points[8] = avgZ
			pointsList = append(pointsList, points)

			if avgZ < minZ {
				minZ = avgZ
			}
			if maxZ < avgZ {
				maxZ = avgZ
			}
		}
	}
	return pointsList, minZ, maxZ
}

func renderSVG(w io.Writer, pointsList [][]float64, minZ float64, maxZ float64) {
	fmt.Fprintf(w, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>\n", width, height)
	for i := range pointsList {
		p := pointsList[i]
		fmt.Fprintf(w, "<polygon points='%g,%g %g,%g %g,%g %g,%g' style='stroke:%s' />\n",
			p[0], p[1], p[2], p[3], p[4], p[5], p[6], p[7], z2c(p[8], minZ, maxZ))
	}
	fmt.Fprintf(w, "</svg>\n")
}

func z2c(z, min, max float64) string {
	normal := (z - min) / (max - min)
	r := uint8(normal * 255)
	b := uint8((1.0 - normal) * 255)
	return fmt.Sprintf("#%02x00%02x", r, b)
}

func corner(i, j int) (float64, float64, float64) {
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	z := f(x, y)
	if math.IsNaN(z) {
		return math.NaN(), math.NaN(), math.NaN()
	}

	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy, z
}

func f(x, y float64) float64 {
	r := math.Hypot(x, y)
	return math.Sin(r) / r
}

/*
 * 鞍馬(2次方程式の複素数解が表す曲面)
 */
func f2(x, y float64) float64 {
	r := 0.1 * x * x - 0.1 * y * y - x + y
	return r * 0.02
}
