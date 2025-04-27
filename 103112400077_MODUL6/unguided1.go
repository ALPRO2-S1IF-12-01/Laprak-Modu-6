package main
import (
"fmt"
"math"
)
func j(px, py, sx, sy float64) float64 {
return math.Sqrt(math.Pow(px-sx, 2) + math.Pow(py-sy, 2))
}
func dalam(cx, cy, r, x, y float64) bool {
return j(cx, cy, x, y) <= r
}
func main() {
var centralx1, centraly1, radius1, centralx2, centraly2, radius2, x, y float64
fmt.Scan(&centralx1, &centraly1, &radius1)
fmt.Scan(&centralx2, &centraly2, &radius2)
fmt.Scan(&x, &y)
jarakLingkaran1 := dalam(centralx1, centraly1, radius1, x, y)
jarakLingkaran2 := dalam(centralx2, centraly2, radius2, x, y)
if jarakLingkaran1 && jarakLingkaran2 {
fmt.Print("Titik di dalam lingkaran 1 dan 2")
} else if jarakLingkaran1 {
fmt.Print("Titik di dalam lingkaran 1")
} else if jarakLingkaran2 {
fmt.Print("Titik di dalam lingkaran 2")
} else {
fmt.Print("Titik di luar lingkaran 1 dan 2")
}
}
