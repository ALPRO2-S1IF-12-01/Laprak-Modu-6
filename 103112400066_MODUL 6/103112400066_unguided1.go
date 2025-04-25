// DWI OKTA SURYANINGRUM
// 103112400066

package main

import "fmt"

// mendefinisikan struct untuk Titik (x,y) dan Lingkaran (pusat Titik, radius)
type Titik struct { X, Y int }
type Lingkaran struct { Pusat Titik; Radius int }

func main() {
    // mendeklarasi variabel untuk input
    var l1, l2 Lingkaran
    var p Titik

    // membaca input lingkaran 1 (x y radius)
    fmt.Scan(&l1.Pusat.X, &l1.Pusat.Y, &l1.Radius)
    
    // membaca input lingkaran 2 (x y radius)
    fmt.Scan(&l2.Pusat.X, &l2.Pusat.Y, &l2.Radius)
    
    // membaca input titik yang akan dicek (x y)
    fmt.Scan(&p.X, &p.Y)

    // membuat fungsi untuk cek apakah titik di dalam lingkaran
    cekLingkaran := func(p Titik, l Lingkaran) bool {
        dx := p.X - l.Pusat.X  // selisih koordinat x
        dy := p.Y - l.Pusat.Y  // selisih koordinat y
        // Jika jarak^2 <= radius^2 maka titik ada di dalam
        return dx*dx + dy*dy <= l.Radius*l.Radius
    }

    // cek posisi titik terhadap kedua lingkaran
    inL1 := cekLingkaran(p, l1)  // apakah di lingkaran 1?
    inL2 := cekLingkaran(p, l2)  // apakah di lingkaran 2?

    // menentukan output berdasarkan beberapa kondisi
    switch {
    case inL1 && inL2:
        fmt.Println("Titik di dalam lingkaran 1 dan 2")
    case inL1:
        fmt.Println("Titik di dalam lingkaran 1")
    case inL2:
        fmt.Println("Titik di dalam lingkaran 2")
    default:
        fmt.Println("Titik di luar lingkaran 1 dan 2")
    }
}