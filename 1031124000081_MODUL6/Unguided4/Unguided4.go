//RYAN AKEYLA NOVIANTO WIDODO
//103112400081

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const NMAX int = 127

func bacaInput() ([]rune, error) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Masukkan karakter (gunakan titik untuk berhenti):")
	input, err := reader.ReadString('.')
	if err != nil {
		return nil, fmt.Errorf("kesalahan membaca input: %w", err)
	}
	input = strings.TrimSuffix(input, ".")
	input = strings.TrimSpace(input)
	return []rune(input), nil
}

func balikRune(s []rune) {
	for i, j := 0, len(s)-1; 1 < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func isPalindrom(s []rune) bool {
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}

func main() {
	runeSlice, err := bacaInput()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	if len(runeSlice) > NMAX {
		fmt.Printf("Input terlalu panjang (maksimal %d karakter).\n", NMAX)
		return
	}

	fmt.Print("Teks Awal: ")
	fmt.Println(string(runeSlice))

	balikRune(runeSlice)
	fmt.Print("Teks Terbalik: ")
	fmt.Println(string(runeSlice))

	if isPalindrom(runeSlice) {
		fmt.Println("Palindrom? true")
	} else {
		fmt.Println("Palindrom? false")
	}
}
