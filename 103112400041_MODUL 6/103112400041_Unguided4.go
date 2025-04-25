//BERTHA ADELA
//103112400041
package main

import (
    "fmt"
    "strings"
)

const NMAX int = 127

type tabel struct {
    arr [NMAX]rune
    m   int
}

func isiArray(t *tabel) {
    var c rune
    t.m = 0
    fmt.Print("Teks: ")
    for t.m < NMAX && c != '.' {
        fmt.Scanf("%c", &c)
        if c != '.' {
            t.arr[t.m] = c
            t.m=t.m+1
        }
    }
}


func cetakArray(t tabel) {
    for i := 0; i < t.m; i++ {
        fmt.Printf("%c", t.arr[i])
    }
    fmt.Println()
}

func balikanArray(t *tabel) {
    for i := 0; i < t.m/2; i++ {
        t.arr[i], t.arr[t.m-i-1] = t.arr[t.m-i-1], t.arr[i]
    }
}

func palindrom(t tabel) bool {
    // Membersihkan input untuk mengabaikan spasi dan simbol
    cleaned := ""
    for i := 0; i < t.m; i++ {
        if t.arr[i] != ' ' && t.arr[i] != '.' { // Abaikan spasi dan titik
            cleaned += strings.ToLower(string(t.arr[i]))
        }
    }

    // Periksa apakah string yang sudah dibersihkan adalah palindrom
    length := len(cleaned)
    for i := 0; i < length/2; i++ {
        if cleaned[i] != cleaned[length-i-1] {
            return false
        }
    }
    return true
}

func main() {
    var tab tabel
    isiArray(&tab)
    if palindrom(tab) {
        fmt.Println("Palindrom:",palindrom(tab))
    } else {
        fmt.Println("Palindrom:",palindrom(tab))
    }
}
