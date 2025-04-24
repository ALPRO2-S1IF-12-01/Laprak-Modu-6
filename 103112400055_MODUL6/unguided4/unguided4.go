//Feros Pedrosa

package main

import "fmt"

const NMAX int = 127

type tabel [NMAX]rune

func isiArray(t *tabel, n *int) {
	fmt.Println("Masukkan karakter (gunakan titik untuk berhenti):")
	for i := 0; i < NMAX; i++ {
		var ch rune
		fmt.Scanf("%c", &ch)
		if ch == '.' {
			break
		}
		t[i] = ch
		*n++
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
		t[i], t[n-i-1] = t[n-i-1], t[i]
	}
}

func palindrom(t tabel, n int) bool {
	for i := 0; i < n/2; i++ {
		if t[i] != t[n-i-1] {
			return false
		}
	}
	return true
}

func main() {
	var tab tabel
	var m int
	isiArray(&tab, &m)
	balikanArray(&tab, m)
	fmt.Print("Teks: ")
	cetakArray(tab, m)

	if palindrom(tab, m) {
		fmt.Println("Palindrom? true")
	} else {
		fmt.Println("Palindrom? false")
	}
}
