// M.HANIF AL FAIZ
// 103112400042
package main

import (
	"fmt"
	"math"
)

func jarak(x1, y1, x2, y2 int) float64 {
	return math.Sqrt(float64((x2-x1)*(x2-x1) + (y2-y1)*(y2-y1)))
}

func main() {
	var x1, y1, r1 int
	var x2, y2, r2 int
	var xt, yt int

	fmt.Scan(&x1, &y1, &r1)
	fmt.Scan(&x2, &y2, &r2)
	fmt.Scan(&xt, &yt)

	jarakKeLingkaran1 := jarak(x1, y1, xt, yt)
	jarakKeLingkaran2 := jarak(x2, y2, xt, yt)

	if jarakKeLingkaran1 < float64(r1) && jarakKeLingkaran2 < float64(r2) {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if jarakKeLingkaran1 < float64(r1) {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if jarakKeLingkaran2 < float64(r2) {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}
