package main

import (
	"fmt"
)
const MAKS = 127
type Tabel [MAKS]rune

func isiArray(tab *Tabel, n *int) {
	var kar rune
	*n = 0
	fmt.Print("Masukkan karakter:")
	fmt.Scanf("%c", &kar)
	for kar != '.' {
		tab[*n] = kar
		*n++
		fmt.Scanf("%c", &kar) 
	}
}

func cetakArray(tab Tabel, n int) {
	fmt.Print("Teks: ")
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
	var tab Tabel
	var n int
	isiArray(&tab, &n)
	cetakArray(tab, n)

	if palindrom(tab, n) {
		fmt.Println("Palindrom: ?True")
	} else {
		fmt.Println("Palindrom: ?False")
	}

	fmt.Print("Reverse teks: ")
	balik := balikArray(&tab, n)
	for i := 0; i < n; i++ {
		fmt.Printf("%c", balik[i])
	}
	fmt.Println()
} 