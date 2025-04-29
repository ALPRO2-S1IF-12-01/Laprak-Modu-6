package main

import (
	"fmt"
	"math"
)

type Titik struct {
	x, y float64
}

type Lingkaran struct {
	titikPusat Titik
	radius     float64
}

func jarak(p, q Titik) float64 {
	return math.Sqrt(math.Pow(p.x-q.x, 2) + math.Pow(p.y-q.y, 2))
}

func didalam(p Titik, l Lingkaran) bool {
	return jarak(p, l.titikPusat) <= l.radius
}

func main() {

	var x1, y1, r1 float64
	var x2, y2, r2 float64
	var xp, yp float64

	fmt.Print("Masukkan x1 y1 r1: ")
	fmt.Scan(&x1, &y1, &r1)
	fmt.Print("Masukkan x2 y2 r2: ")
	fmt.Scan(&x2, &y2, &r2)
	fmt.Print("Masukkan xp yp: ")
	fmt.Scan(&xp, &yp)

	ling1 := Lingkaran{Titik{x1, y1}, r1}
	ling2 := Lingkaran{Titik{x2, y2}, r2}
	titik := Titik{xp, yp}

	in1 := didalam(titik, ling1)
	in2 := didalam(titik, ling2)

	if in1 && in2 {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if in1 {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if in2 {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}
