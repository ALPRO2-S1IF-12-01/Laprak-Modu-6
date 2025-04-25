//Ahmad Ruba'i
//103112400074 
package main

import "fmt"

const NMAX int = 127

type tabel [NMAX]rune

func isiArray(t *tabel, n *int) {
	var char rune
	*n = 0
	fmt.Print("Karakter\t:")
	fmt.Scanf("%c", &char)
	for char != '.' && *n < NMAX {
		if char != ' ' {  
			t[*n] = char
			*n++
		}
		fmt.Scanf("%c", &char)
	}
}

func cetakArray(t tabel, n int) {
	fmt.Print("Reverse teks\t:")
	for i := 0; i < n; i++ {
		fmt.Printf("%c", t[i])
	}
	fmt.Println()
}

func balikUrutanArray(t *tabel, n int) {
	var temp tabel
	for i := 0; i < n; i++ {
		temp[i] = t[n-1-i]
	}
	for i := 0; i < n; i++ {
		t[i] = temp[i]
	}
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
	var tab tabel
	var m int

	isiArray(&tab, &m)
	fmt.Print("Teks\t\t:")
	for i := 0; i < m; i++ {
		fmt.Printf("%c", tab[i])
	}
	fmt.Println()
	balikUrutanArray(&tab, m)
	cetakArray(tab, m)
	balikUrutanArray(&tab, m)
	fmt.Print("Palindrom\t?")
	if palindrom(tab, m) {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}