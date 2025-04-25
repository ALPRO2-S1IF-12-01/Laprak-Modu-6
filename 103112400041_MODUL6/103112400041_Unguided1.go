//BERTHA ADELA
//103112400041
package main
import (
	"fmt"
	"math"
)
type Titik struct {
    x, y float64
}

type Lingkaran struct {
    pusat Titik
    radius float64
}
func main() {
	var lingkaran1, lingkaran2 Lingkaran
    var titik Titik
	
	fmt.Scanln(&lingkaran1.pusat.x, &lingkaran1.pusat.y, &lingkaran1.radius)
    fmt.Scanln(&lingkaran2.pusat.x, &lingkaran2.pusat.y, &lingkaran2.radius)
    fmt.Scanln(&titik.x, &titik.y)

	if Didalam(lingkaran1, titik) && Didalam(lingkaran2, titik) {
        fmt.Print("Titik di dalam lingkaran 1 dan 2")
    } else if Didalam(lingkaran1, titik) {
        fmt.Print("Titik di dalam lingkaran 1")
    } else if Didalam(lingkaran2, titik) {
        fmt.Print("Titik di dalam lingkaran 2")
    } else {
        fmt.Print("Titik di luar lingkaran 1 dan 2")
    }
}

func Jarak(p, q Titik) float64 {
    return math.Sqrt((q.x-p.x)*(q.x-p.x) + (q.y-p.y)*(q.y-p.y))
}

func Didalam(c Lingkaran, p Titik) bool {
    return Jarak(c.pusat, p) <= c.radius
}