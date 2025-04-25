//Achmad Zulvan Nur Hakim 103112400070
package main

import "fmt"

const NMAX int = 127

type tabel [NMAX]rune

func isiArray(t *tabel, n *int) {
    var char rune
    *n = 0
    for {
        fmt.Scanf("%c", &char)
        if char == '.' || *n >= NMAX {
            break
        }
        if char != ' ' && char != '\n' {
            t[*n] = char
            *n++
        }
    }
}

func cetakArray(t tabel, n int) {
    for i := 0; i < n; i++ {
        fmt.Printf("%c ", t[i])
    }
    fmt.Println()
}

func balikanArray(t *tabel, n int) {
    for i := 0; i < n/2; i++ {
        t[i], t[n-1-i] = t[n-1-i], t[i]
    }
}

func isPalindrom(t tabel, n int) bool {
    for i := 0; i < n/2; i++ {
        if t[i] != t[n-1-i] {
            return false
        }
    }
    return true
}

func main() {
    var tab tabel
    var m int

    fmt.Print("Teks : ")
    isiArray(&tab, &m)

    copyTab := tab

    fmt.Print("Reverse : ")
    balikanArray(&tab, m)
    cetakArray(tab, m)

    fmt.Printf("Palindrom ? %t\n", isPalindrom(copyTab, m))
}