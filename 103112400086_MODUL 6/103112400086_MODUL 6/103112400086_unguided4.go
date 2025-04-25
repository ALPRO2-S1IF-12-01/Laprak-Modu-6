package main

import "fmt"

const NMAX = 127

type tabel [NMAX]rune

func isiArray(t *tabel, n *int) {
	var Teks rune
	*n = 0
	fmt.Print("Teks		: ")
	for {
		fmt.Scanf("%c", &Teks)
		if Teks == '.' || *n >= NMAX {
			break
		}
		if Teks != ' ' && Teks != '\n' {
			(*t)[*n] = Teks
			*n++
		}
	}
}

func balikArray(t tabel, n int) tabel {
	var hasil tabel
	for i := 0; i < n; i++ {
		hasil[i] = t[n-1-i]
	}
	return hasil
}

func palindrom(t tabel, n int) bool {
	for i := 0; i < n/2; i++ {
		if t[i] != t[n-1-i] {
			return false
		}
	}
	return true
}

func main() {
	var tab, reversed tabel
	var n int

	isiArray(&tab, &n)

	fmt.Print("Teks		: ")
	for i := 0; i < n; i++ {
		fmt.Printf("%c ", tab[i])
	}
	fmt.Println()

	reversed = balikArray(tab, n)

	fmt.Print("Reverse teks	: ")
	for i := 0; i < n; i++ {
		fmt.Printf("%c ", reversed[i])
	}
	fmt.Println()

	if palindrom(tab, n) {
		fmt.Println("Palindrom	? true")
	} else {
		fmt.Println("Palindrom	? false")
	}
}

// CHILA
