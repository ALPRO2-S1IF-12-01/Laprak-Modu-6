package main
//Setyo Nugroho
//103112400024

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

func hitungJarak(a, b Titik) float64 {
	selisihX := float64(a.x - b.x)
	selisihY := float64(a.y - b.y)
	return math.Sqrt(selisihX*selisihX + selisihY*selisihY)
}

func cekDalam(l Lingkaran, t Titik) bool {
	return hitungJarak(l.pusat, t) <= float64(l.r)
}

func main() {
	var ling1, ling2 Lingkaran
	var titik Titik
	fmt.Scan(&ling1.pusat.x, &ling1.pusat.y, &ling1.r)
	fmt.Scan(&ling2.pusat.x, &ling2.pusat.y, &ling2.r)
	fmt.Scan(&titik.x, &titik.y)
	dalam1 := cekDalam(ling1, titik)
	dalam2 := cekDalam(ling2, titik)

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
