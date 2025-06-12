package main

import (
	"fmt"
	"math"
)

type Titik struct {
	x, y int
}

type Lingkaran struct {
	pusat Titik
	r     int
}

func jarak(a, b Titik) float64 {
	dx := float64(a.x - b.x)
	dy := float64(a.y - b.y)
	return math.Sqrt(dx*dx + dy*dy)
}

func DL(l Lingkaran, t Titik) bool {
	return jarak(l.pusat, t) <= float64(l.r)
}

func main() {
	fmt.Println("Program Pengecekan Titik dalam Lingkaran")
	fmt.Println("======================================")

	var cx1, cy1, r1 int
	var cx2, cy2, r2 int
	var x, y int
	fmt.Println("\nLingkaran 1:")
	fmt.Print("Pusat (x y): ")
	fmt.Scan(&cx1, &cy1)
	fmt.Print("Jari-jari: ")
	fmt.Scan(&r1)
	fmt.Println("\nLingkaran 2:")
	fmt.Print("Pusat (x y): ")
	fmt.Scan(&cx2, &cy2)
	fmt.Print("Jari-jari: ")
	fmt.Scan(&r2)
	fmt.Println("\nTitik yang akan dicek:")
	fmt.Print("Koordinat (x y): ")
	fmt.Scan(&x, &y)
	if r1 <= 0 || r2 <= 0 {
		fmt.Println("\nError: Jari-jari harus lebih besar dari 0!")
		return
	}
	l1 := Lingkaran{Titik{cx1, cy1}, r1}
	l2 := Lingkaran{Titik{cx2, cy2}, r2}
	titik := Titik{x, y}
	dalam1 := DL(l1, titik)
	dalam2 := DL(l2, titik)
	fmt.Println("\nHasil Pengecekan:")
	fmt.Printf("Titik (%d, %d) ", x, y)
	if dalam1 && dalam2 {
		fmt.Println("berada di dalam lingkaran 1 dan 2")
	} else if dalam1 {
		fmt.Println("berada di dalam lingkaran 1")
	} else if dalam2 {
		fmt.Println("berada di dalam lingkaran 2")
	} else {
		fmt.Println("berada di luar lingkaran 1 dan 2")
	}
}
