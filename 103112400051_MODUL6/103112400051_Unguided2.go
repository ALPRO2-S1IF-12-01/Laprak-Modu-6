package main

import "fmt"

func main() {
    const maxSize = 100
    var arr [maxSize]int
    var n, x, index, num int
    
    fmt.Print("Masukkan jumlah elemen : ")
    fmt.Scan(&n)
    
    fmt.Println("Masukkan", n, "elemen array:")
    for i := 0; i < n; i++ {
        fmt.Scan(&arr[i])
    }
    
    // a. Tampilkan seluruh array
    fmt.Println("\na. Isi array lengkap:")
    for i := 0; i < n; i++ {
        fmt.Print(arr[i], " ")
    }
    fmt.Println()
    
    // b. Indeks ganjil
    fmt.Println("\nb. Elemen indeks ganjil:")
    for i := 1; i < n; i += 2 {
        fmt.Print(arr[i], " ")
    }
    fmt.Println()
    
    // c. Indeks genap
    fmt.Println("\nc. Elemen indeks genap:")
    for i := 0; i < n; i += 2 {
        fmt.Print(arr[i], " ")
    }
    fmt.Println()
    
    // d. Kelipatan x
    fmt.Print("\nd. Masukkan nilai x untuk indeks kelipatan: ")
    fmt.Scan(&x)
    fmt.Println("Elemen dengan indeks kelipatan", x, ":")
    for i := 0; i < n; i++ {
        if i%x == 0 {
            fmt.Print(arr[i], " ")
        }
    }
    fmt.Println()
    
    // e. Hapus elemen
    fmt.Print("\ne. Masukkan indeks yang akan dihapus (0-", n-1, "): ")
    fmt.Scan(&index)
    for i := index; i < n-1; i++ {
        arr[i] = arr[i+1]
    }
    n--
    fmt.Println("Array setelah penghapusan:")
    for i := 0; i < n; i++ {
        fmt.Print(arr[i], " ")
    }
    fmt.Println()
    
    // f. Rata-rata 
    sum := 0
    for i := 0; i < n; i++ {
        sum += arr[i]
    }
    fmt.Printf("\nf. Rata-rata: %.2f\n", float64(sum)/float64(n))
    
    // g. Standar deviasi
    var variance float64
    avg := float64(sum) / float64(n)
    for i := 0; i < n; i++ {
        diff := float64(arr[i]) - avg
        variance += diff * diff
    }
    fmt.Printf("g. Standar deviasi : %.2f\n", sqrtApprox(variance/float64(n)))
    
    // h. Frekuensi bilangan
    fmt.Print("\nh. Masukkan bilangan yang dicari: ")
    fmt.Scan(&num)
    count := 0
    for i := 0; i < n; i++ {
        if arr[i] == num {
            count++
        }
    }
    fmt.Println("Frekuensi", num, ":", count)
}

func sqrtApprox(x float64) float64 {
    z := 1.0
    for i := 0; i < 10; i++ {
        z -= (z*z - x) / (2 * z)
    }
    return z
}