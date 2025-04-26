package main

import "fmt"

const NMAX int = 127

type tabel [NMAX]rune

func isiArray(t *tabel, n *int) {

	var c rune
	*n = 0
	fmt.Print("Teks: ")
	for {
		fmt.Scanf("%c", &c)
		if c == '.' || *n >= NMAX {
			break
		}
		t[*n] = c
		*n++
	}
}

func cetakArray(t tabel, n int) {
	for i := 0; i < n; i++ {
		fmt.Printf("%c", t[i])
	}
	fmt.Println()
}

func balikArray(t *tabel, n int) {
	var temp tabel
	for i := 0; i < n; i++ {
		temp[i] = t[n-1-i]
	}
	for i := 0; i < n; i++ {
		t[i] = temp[i]
	}
}

func palindrom(t tabel, n int) bool {

	var temp tabel

	for i := 0; i < n; i++ {
		temp[i] = t[i]
	}

	balikArray(&temp, n)

	for i := 0; i < n; i++ {
		if t[i] != temp[i] {
			return false
		}
	}
	return true
}

func main() {
	var tab tabel
	var m int
	isiArray(&tab, &m)

	fmt.Print("Teks: ")
	cetakArray(tab, m)

	balikArray(&tab, m)
	fmt.Print("Reverse teks: ")
	cetakArray(tab, m)

	balikArray(&tab, m)

	fmt.Print("Palindrom: ")
	if palindrom(tab, m) {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}
