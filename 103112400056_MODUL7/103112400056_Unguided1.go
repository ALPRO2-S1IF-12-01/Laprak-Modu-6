// Felix Pedrosa V 

package main

import (
    "fmt"
    "math"
)

// Fungsi untuk menghitung jarak antara dua titik
func hitungJarak(x1, y1, x2, y2 float64) float64 {
    return math.Sqrt(math.Pow(x1-x2, 2) + math.Pow(y1-y2, 2))
}

// Fungsi untuk memeriksa apakah titik berada di dalam lingkaran
func titikDalamLingkaran(centerX, centerY, radius, pointX, pointY float64) bool {
    return hitungJarak(centerX, centerY, pointX, pointY) <= radius
}

func main() {
    var centerX1, centerY1, radius1 float64
    var centerX2, centerY2, radius2 float64
    var pointX, pointY float64

    // Input untuk lingkaran pertama
    fmt.Scan(&centerX1, &centerY1, &radius1)
    // Input untuk lingkaran kedua
    fmt.Scan(&centerX2, &centerY2, &radius2)
    // Input untuk titik
    fmt.Scan(&pointX, &pointY)

    // Memeriksa apakah titik berada di dalam lingkaran
    isInCircle1 := titikDalamLingkaran(centerX1, centerY1, radius1, pointX, pointY)
    isInCircle2 := titikDalamLingkaran(centerX2, centerY2, radius2, pointX, pointY)

    // Menentukan dan mencetak hasil
    if isInCircle1 && isInCircle2 {
        fmt.Println("Titik di dalam lingkaran 1 dan 2")
    } else if isInCircle1 {
        fmt.Println("Titik di dalam lingkaran 1")
    } else if isInCircle2 {
        fmt.Println("Titik di dalam lingkaran 2")
    } else {
        fmt.Println("Titik di luar lingkaran 1 dan 2")
    }
}