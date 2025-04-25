package main
//Nufail Alauddin Tsaqif
//013112400084
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
	return math.Hypot(a.X-b.X, a.Y-b.Y)
}

func titikDalamLingkaran(l Lingkaran, t Titik) bool {
	return hitungJarak(l.Pusat, t) <= l.Radius
}

func inputLingkaran(nama string) Lingkaran {
	var x, y, r float64
	fmt.Printf("Masukkan %s (x y r): ", nama)
	fmt.Scan(&x, &y, &r)
	return Lingkaran{Pusat: Titik{X: x, Y: y}, Radius: r}
}

func inputTitik() Titik {
	var x, y float64
	fmt.Print("Masukkan koordinat titik (x y): ")
	fmt.Scan(&x, &y)
	return Titik{X: x, Y: y}
}

func main() {
	lingkaran1 := inputLingkaran("lingkaran 1")
	lingkaran2 := inputLingkaran("lingkaran 2")
	titik := inputTitik()

	dalamL1 := titikDalamLingkaran(lingkaran1, titik)
	dalamL2 := titikDalamLingkaran(lingkaran2, titik)

	if dalamL1 && dalamL2 {
		fmt.Println("Titik berada di dalam lingkaran 1 dan 2.")
	} else if dalamL1 {
		fmt.Println("Titik berada di dalam lingkaran 1 saja.")
	} else if dalamL2 {
		fmt.Println("Titik berada di dalam lingkaran 2 saja.")
	} else {
		fmt.Println("Titik berada di luar kedua lingkaran.")
	}
}
