package main

import (
	"fmt"
	"math"
)

type titik struct {
	x, y int
}

type lingkaran struct {
	pusatLingkaran titik
	radius int
}

func main() {
	var (
		lingkaran1, lingkaran2 lingkaran
		p titik
		dalamLingkaran1, dalamLingkaran2 bool
	)
	fmt.Scan(&lingkaran1.pusatLingkaran.x, &lingkaran1.pusatLingkaran.y, &lingkaran1.radius)
	fmt.Scan(&lingkaran2.pusatLingkaran.x, &lingkaran2.pusatLingkaran.y, &lingkaran2.radius)
	fmt.Scan(&p.x, &p.y)

	dalamLingkaran1 = didalam(lingkaran1, p)
	dalamLingkaran2 = didalam(lingkaran2, p)

	if dalamLingkaran1 == true && dalamLingkaran2 == true {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if dalamLingkaran1 == true {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if dalamLingkaran2 == true {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}

func jarak(p, q titik) float64 {
	cx := p.x - q.x
	cy := p.y - q.y

	return math.Sqrt(float64(cx*cx + cy*cy))
}

func didalam(c lingkaran, p titik) bool {
	jarakMenujuPusat := jarak(c.pusatLingkaran, p)
	return jarakMenujuPusat <= float64(c.radius)
}
