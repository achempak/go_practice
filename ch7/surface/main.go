package main

import (
	"fmt"
	"io"
	"log"
	"math"
	"net/http"

	"goPractice/ch7/eval"
)

const (
	width, height = 600, 320            // canvas size in pixels
	cells         = 100                 // number of grid cells per dimension
	xyrange       = 30.0                // axis range
	xyscale       = width / 2 / xyrange // pixels per unit
	zscale        = height * 0.5        // pixels per z unit
	angle         = math.Pi / 6         // angle of x,y axes
)

var sin30, cos30 = math.Sin(angle), math.Cos(angle)

func parseAndCheck(s string) (eval.Expr, error) {
	if s == "" {
		return nil, fmt.Errorf("empty expression")
	}
	expr, err := eval.Parse(s)
	if err != nil {
		return nil, err
	}
	vars := make(map[eval.Var]bool)
	if err := expr.Check(vars); err != nil {
		return nil, err
	}
	for v := range vars {
		if v != "x" && v != "y" && v != "r" {
			return nil, fmt.Errorf("undefined variable: %s", v)
		}
	}
	return expr, nil
}

func plot(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	expr, err := parseAndCheck(r.Form.Get("expr"))
	if err != nil {
		http.Error(w, "bad expr: "+err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "image/svg+xml")
	surface(w, expr)
}

func surface(w io.Writer, expr eval.Expr) {
	fmt.Fprintf(w, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(i+1, j, expr)
			bx, by := corner(i, j, expr)
			cx, cy := corner(i, j+1, expr)
			dx, dy := corner(i+1, j+1, expr)
			fmt.Fprintf(w, "<polygon points='%g,%g,%g,%g,%g,%g,%g,%g'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Fprintf(w, "</svg>")
	// http.HandleFunc("/", handler)
	// log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func corner(i, j int, expr eval.Expr) (float64, float64) {
	// Find point (x,y) at corner of cell (i,j)
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)
	// Compute surface height z.
	z := f(x, y, expr)
	// Project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy).
	sx := width/2 + (x-y)*cos30*xyscale
	sy := height/2 + (x+y)*sin30*xyscale - z*zscale
	return sx, sy
}

func f(x, y float64, expr eval.Expr) float64 {
	r := math.Hypot(x, y) // distance from (0,0)
	env := eval.Env{"x": x, "y": y, "r": r}
	return expr.Eval(env)
}

func main() {
	http.HandleFunc("/plot", plot)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
