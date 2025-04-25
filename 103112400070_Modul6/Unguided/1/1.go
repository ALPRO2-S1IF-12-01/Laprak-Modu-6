//Achmad Zulvan Nur Hakim 103112400070
package main

import (
	"fmt"
	"math"
)

type Titik struct {
	x, y int
}

type L struct {
	cen Titik
	r   int
}

func jarak(p, q Titik) float64 {
	return math.Sqrt(math.Pow(float64(p.x-q.x), 2) + math.Pow(float64(p.y-q.y), 2))
}

func didalam(l L, p Titik) bool {
	return jarak(p, l.cen) <= float64(l.r)
}

func main() {
	var L1, L2 L
	var titik Titik

	fmt.Scan(&L1.cen.x, &L1.cen.y, &L1.r)
	fmt.Scan(&L2.cen.x, &L2.cen.y, &L2.r)
	fmt.Scan(&titik.x, &titik.y)

	dlmL1 := didalam(L1, titik)
	dlmL2 := didalam(L2, titik)

	if dlmL1 && dlmL2 {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if dlmL1 {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if dlmL2 {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}