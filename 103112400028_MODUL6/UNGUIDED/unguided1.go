// MUAHMMAD GAMEL AL GHIFARI
// 103112400028
package main

import (
	"fmt"
	"math"
)

func lingkaran(x1, y1, x2, y2 int) float64 {
	return math.Sqrt(math.Pow(float64(x2-x1), 2) + math.Pow(float64(y2-y1), 2))
}

func main() {
	var x1, y1, r1 int
	fmt.Scan(&x1, &y1, &r1)
	var x2, y2, r2 int
	fmt.Scan(&x2, &y2, &r2)
	var xt, yt int
	fmt.Scan(&xt, &yt)
	jarakKeLingkaran1 := lingkaran(x1, y1, xt, yt)
	jarakKeLingkaran2 := lingkaran(x2, y2, xt, yt)

	if jarakKeLingkaran1 <= float64(r1) && jarakKeLingkaran2 <= float64(r2) {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if jarakKeLingkaran1 <= float64(r1) {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if jarakKeLingkaran2 <= float64(r2) {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}
