// daffa tsaqifna f
package main

import "fmt"

const Nmax int = 127

type table [Nmax]rune

func isiArray(t *table, n *int) {
	var c rune
	*n = 0
	fmt.Print("Teks  : ")
	for {
		fmt.Scanf("%c", &c)
		if c == '.' || *n >= Nmax {
			break
		}
		t[*n] = c
		*n++
	}
}

func cetakArray(t table, n int) {
	for i := 0; i < n; i++ {
		fmt.Printf("%c ", t[i])
	}
	fmt.Println()
}

func balikanArray(t *table, n int) {
	for i := 0; i < n/2; i++ {
		t[i], t[n-1-i] = t[n-1-i], t[i]
	}
}

func palindrom(t table, n int) bool {
	for i := 0; i < n/2; i++ {
		if t[i] != t[n-1-i] {
			return false
		}
	}
	return true
}

func main() {
	var tab table
	var m int

	isiArray(&tab, &m)

	fmt.Print("Reverse teks : ")
	var reversedTab table = tab // salin isi
	balikanArray(&reversedTab, m)
	cetakArray(reversedTab, m)

	fmt.Println("Palindrom ?", palindrom(tab, m))
}
