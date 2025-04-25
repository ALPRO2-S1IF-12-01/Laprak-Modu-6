package main

import (
	"fmt"
	"math"
)

type Titik struct {
	x, y int
}

type Lingkaran struct {
	pusat Titik
	r     int
}

func jarak(a, b Titik) float64 {
	dx := float64(a.x - b.x)
	dy := float64(a.y - b.y)
	return math.Sqrt(dx*dx + dy*dy)
}

func DL(l Lingkaran, t Titik) bool {
	return jarak(l.pusat, t) <= float64(l.r)
}

func main() {
	var cx1, cy1, r1 int
	var cx2, cy2, r2 int
	var x, y int
	fmt.Println("NAMA: MULIA AKBAR NANDA PRATAMA\nNIM: 103112400034")
	fmt.Scan(&cx1, &cy1, &r1)
	fmt.Scan(&cx2, &cy2, &r2)
	fmt.Scan(&x, &y)

	l1 := Lingkaran{Titik{cx1, cy1}, r1}
	l2 := Lingkaran{Titik{cx2, cy2}, r2}
	titik := Titik{x, y}

	dalam1 := DL(l1, titik)
	dalam2 := DL(l2, titik)

	if dalam1 && dalam2 {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if dalam1 {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if dalam2 {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}
