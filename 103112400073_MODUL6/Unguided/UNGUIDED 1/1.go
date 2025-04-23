//Muhammad Zaky Mubarok
package main

import "fmt"

type Titik struct {
	x, y int
}

type Lingkaran struct {
	pusat   Titik
	jariJari int
}

func (l Lingkaran) berisi(t Titik) bool {
	dx := t.x - l.pusat.x
	dy := t.y - l.pusat.y
	jarakKuadrat := dx*dx + dy*dy
	jariJariKuadrat := l.jariJari * l.jariJari
	return jarakKuadrat <= jariJariKuadrat
}

func main() {
	var l1, l2 Lingkaran
	var t Titik

	fmt.Scan(&l1.pusat.x, &l1.pusat.y, &l1.jariJari)
	fmt.Scan(&l2.pusat.x, &l2.pusat.y, &l2.jariJari)
	fmt.Scan(&t.x, &t.y)

	diL1 := l1.berisi(t)
	diL2 := l2.berisi(t)

	if diL1 && diL2 {
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	} else if diL1 {
		fmt.Println("Titik di dalam lingkaran 1")
	} else if diL2 {
		fmt.Println("Titik di dalam lingkaran 2")
	} else {
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}