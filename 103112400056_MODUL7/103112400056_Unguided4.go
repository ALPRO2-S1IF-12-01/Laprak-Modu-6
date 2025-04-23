// Felix Pedrosa V 

package main

import (
	"fmt"
)

const NMAX int = 127

type table [NMAX]rune

// Fungsi untuk mengisi array dengan karakter
func fillArray(t *table, count *int) {
	var character rune
	*count = 0
	fmt.Print("Masukkan teks: ")
	for {
		fmt.Scanf("%c", &character)

		// Hentikan jika karakter adalah titik
		if character == '.' || *count >= NMAX {
			break
		}

		// Abaikan spasi dan newline
		if character != ' ' && character != '\n' {
			(*t)[*count] = character
			*count++
		}
	}
}

// Fungsi untuk mencetak array
func printArray(t table, count int) {
	for i := 0; i < count; i++ {
		fmt.Printf("%c ", t[i])
	}
	fmt.Println()
}

// Fungsi untuk membalikkan array
func reverseArray(t *table, count int) {
	for i := 0; i < count/2; i++ {
		t[i], t[count-1-i] = t[count-1-i], t[i]
	}
}

// Fungsi untuk memeriksa apakah array adalah palindrom
func isPalindrome(t table, count int) bool {
	var copyTable table
	copy(copyTable[:], t[:])
	reverseArray(&copyTable, count)
	for i := 0; i < count; i++ {
		if t[i] != copyTable[i] {
			return false
		}
	}
	return true
}

func main() {
	var charTable table
	var length int

	fillArray(&charTable, &length)

	fmt.Print("Teks: ")
	printArray(charTable, length)

	fmt.Printf("Apakah palindrom? %v\n", isPalindrome(charTable, length))
}