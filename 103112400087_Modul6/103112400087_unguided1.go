package main

import (
	"fmt"
	"math"
)

type titik struct {
	x, y int
}
type lingkaran struct {
	pusat  titik
	radius int
}

func jarak(p, q titik) float64 {
	return math.Sqrt(float64((p.x-q.x)*(p.x-q.x) + (p.y-q.y)*(p.y-q.y)))
}
func didalam(c lingkaran, p titik) bool {
	return jarak(c.pusat, p) <= float64(c.radius)
}
func main() {
	var l1, l2 lingkaran
	var p titik

	fmt.Print("Masukkan pusat dan radius lingkaran 1 (x y r): ")
	fmt.Scan(&l1.pusat.x, &l1.pusat.y, &l1.radius)

	fmt.Print("Masukkan pusat dan radius lingkaran 2 (x y r): ")
	fmt.Scan(&l2.pusat.x, &l2.pusat.y, &l2.radius)

	fmt.Print("Masukkan titik sembarang (x y): ")
	fmt.Scan(&p.x, &p.y)

	inL1 := didalam(l1, p)
	inL2 := didalam(l2, p)

	if inL1 && inL2 {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if inL1 {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if inL2 {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}
