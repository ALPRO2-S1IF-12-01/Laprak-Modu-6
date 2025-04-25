package main

import (
	"fmt"
	"math"
)

type Titik struct {
	x, y int
}

type Lingkaran struct {
	pusat  Titik
	radius int
}

func hitungJarak(a, b Titik) float64 {
	return math.Sqrt(math.Pow(float64(a.x-b.x), 2) + math.Pow(float64(a.y-b.y), 2))
}

func diDalamLingkaran(t Titik, l Lingkaran) bool {
	return hitungJarak(t, l.pusat) <= float64(l.radius)
}

func main() {
	var l1, l2 Lingkaran
	var titik Titik

	// Input dari user
	fmt.Println("Masukkan data lingkaran 1 (x y radius):")
	fmt.Scan(&l1.pusat.x, &l1.pusat.y, &l1.radius)

	fmt.Println("Masukkan data lingkaran 2 (x y radius):")
	fmt.Scan(&l2.pusat.x, &l2.pusat.y, &l2.radius)

	fmt.Println("Masukkan titik sembarang (x y):")
	fmt.Scan(&titik.x, &titik.y)

	// Logika posisi titik terhadap dua lingkaran
	dalam1 := diDalamLingkaran(titik, l1)
	dalam2 := diDalamLingkaran(titik, l2)

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
