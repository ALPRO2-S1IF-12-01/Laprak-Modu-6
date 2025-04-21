// Raihan Adi Arba
// 103112400071

package main

import "fmt"

const NMAX int = 127

type tabel [NMAX]rune

func isiArray(t *tabel, n *int) {

	var input rune
	*n = 0
	fmt.Println("Masukkan karakter [akhirki dengan . ]:")

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

	fmt.Print("Isi array: ")
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
	var m int

	isiArray(&tab, &m)

	fmt.Println("\nArray sebelum dibalik:")
	cetakArray(tab, m)

	if isPalindrom(tab, m) {
		fmt.Println("Array ini palindrom")
	} else {
		fmt.Println("Array bukan palindrom")
	}

	balikanArray(&tab, m)

	fmt.Println("\nArray setelah dibalik:")
	cetakArray(tab, m)

	if isPalindrom(tab, m) {
		fmt.Println("Array ini palindrom")
	} else {
		fmt.Println("Array bukan palindrom")
	}
}
