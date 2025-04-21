// Raihan Adi Arba
// 103112400071

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
		fmt.Println("\n==========================================")
		fmt.Println("	SELAMAT DATANG DI MENU ARRAY	")
		fmt.Println("==========================================")
		fmt.Println("1. TAMPILKAN SEMUA ARRAY")
		fmt.Println("2. TAMPILKAN SEMUA ARRAY GENAP")
		fmt.Println("3. TAMPILKAN SEMUA ARRAY GANJIL")
		fmt.Println("4. TAMPILKAN ARRAY DENGAN INDEKS KELIPATAN X")
		fmt.Println("5. HAPUS ELEMEN PADA INDEKS TERTENTU")
		fmt.Println("6. HITUNG RATA-RATA ARRAY")
		fmt.Println("7. HITUNG STANDAR DEVIASI ARRAY")
		fmt.Println("8. HITUNG FREKUENSI BILANGAN TERTENTU")
		fmt.Println("9. KELUAR")
		fmt.Print("Pilihan Anda: ")

		var choice int
		fmt.Scan(&choice)

		if choice == 1 {
			displayAll(arr)
		} else if choice == 2 {
			displayEvenIndex(arr)
		} else if choice == 3 {
			displayOddIndex(arr)
		} else if choice == 4 {
			displayMultiples(arr)
		} else if choice == 5 {
			arr = deleteElement(arr)
		} else if choice == 6 {
			calculateAverage(arr)
		} else if choice == 7 {
			calculateStdDev(arr)
		} else if choice == 8 {
			calculateFrequency(arr)
		} else if choice == 9 {
			fmt.Println("Keluar dari program.")
			return
		} else {
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

func displayEvenIndex(arr []int) {
	fmt.Println("\nElemen dengan indeks genap:")
	for i := 0; i < len(arr); i += 2 {
		fmt.Printf("Indeks %d: %d\n", i, arr[i])
	}
}

func displayOddIndex(arr []int) {
	fmt.Println("\nElemen dengan indeks ganjil:")
	for i := 1; i < len(arr); i += 2 {
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

	if index < 0 || index >= len(arr) {
		fmt.Println("Indeks tidak valid!")
		return arr
	}

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
