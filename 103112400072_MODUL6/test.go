package main

import "fmt"

func main() {
	var angka = [...]int{5, 4, 3, 2, 1}
	fmt.Println(angka)
	fmt.Println(len(angka))

	var matriks = [2][3]int{{1, 2, 3},
		{4, 5, 6}}

	fmt.Println(matriks)

	var check = [3][4]int{{4, 1, 2, 4},
		{4, 6, 7, 8},
		{0, 1, 1, 4}}
	fmt.Print(check)

	var hewan = make([]string, 2)
	hewan [0] = "raja"
	hewan [1] = "anjing"
	fmt.Print(hewan)
}
 