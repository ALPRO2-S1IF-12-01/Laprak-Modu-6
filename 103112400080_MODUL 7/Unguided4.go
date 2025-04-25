// JESIKA METANIA RAHMA ARIFIN
// 103112400080

package main

import (
	"fmt"
	"strings"
)

// Fungsi untuk mengecek apakah kata adalah palindrom
func isPalindrom(kata string) bool {
	kata = strings.ToUpper(kata)
	n := len(kata)
	for i := 0; i < n/2; i++ {
		if kata[i] != kata[n-1-i] {
			return false
		}
	}
	return true
}

// Fungsi untuk mencetak karakter per huruf dipisah spasi
func tampilkanTeks(kata string) string {
	chars := []rune(strings.ToUpper(kata))
	var hasil strings.Builder
	for i, c := range chars {
		hasil.WriteRune(c)
		if i != len(chars)-1 {
			hasil.WriteString(" ")
		}
	}
	return hasil.String()
}

func main() {
	var kata string
	fmt.Scan(&kata)

	fmt.Println(kata)
	fmt.Println(isPalindrom(kata))
}
