package main

//103112400049 Hisyam Nurdiatmoko

import (
	"fmt"
	"math"
)

type titik struct {
	x, y int
}

type lingkaran struct {
	cx, cy, r int
}

func jarak(p, q titik) float64 {
	return math.Sqrt(float64((p.x-q.x)*(p.x-q.x) + (p.y-q.y)*(p.y-q.y)))
}

func didalam(c lingkaran, p titik) bool {
	return jarak(titik{c.cx, c.cy}, p) <= float64(c.r)
}

func main() {
	var c1, c2 lingkaran
	var hisyam titik
	fmt.Scan(&c1.cx, &c1.cy, &c1.r)
	fmt.Scan(&c2.cx, &c2.cy, &c2.r)
	fmt.Scan(&hisyam.x, &hisyam.y)
	diC1 := didalam(c1, hisyam)
	diC2 := didalam(c2, hisyam)
	if diC1 && diC2 {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if diC1 {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if diC2 {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}
