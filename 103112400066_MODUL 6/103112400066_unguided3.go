// DWI OKTA SURYANINGRUM
// 103112400066

package main

import "fmt"

// Struct Pertandingan menyimpan data dua klub dan daftar pemenang dari beberapa pertandingan
type Pertandingan struct {
    KlubA    string   // Menyimpan nama klub A
    KlubB    string   // Menyimpan nama klub B
    Pemenang []string // Slice untuk menyimpan hasil pemenang setiap pertandingan
}

func main() {
    var p Pertandingan // Inisialisasi variabel struct Pertandingan
    
    // Input nama klub
    fmt.Print("Klub A : ")
    fmt.Scan(&p.KlubA) // Meminta input nama Klub A
    fmt.Print("Klub B : ")
    fmt.Scan(&p.KlubB) // Meminta input nama Klub B
    
    // Proses input skor pertandingan berulang
    for i := 1; ; i++ { // Loop tanpa batas, dihentikan dengan kondisi tertentu
        var skorA, skorB int // Variabel untuk menyimpan skor tiap klub
        fmt.Printf("Pertandingan %d : ", i)
        fmt.Scan(&skorA, &skorB) // Input skor Klub A dan Klub B
        
        // Validasi skor negatif untuk menghentikan input
        if skorA < 0 || skorB < 0 {
            fmt.Println("Pertandingan selesai")
            break // Keluar dari loop jika skor negatif dimasukkan
        }
        
        // Menentukan pemenang berdasarkan skor
        switch {
        case skorA > skorB:
            p.Pemenang = append(p.Pemenang, p.KlubA) // Klub A menang
            fmt.Printf("Hasil %d : %s\n", i, p.KlubA)
        case skorB > skorA:
            p.Pemenang = append(p.Pemenang, p.KlubB) // Klub B menang
            fmt.Printf("Hasil %d : %s\n", i, p.KlubB)
        default:
            p.Pemenang = append(p.Pemenang, "Draw") // Skor imbang
            fmt.Printf("Hasil %d : Draw\n", i)
        }
    }
    
    // Menampilkan daftar pemenang dari semua pertandingan
    fmt.Println("\nDaftar Klub Pemenang:")
    for i, pemenang := range p.Pemenang {
        fmt.Printf("Pertandingan %d: %s\n", i+1, pemenang)
    }
}