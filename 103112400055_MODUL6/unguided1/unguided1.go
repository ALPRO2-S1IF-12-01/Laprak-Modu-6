//Feros Pedrosa Valentino

package main

import (
	"fmt"
	"math"
)

type Titik struct {
	x int
	y int
}

type Lingkaran struct {
	Pusat  Titik
	Radius int
}

func jarak(p1, p2 Titik) float64 {
	return math.Sqrt(float64((p1.x-p2.x)*(p1.x-p2.x) + (p1.y-p2.y)*(p1.y-p2.y)))
}

func adalahTitikDalamLingkaran(titik Titik, lingkaran Lingkaran) bool {
	return jarak(titik, lingkaran.Pusat) < float64(lingkaran.Radius)
}

func main() {
	var l1, l2 Lingkaran
	var t Titik

	fmt.Scan(&l1.Pusat.x, &l1.Pusat.y, &l1.Radius)

	fmt.Scan(&l2.Pusat.x, &l2.Pusat.y, &l2.Radius)

	fmt.Scan(&t.x, &t.y)

	dalamLingkaran1 := adalahTitikDalamLingkaran(t, l1)
	dalamLingkaran2 := adalahTitikDalamLingkaran(t, l2)

	if dalamLingkaran1 && dalamLingkaran2 {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if dalamLingkaran1 {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if dalamLingkaran2 {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}
