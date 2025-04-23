package main

import (
	"fmt"
)

const NMAX int = 127

type tabel [NMAX]rune

func isiArray(t *tabel, n *int) {
	var ch rune
	*n = 0
	fmt.Print("Teks : ")
	for {
		fmt.Scanf("%c", &ch)

		// Hentikan kalau titik
		if ch == '.' || *n >= NMAX {
			break
		}

		// Abaikan spasi dan newline
		if ch != ' ' && ch != '\n' {
			(*t)[*n] = ch
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
	var tCopy tabel
	copy(tCopy[:], t[:])
	balikanArray(&tCopy, n)
	for i := 0; i < n; i++ {
		if t[i] != tCopy[i] {
			return false
		}
	}
	return true
}

func main() {
	var tab tabel
	var m int

	isiArray(&tab, &m)

	fmt.Print("Teks : ")
	cetakArray(tab, m)

	fmt.Printf("Palindrom ? %v\n", palindrom(tab, m))
}
