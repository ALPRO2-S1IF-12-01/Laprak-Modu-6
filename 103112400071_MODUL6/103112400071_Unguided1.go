// Raihan Adi Arba
// 103112400071

package main

import (
	"fmt"
	"math"
)

type titik struct {
	x, y int
}

type lingkaran struct {
	pusatlingkaran titik
	radius         int
}

func jarak(p, q titik) float64 {
	cx := p.x - q.x
	cy := p.y - q.y
	return math.Sqrt(float64(cx*cx + cy*cy))
}

func didalam(c lingkaran, p titik) bool {
	jarakkepusat := jarak(c.pusatlingkaran, p)
	return jarakkepusat <= float64(c.radius)
}

func main() {
	var (
		lingkaran1, lingkaran2       lingkaran
		p                            titik
		dlmlingkaran1, dlmlingkaran2 bool
	)

	fmt.Scan(&lingkaran1.pusatlingkaran.x, &lingkaran1.pusatlingkaran.y, &lingkaran1.radius)
	fmt.Scan(&lingkaran2.pusatlingkaran.x, &lingkaran2.pusatlingkaran.y, &lingkaran2.radius)
	fmt.Scan(&p.x, &p.y)

	dlmlingkaran1 = didalam(lingkaran1, p)
	dlmlingkaran2 = didalam(lingkaran2, p)

	switch {
	case dlmlingkaran1 && dlmlingkaran2:
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	case dlmlingkaran1:
		fmt.Println("Titik di dalam lingkaran 1")
	case dlmlingkaran2:
		fmt.Println("Titik di dalam lingkaran 2")
	default:
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}
