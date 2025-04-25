//Ahmad Ruba'i
//103112400074
package main

import (
	"fmt"
	"math"
)

func jarak(a int, b int, c int, d int) float64 {
	dx := float64(a - c)
	dy := float64(b - d)
	return math.Sqrt(dx*dx + dy*dy)
}

func didalam(cx int, cy int, r int, px int, py int) bool {
	jarakPusatKeTitik := jarak(cx, cy, px, py)
	return jarakPusatKeTitik < float64(r)
}

func main() {
	var cx1, cy1, r1, cx2, cy2, r2, px, py int

	fmt.Scanln(&cx1, &cy1, &r1)
	fmt.Scanln(&cx2, &cy2, &r2)
	fmt.Scanln(&px, &py)

	diDalam1 := didalam(cx1, cy1, r1, px, py)
	diDalam2 := didalam(cx2, cy2, r2, px, py)

	if diDalam1 && diDalam2 {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if diDalam1 {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if diDalam2 {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}