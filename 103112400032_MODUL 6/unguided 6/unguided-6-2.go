// daffa tsaqifna f
package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Print("Masukkan jumlah elemen array: ")
	fmt.Scan(&n)

	arr := make([]int, n)

	fmt.Println("Masukkan elemen array:")
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	for {
		fmt.Println("\nMenu:")
		fmt.Println("a. Tampilkan seluruh isi array")
		fmt.Println("b. Tampilkan elemen dengan indeks ganjil")
		fmt.Println("c. Tampilkan elemen dengan indeks genap")
		fmt.Println("d. Tampilkan elemen dengan indeks kelipatan x")
		fmt.Println("e. Hapus elemen pada indeks tertentu")
		fmt.Println("f. Tampilkan rata-rata array")
		fmt.Println("g. Tampilkan standar deviasi array")
		fmt.Println("h. Tampilkan frekuensi suatu bilangan")
		fmt.Println("0. keluar")
		var opsi string
		fmt.Print("Pilih menu (a-h): ")
		fmt.Scan(&opsi)

		switch opsi {
		case "a":
			fmt.Println("Isi array:", arr)

		case "b":
			fmt.Print("Indeks ganjil: ")
			for i := 1; i < len(arr); i += 2 {
				fmt.Print(arr[i], " ")
			}
			fmt.Println()

		case "c":
			fmt.Print("Indeks genap: ")
			for i := 0; i < len(arr); i += 2 {
				fmt.Print(arr[i], " ")
			}
			fmt.Println()

		case "d":
			var x int
			fmt.Print("Masukkan nilai x: ")
			fmt.Scan(&x)
			fmt.Print("Indeks kelipatan ", x, ": ")
			for i := 0; i < len(arr); i++ {
				if i%x == 0 {
					fmt.Print(arr[i], " ")
				}
			}
			fmt.Println()

		case "e":
			var index int
			fmt.Print("Masukkan indeks yang ingin dihapus: ")
			fmt.Scan(&index)
			if index >= 0 && index < len(arr) {
				arr = append(arr[:index], arr[index+1:]...)
				fmt.Println("Array setelah penghapusan:", arr)
			} else {
				fmt.Println("Indeks tidak valid")
			}

		case "f":
			sum := 0
			for _, v := range arr {
				sum += v
			}
			rata := float64(sum) / float64(len(arr))
			fmt.Println("Rata-rata:", rata)

		case "g":
			sum := 0
			for _, v := range arr {
				sum += v
			}
			rata := float64(sum) / float64(len(arr))
			var std float64
			for _, v := range arr {
				std += math.Pow(float64(v)-rata, 2)
			}
			std = math.Sqrt(std / float64(len(arr)))
			fmt.Println("Standar deviasi:", std)

		case "h":
			var target int
			fmt.Print("Masukkan bilangan yang ingin dicari frekuensinya: ")
			fmt.Scan(&target)
			count := 0
			for _, v := range arr {
				if v == target {
					count++
				}
			}
			fmt.Println("Frekuensi:", count)
		case "0":
			return

		default:
			fmt.Println("Opsi tidak dikenali.")
		}
	}
}
