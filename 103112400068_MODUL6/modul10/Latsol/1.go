package main

import "fmt"

func main() {

	var x1, y1, r1, x2, y2, r2, x, y float64

	fmt.Scan(&x1, &y1, &r1)

	fmt.Scan(&x2, &y2, &r2)

	fmt.Scan(&x, &y)

	jarak1 := (x-x1)*(x-x1) + (y-y1)*(y-y1)

	jarak2 := (x-x2)*(x-x2) + (y-y2)*(y-y2)

	diDalamLingkaran1 := jarak1 <= r1*r1
	diDalamLingkaran2 := jarak2 <= r2*r2

	if diDalamLingkaran1 && diDalamLingkaran2 {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if diDalamLingkaran1 {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if diDalamLingkaran2 {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}
