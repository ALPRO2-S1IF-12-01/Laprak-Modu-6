package main

import "fmt"

func isPalindrom(text []rune) bool {
	for i := 0; i < len(text)/2; i++ {
		if text[i] != text[len(text)-1-i] {
			return false
		}
	}
	return true
}

func main() {
	var input string
	fmt.Print("Masukkan teks (akhiri dengan titik): ")
	fmt.Scan(&input)

	text := []rune{}
	for _, ch := range input {
		if ch == '.' {
			break
		}
		text = append(text, ch)
	}

	fmt.Print("Teks dibalik: ")
	for i := len(text) - 1; i >= 0; i-- {
		fmt.Printf("%c ", text[i])
	}
	fmt.Println()

	fmt.Println("Palindrom?", isPalindrom(text))
}
