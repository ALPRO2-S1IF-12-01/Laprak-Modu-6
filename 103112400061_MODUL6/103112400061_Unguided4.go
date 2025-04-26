package main

import "fmt"

const NMAX int = 127
type tabel [NMAX]rune

func main() {
	var (
		tabel tabel
		n int
		isPalindrom bool
	)

	isiArray(&tabel, &n)

	isPalindrom = palindrom(tabel, n)

	balikanArray(&tabel, n)

	fmt.Printf("%-10s : ", "Reverse teks")
	cetakArray(tabel, n)

	fmt.Printf("%-10s %t\n", "Palindrom ?", isPalindrom)
}

func isiArray(t *tabel, n *int) {
	var char rune
	
	fmt.Printf("%-12s : ", "Teks")
	for {
		fmt.Scanf("%c", &char)

		if char == ' ' {
			continue
		}

		if char == '.' {
			break
		}

		if *n < NMAX {
			(*t)[*n] = char
			*n++
		} else {
			break
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
		(*t)[i], (*t)[n-1-i] = (*t)[n-1-i], (*t)[i]
	}
}

func palindrom(t tabel, n int) bool {
	var tempTabel tabel

	for i := 0; i < n; i++ {
		tempTabel[i] = t[i]
	}
	balikanArray(&tempTabel, n)
	for i := 0; i < n; i++ {
		if t[i] != tempTabel[i] {
			return false
		}
	}
	return true
}