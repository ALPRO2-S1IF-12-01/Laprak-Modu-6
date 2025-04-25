// Savila Nur Fadilla
// 103112400031

package main

import "fmt"

const NMAX int = 127

type tabel [NMAX]rune

type Tabel struct {
    tab tabel
    m   int
}

// fungsi untuk mengisi array karakter
func isiArray(t *Tabel, n *int) {
    var c rune
    *n = 0
    fmt.Print("Masukkan karakter (akhiri dengan tanda TITIK) : ")
    for {
        fmt.Scanf("%c", &c)
        if c == '.' || *n >= NMAX {
            break
        }
        t.tab[*n] = c
        *n = *n + 1
    }
}

// fungsi untuk mencetak isi array
func cetakArray(t Tabel, n int) {
    for i := 0; i < n; i++ {
        fmt.Printf("%c", t.tab[i])
    }
    fmt.Println()
}

// fungsi untuk mereverse isi array
func balikanArray(t *Tabel, n int) {
    for i := 0; i < n/2; i++ {
        t.tab[i], t.tab[n-1-i] = t.tab[n-1-i], t.tab[i]
    }
}

// fungsi untuk mengecek apakah isi array adalah palindrome
func palindrome(t Tabel, n int) bool {
    for i := 0; i < n/2; i++ {
        if t.tab[i] != t.tab[n-1-i] {
            return false
        }
    }
    return true
}

func main() {
    var tab Tabel
    var m int

    isiArray(&tab, &m)

    fmt.Print("Teks		: ")
    cetakArray(tab, m)

    var reversed Tabel = tab
    balikanArray(&reversed, m)

    fmt.Print("Reverse teks	: ")
    cetakArray(reversed, m)

    fmt.Print("Palindrome	? ")
    if palindrome(tab, m) {
        fmt.Println("true")
    } else {
        fmt.Println("false")
    }
}
