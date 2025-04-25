package main

//103112400050
import (
	"fmt"
	"math"
)

type Titik struct{ x, y int }
type Lingkaran struct{ x, y, r int }

func jarak(a, b Titik) float64 {
	dx := a.x - b.x
	dy := a.y - b.y
	return math.Sqrt(float64(dx*dx + dy*dy))
}

func dalam(c Lingkaran, t Titik) bool {
	return jarak(Titik{c.x, c.y}, t) <= float64(c.r)
}

func main() {
	var c1, c2 Lingkaran
	var A Titik

	fmt.Scan(&c1.x, &c1.y, &c1.r)
	fmt.Scan(&c2.x, &c2.y, &c2.r)
	fmt.Scan(&A.x, &A.y)

	d1 := dalam(c1, A)
	d2 := dalam(c2, A)

	if d1 && d2 {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if d1 {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if d2 {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}
