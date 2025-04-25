//Anastasia Adinda Narendra Indrianto
package main
import (
    "fmt"
    "math"
)

type Lingkaran struct {
    x, y, r float64
}
type Titik struct {
    x, y float64
}
func hitungJarakAnastasia(x1, y1, x2, y2 float64) float64 {
    return math.Sqrt(math.Pow(x1-x2, 2) + math.Pow(y1-y2, 2))
}
func titikDalamLingkaran(lingkaran Lingkaran, titik Titik) bool {
    return hitungJarakAnastasia(lingkaran.x, lingkaran.y, titik.x, titik.y) <= lingkaran.r
}

func main() {
    var l1, l2 Lingkaran
    var t Titik

    fmt.Print("Masukkan pusat dan radius lingkaran 1 (x, y, radius): ")
    fmt.Scan(&l1.x, &l1.y, &l1.r)
    fmt.Print("Masukkan pusat dan radius lingkaran 2 (x, y, radius): ")
    fmt.Scan(&l2.x, &l2.y, &l2.r)
    fmt.Println("Masukkan koordinat titik Anastasia (x, y):")
    fmt.Scan(&t.x, &t.y)
    lingkaran1 := titikDalamLingkaran(l1, t)
    lingkaran2 := titikDalamLingkaran(l2, t)

    if lingkaran1 && lingkaran2 {
        fmt.Println("Titik di dalam lingkaran 1 dan 2")
    } else if lingkaran1 {
        fmt.Println("Titik di dalam lingkaran 1")
    } else if lingkaran2 {
        fmt.Println("Titik di dalam lingkaran 2")
    } else {
        fmt.Println("Titik di luar lingkaran 1 dan 2")
    }
}
