//Anastasia Adinda Narendra Indrianto
package main

import "fmt"

const NMAX int = 127
type tabel [NMAX]rune
func AnastasiaIsiTeks(t *tabel, jumlah *int) {
	var karakter rune
	*jumlah = 0
	fmt.Print("Masukkan teks: ")
	for {
		fmt.Scanf("%c", &karakter)
		if karakter == '.' || *jumlah >= NMAX {
			break
		}
		if karakter != ' ' && karakter != '\n' {
			(*t)[*jumlah] = karakter
			*jumlah++
		}
	}
}
func tampilkanArray(t tabel, jumlah int) {
	fmt.Print("Teks yang dimasukkan: ")
	for i := 0; i < jumlah; i++ {
		fmt.Printf("%c ", t[i])
	}
	fmt.Println()
}
func balikArray(t *tabel, jumlah int) {
	for i := 0; i < jumlah/2; i++ {
		t[i], t[jumlah-1-i] = t[jumlah-1-i], t[i]
	}
}
func cekPalindrom(t tabel, jumlah int) bool {
	var salinan tabel
	copy(salinan[:], t[:])
	balikArray(&salinan, jumlah)
	for i := 0; i < jumlah; i++ {
		if t[i] != salinan[i] {
			return false
		}
	}
	return true
}

func main() {
	var teks tabel
	var panjang int
	AnastasiaIsiTeks(&teks, &panjang)

	tampilkanArray(teks, panjang)
	balikArray(&teks, panjang)
	tampilkanArray(teks, panjang)

	fmt.Print("Apakah teks palindrom? ")
	if cekPalindrom(teks, panjang) {
		fmt.Println("Ya")
	} else {
		fmt.Println("Tidak")
	}
}
