package main

import "fmt"

const NMAX int = 127
type tabel [NMAX]rune

func isiArray(t *tabel, n *int) {
	var huruf rune
	*n = 0
	fmt.Print("Teks : ")
	for {
		fmt.Scanf("%c", &huruf)
		if huruf == '.' || *n >= NMAX {
			break
		}
		if huruf != '\n' && huruf != '\r' {
			t[*n] = huruf
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
	var salin tabel

	for i := 0; i < n; i++ {
		salin[i] = t[i]
	}
	
	balikanArray(&salin, n)
	
	for i := 0; i < n; i++ {
		if t[i] != salin[i] {
			return false
		}
	}
	return true
}

func main() {
	var tab tabel
	var m int

	isiArray(&tab, &m)
	fmt.Print("Reverse teks : ")
	var salin tabel
	
	for i := 0; i < m; i++ {
		salin[i] = tab[i]
	}
	balikanArray(&salin, m)
	cetakArray(salin, m)

	fmt.Print("Palindrom ? ")
	if palindrom(tab, m) {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}
