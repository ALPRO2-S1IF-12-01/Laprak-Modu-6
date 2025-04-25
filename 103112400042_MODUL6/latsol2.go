// M.HANIF AL FAIZ
// 1031124000042
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
		fmt.Printf("Elemen ke-%d: ", i+1)
		fmt.Scan(&arr[i])
	}

	for {
		fmt.Println("\nMenu Operasi Array:")
		fmt.Println("a. Tampilkan seluruh isi array")
		fmt.Println("b. Tampilkan elemen dengan indeks ganjil")
		fmt.Println("c. Tampilkan elemen dengan indeks genap")
		fmt.Println("d. Tampilkan elemen dengan indeks kelipatan x")
		fmt.Println("e. Hapus elemen pada indeks tertentu")
		fmt.Println("f. Hitung rata-rata array")
		fmt.Println("g. Hitung standar deviasi array")
		fmt.Println("h. Hitung frekuensi bilangan tertentu")
		fmt.Println("x. Keluar")

		var choice string
		fmt.Print("Pilihan Anda: ")
		fmt.Scan(&choice)

		switch choice {
		case "a":
			displayAll(arr)
		case "b":
			displayOddIndex(arr)
		case "c":
			displayEvenIndex(arr)
		case "d":
			displayMultiples(arr)
		case "e":
			arr = deleteElement(arr)
		case "f":
			calculateAverage(arr)
		case "g":
			calculateStdDev(arr)
		case "h":
			calculateFrequency(arr)
		case "x":
			return
		default:
			fmt.Println("Pilihan tidak valid!")
		}
	}
}

func displayAll(arr []int) {
	fmt.Println("\nIsi array:")
	for i, val := range arr {
		fmt.Printf("Indeks %d: %d\n", i, val)
	}
}

func displayOddIndex(arr []int) {
	fmt.Println("\nElemen dengan indeks ganjil:")
	for i := 1; i < len(arr); i += 2 {
		fmt.Printf("Indeks %d: %d\n", i, arr[i])
	}
}

func displayEvenIndex(arr []int) {
	fmt.Println("\nElemen dengan indeks genap:")
	for i := 0; i < len(arr); i += 2 {
		fmt.Printf("Indeks %d: %d\n", i, arr[i])
	}
}
func displayMultiples(arr []int) {
	var x int
	fmt.Print("Masukkan nilai x: ")
	fmt.Scan(&x)

	fmt.Printf("\nElemen dengan indeks kelipatan %d:\n", x)
	for i := 0; i < len(arr); i++ {
		if i%x == 0 {
			fmt.Printf("Indeks %d: %d\n", i, arr[i])
		}
	}
}

func deleteElement(arr []int) []int {
	var index int
	fmt.Print("Masukkan indeks yang akan dihapus: ")
	fmt.Scan(&index)

	newArr := append(arr[:index], arr[index+1:]...)

	fmt.Println("\nArray setelah penghapusan:")
	for i, val := range newArr {
		fmt.Printf("Indeks %d: %d\n", i, val)
	}

	return newArr
}

func calculateAverage(arr []int) {
	if len(arr) == 0 {
		fmt.Println("Array kosong")
		return
	}

	sum := 0
	for _, val := range arr {
		sum += val
	}
	average := float64(sum) / float64(len(arr))
	fmt.Printf("\nRata-rata: %.2f\n", average)
}

func calculateStdDev(arr []int) {
	if len(arr) == 0 {
		fmt.Println("Array kosong")
		return
	}

	sum := 0
	for _, val := range arr {
		sum += val
	}
	mean := float64(sum) / float64(len(arr))

	variance := 0.0
	for _, val := range arr {
		variance += math.Pow(float64(val)-mean, 2)
	}
	variance /= float64(len(arr))

	stdDev := math.Sqrt(variance)
	fmt.Printf("\nStandar deviasi: %.2f\n", stdDev)
}
func calculateFrequency(arr []int) {
	var target int
	fmt.Print("Masukkan bilangan yang akan dihitung frekuensinya: ")
	fmt.Scan(&target)

	count := 0
	for _, val := range arr {
		if val == target {
			count++
		}
	}
	fmt.Printf("\nFrekuensi bilangan %d: %d\n", target, count)
}
