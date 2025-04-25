// MUHAMMAD GAMEL AL GHIFARI
// 103112400028
package main

import (
	"fmt"
)

const NMAX int = 127

type tabel [NMAX]rune

func isiArray(t *tabel, n *int) {
	var input rune
	*n = 0
	fmt.Println("Masukkan karakter (akhiri dengan titik '.'):")

	for {
		fmt.Scanf("%c", &input)

		if input == '\n' {
			continue
		}
		if input == '.' || *n >= NMAX {
			break
		}
		t[*n] = input
		(*n)++
	}
}

func cetakArray(t tabel, n int) {
	fmt.Print("Teks: ")
	for i := 0; i < n; i++ {
		fmt.Printf("%c", t[i])
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
	var n int

	isiArray(&tab, &n)

	fmt.Println("\nTeks asli:")
	cetakArray(tab, n)

	balikanArray(&tab, n)
	fmt.Println("\nTeks setelah dibalik:")
	cetakArray(tab, n)

	if isPalindrom(tab, n) {
		fmt.Println("\nPalindrom? true")
	} else {
		fmt.Println("\nPalindrom? false")
	}
}
