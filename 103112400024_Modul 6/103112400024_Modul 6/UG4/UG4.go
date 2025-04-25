package main
//Setyo Nugroho
//103112400024

import "fmt"

const NMAX int = 127

type tabel [NMAX]rune

func isiArray(t *tabel, n *int) {
	var Teks rune
	*n = 0
	fmt.Print("Teks: ")
	for {
		fmt.Scanf("%c", &Teks)
		if Teks == '.' || *n > NMAX {
			break
		}
		if Teks != ' ' && Teks != '\n' {
			(*t)[*n] = Teks
			*n++
		}
	}
}

func cetakArray(t tabel, n int) {
	fmt.Print("Reverse teks: ")
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
	var m int
	isiArray(&tab, &m)
	balikanArray(&tab, m)
	cetakArray(tab, m)
	fmt.Print("Palindrom: ", palindrom(tab, m))
}