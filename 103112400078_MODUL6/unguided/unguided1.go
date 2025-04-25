package main
// Mohammad Reyhan Aretha Fatin
// 103112400078
import (
	"fmt"
	"math"
)

type Titik struct {
	X, Y float64
}

type Lingkaran struct {
	Pusat  Titik
	Radius float64
}

func hitungJarak(a, b Titik) float64 {
	return math.Sqrt(math.Pow(a.X-b.X, 2) + math.Pow(a.Y-b.Y, 2))
}

func diDalamLingkaran(l Lingkaran, t Titik) bool {
	return hitungJarak(l.Pusat, t) <= l.Radius
}

func main() {
	var l1, l2 Lingkaran
	var titik Titik

	fmt.Print("Masukkan koordinat dan radius lingkaran 1 (x y r): ")
	fmt.Scan(&l1.Pusat.X, &l1.Pusat.Y, &l1.Radius)

	fmt.Print("Masukkan koordinat dan radius lingkaran 2 (x y r): ")
	fmt.Scan(&l2.Pusat.X, &l2.Pusat.Y, &l2.Radius)

	fmt.Print("Masukkan koordinat titik (x y): ")
	fmt.Scan(&titik.X, &titik.Y)

	dalamL1 := diDalamLingkaran(l1, titik)
	dalamL2 := diDalamLingkaran(l2, titik)

	switch {
	case dalamL1 && dalamL2:
		fmt.Println("Titik di dalam lingkaran 1 dan 2")
	case dalamL1:
		fmt.Println("Titik di dalam lingkaran 1")
	case dalamL2:
		fmt.Println("Titik di dalam lingkaran 2")
	default:
		fmt.Println("Titik di luar lingkaran 1 dan 2")
	}
}
