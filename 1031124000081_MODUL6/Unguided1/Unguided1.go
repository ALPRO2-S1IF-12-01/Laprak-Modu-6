//RYAN AKEYLA NOVIANTO WIDODO
//103112400081

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
	return math.Sqrt(float64((p1.x-p2.x)*(p1.x-p2.x) + (p2.y-p2.y)*(p2.y-p2.y)))
}

func isPointInCircles(point Titik, circles [2]Lingkaran) []int {
	inside := []int{}
	for i, circle := range circles {
		if jarak(point, circle.Pusat) < float64(circle.Radius) {
			inside = append(inside, i+1)
		}
	}
	return inside
}

func main() {
	var circles [2]Lingkaran
	var point Titik
	fmt.Scan(&circles[0].Pusat.x, &circles[0].Pusat.y, &circles[0].Radius)
	fmt.Scan(&circles[1].Pusat.x, &circles[1].Pusat.y, &circles[1].Radius)
	fmt.Scan(&point.x, &point.y)

	insideCircles := isPointInCircles(point, circles)

	switch len(insideCircles) {
	case 0:
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	case 1:
		fmt.Println("Titik di dalam lingkaran %d\n", insideCircles[0])
	case 2:
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	}
}
