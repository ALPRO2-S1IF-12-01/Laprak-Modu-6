//ABID FADHILAH M
//103112400046
package main

import "fmt"

const NMAX int = 127

type tabel [NMAX]rune

func isiArray(t *tabel, n *int) {
	var karakter rune
	*n = 0
	fmt.Print("Teks: ")
	for {
		fmt.Scanf("%c", &karakter)
		if karakter == '.' || *n >= NMAX {
			break
		}
		if karakter != ' ' && karakter != '\n' && karakter != '\r' {
			t[*n] = karakter
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

func palindrom(t tabel, n int) bool {
	var salin tabel
	copy(salin[:], t[:]) 
	balikanArray(&salin, n)
	for i := 0; i < n; i++ {
		if t[i] != salin[i] {
			return false
		}
	}
	return true
}

func main() {
	var teks tabel
	var jumlah int
	isiArray(&teks, &jumlah)
	cetakArray(teks, jumlah)

	if palindrom(teks, jumlah) {
		fmt.Println("Palindrom? true")
	} else {
		fmt.Println("Palindrom? false")
	}
}
