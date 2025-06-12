package main

import (
	"fmt"
	"unicode"
)

const MAKS = 127

type Tabel [MAKS]rune

func isiArray(tab *Tabel, n *int) {
	var kar rune
	*n = 0
	fmt.Println("Masukkan teks (akhiri dengan '.'):")

	for *n < MAKS {
		fmt.Scanf("%c", &kar)
		if kar == '.' {
			break
		}
		if !unicode.IsSpace(kar) {
			tab[*n] = unicode.ToLower(kar)
			*n++
		}
	}
}

func cetakArray(tab Tabel, n int) {
	fmt.Print("Teks yang diinput: ")
	for i := 0; i < n; i++ {
		fmt.Printf("%c", tab[i])
	}
	fmt.Println()
}

func balikArray(tab *Tabel, n int) Tabel {
	var hasil Tabel
	j := 0
	for i := n - 1; i >= 0; i-- {
		hasil[j] = tab[i]
		j++
	}
	return hasil
}

func palindrom(tab Tabel, n int) bool {
	balik := balikArray(&tab, n)
	for i := 0; i < n; i++ {
		if tab[i] != balik[i] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println("Program Pengecekan Palindrom")
	fmt.Println("==========================")

	var tab Tabel
	var n int
	isiArray(&tab, &n)
	if n == 0 {
		fmt.Println("Error: Teks kosong!")
		return
	}
	cetakArray(tab, n)
	fmt.Print("\nHasil Pengecekan:\n")
	if palindrom(tab, n) {
		fmt.Println("Teks tersebut adalah palindrom")
	} else {
		fmt.Println("Teks tersebut bukan palindrom")
	}
	fmt.Print("\nTeks terbalik: ")
	balik := balikArray(&tab, n)
	for i := 0; i < n; i++ {
		fmt.Printf("%c", balik[i])
	}
	fmt.Println()
}
