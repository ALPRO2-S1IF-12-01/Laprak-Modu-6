package main

import (
	"fmt"
	"math"
)

func jarakTitik(a, b, x, y float64) float64 {
	return math.Sqrt(math.Pow(a-x, 2) + math.Pow(b-y, 2))
}

func cekLingkaran(a, b, c, x, y float64) bool {
	return jarakTitik(a, b, x, y) <= c
}

func main() {
	var x1, x2, y1, y2, r1, r2, x, y float64
	fmt.Print("Masukkan koordinat lingkaran 1: ")
	fmt.Scan(&x1, &y1, &r1)
	fmt.Print("Masukkan koordinat lingkaran 2: ")
	fmt.Scan(&x2, &y2, &r2)
	fmt.Print("Masukkan koordinat titik: ")
	fmt.Scan(&x, &y)
	jarakLingkaran1 := cekLingkaran(x1, y1, r1, x, y)
	jarakLingkaran2 := cekLingkaran(x2, y2, r2, x, y)
	if jarakLingkaran1 && jarakLingkaran2 {
		fmt.Print("Titik di dalam lingkaran 1 dan 2")
	} else if jarakLingkaran1 {
		fmt.Print("Titik di dalam lingkaran 1")
	} else if jarakLingkaran2 {
		fmt.Print("Titik di dalam lingkaran 2")
	} else {
		fmt.Print("Titik di luar lingkaran 1 dan 2")
	}
}

/* 
Dimas Ramadhani
103112400065
*/
