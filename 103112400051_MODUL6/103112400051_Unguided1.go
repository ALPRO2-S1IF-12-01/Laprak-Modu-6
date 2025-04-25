package main

import "fmt"

type Titik struct {
	x,y int
}

type Lingkaran struct {
	pusat Titik
	r int
}

func jarakKuadrat(a, b Titik) int {
	dx := a.x - b.x
	dy := a.y - b.y
	return dx*dx + dy*dy
}

func dalamLingkaran(l Lingkaran, t Titik) bool {
	return jarakKuadrat(l.pusat, t) <= l.r*l.r
}

func main() {
	var l1, l2 Lingkaran
	var titik Titik

	fmt.Scan(&l1.pusat.x, &l1.pusat.y, &l1.r)
	fmt.Scan(&l2.pusat.x, &l2.pusat.y, &l2.r)
	fmt.Scan(&titik.x, &titik.y)

	inL1 := dalamLingkaran(l1, titik)
	inL2 := dalamLingkaran(l2, titik)

	switch {
	case inL1 && inL2:
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	case inL1:
		fmt.Println("Titik di dalam lingkaran 1")
	case inL2:
		fmt.Println("Titik di dalam lingkaran 2")
	default:
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}
