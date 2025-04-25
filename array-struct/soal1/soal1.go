//ABID FADHILAH M
//103112400046
package main

import (
	"fmt"
	"math"
)

func dalamLingkaran(x, y, r, xt, yt int) bool {
	return math.Sqrt(float64((xt-x)*(xt-x)+(yt-y)*(yt-y))) <= float64(r)
}

func main() {
	var x1, y1, r1, x2, y2, r2, xt, yt int
	fmt.Scan(&x1, &y1, &r1, &x2, &y2, &r2, &xt, &yt)

	dalam1, dalam2 := dalamLingkaran(x1, y1, r1, xt, yt), dalamLingkaran(x2, y2, r2, xt, yt)
	switch {
	case dalam1 && dalam2:
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	case dalam1:
		fmt.Println("Titik di dalam lingkaran 1")
	case dalam2:
		fmt.Println("Titik di dalam lingkaran 2")
	default:
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}

