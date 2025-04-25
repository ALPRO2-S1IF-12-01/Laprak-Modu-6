package main

//103112400049 Hisyam Nurdiatmoko

import "fmt"

const NMAX int = 127

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

func cetakArray(t tabel, n int) {
	fmt.Print("Reverse teks	: ")
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
	var tabClone tabel
	copy(tabClone[:], t[:])
	balikanArray(&tabClone, n)
	for i := 0; i < n; i++ {
		if t[i] != tabClone[i] {
			return false
		}
	}
	return true
}

func main() {
	var tab tabel
	var hisyam int
	isiArray(&tab, &hisyam)
	fmt.Print("Teks		: ")
	for i := 0; i < hisyam; i++ {
		fmt.Printf("%c ", tab[i])
	}
	fmt.Println()
	balikanArray(&tab, hisyam)
	cetakArray(tab, hisyam)
	fmt.Print("Palindrom	? ")
	if palindrom(tab, hisyam) {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}
