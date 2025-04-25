package main

import "fmt"

const NMAX int = 127

type tabel [NMAX]rune

func isiArray(t *tabel, n *int) {
    var char rune
    *n = 0
    fmt.Print("Teks    :")
    for {
        _, err := fmt.Scanf("%c", &char)
        if err != nil || char == '.' || *n >= NMAX {
            break
        }
        if char != ' ' && char != '\n' && char != '\r' {
            t[*n] = char
            fmt.Printf(" %c", char)
            *n++
        }
    }
    fmt.Println()
}

func cetakArray(t tabel, n int, label string) {
    fmt.Print(label)
    for i := 0; i < n; i++ {
        fmt.Printf(" %c", t[i])
    }
    fmt.Println()
}

func balikkanArray(t *tabel, n int) {
    for i := 0; i < n/2; i++ {
        t[i], t[n-1-i] = t[n-1-i], t[i]
    }
}

func palindrom(t tabel, n int) bool {
    for i := 0; i < n/2; i++ {
        if t[i] != t[n-1-i] {
            return false
        }
    }
    return true
}

func main() {
    var tab, tabAsli tabel
    var m int

    // Isi array tab
    isiArray(&tab, &m)
    
    // Simpan versi asli
    copy(tabAsli[:], tab[:m])
    
    // Periksa palindrom
    fmt.Print("Palindrom   ? ", palindrom(tab, m), "\n\n")
    
    // Balikkan array dan tampilkan
    balikkanArray(&tab, m)
    cetakArray(tab, m, "Reverse teks   :")
}