package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const NMAX int = 127

type tabel [NMAX]rune

func isiArray(t *tabel, n *int) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Teks: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	*n = len(input)
	for i := 0; i < *n && i < NMAX; i++ {
		(*t)[i] = rune(input[i])
	}
}

func cetakArray(t tabel, n int) {
	for i := 0; i < n; i++ {
		fmt.Printf("%c", t[i])
	}
	fmt.Println()
}

func balikArray(t *tabel, n int) {
	for i := 0; i < n/2; i++ {
		t[i], t[n-1-i] = t[n-1-i], t[i]
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
	var n int

	isiArray(&tab, &n)

	fmt.Print("Reverse teks: ")
	balikArray(&tab, n)
	cetakArray(tab, n)

	if palindrom(tab, n) {
		fmt.Println("Teks adalah palindrom.")
	} else {
		fmt.Println("Teks bukan palindrom.")
	}
}
