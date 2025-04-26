package main

// Muhammad Faris Rachmadi
// 103112400079

import (
	"fmt"
	"math"
)

type Titik struct {
	x, y int
}

type Lingkaran struct {
	center Titik
	radius int
}

func jarak(p, q Titik) float64 {
	return math.Sqrt(float64((p.x-q.x)*(p.x-q.x) + (p.y-q.y)*(p.y-q.y)))
}

func didalam(c Lingkaran, p Titik) bool {
	return jarak(c.center, p) <= float64(c.radius)
}

func main() {
	var c1, c2 Lingkaran
	fmt.Println("Masukkan koordinat pusat dan radius lingkaran 1 (cx cy r):")
	fmt.Scan(&c1.center.x, &c1.center.y, &c1.radius)

	fmt.Println("Masukkan koordinat pusat dan radius lingkaran 2 (cx cy r):")
	fmt.Scan(&c2.center.x, &c2.center.y, &c2.radius)

	var p Titik
	fmt.Println("Masukkan koordinat titik sembarang (x y):")
	fmt.Scan(&p.x, &p.y)

	diDalam1 := didalam(c1, p)
	diDalam2 := didalam(c2, p)

	switch {
	case diDalam1 && diDalam2:
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	case diDalam1:
		fmt.Println("Titik di dalam lingkaran 1")
	case diDalam2:
		fmt.Println("Titik di dalam lingkaran 2")
	default:
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}
