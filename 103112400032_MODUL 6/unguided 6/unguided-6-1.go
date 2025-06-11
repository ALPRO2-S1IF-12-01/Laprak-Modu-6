// daffa tsaqifna f
package main

import (
	"fmt"
	"math"
)

type coor struct {
	x, y float64
}
type lingkaran struct {
	titik  coor
	radius float64
}

func jarak(a, b, c, d float64) float64 {
	return math.Sqrt(math.Pow(a-c, 2) + math.Pow(b-d, 2))
}
func isInside(circle lingkaran, point coor) bool {
	d := jarak(circle.titik.x, circle.titik.y, point.x, point.y)
	return d <= circle.radius
}
func main() {
	var cx1, cx2, cy1, cy2, r1, r2, x, y float64
	fmt.Scan(&cx1, &cy1, &r1)
	fmt.Scan(&cx2, &cy2, &r2)
	fmt.Scan(&x, &y)
	l1 := lingkaran{titik: coor{cx1, cy1}, radius: r1}
	l2 := lingkaran{titik: coor{cx2, cy2}, radius: r2}
	titik := coor{x, y}
	in1 := isInside(l1, titik)
	in2 := isInside(l2, titik)
	switch {
	case in1 && in2:
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	case in1:
		fmt.Println("Titik di dalam lingkaran 1")
	case in2:
		fmt.Println("Titik di dalam lingkaran 2")
	default:
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}
