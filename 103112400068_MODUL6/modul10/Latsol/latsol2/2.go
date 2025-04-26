package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Print("Masukkan jumlah elemen: ")
	fmt.Scan(&n)

	arr := make([]int, n)

	fmt.Println("Masukkan nilai-nilai array:")
	for i := 0; i < n; i++ {
		fmt.Scan(&arr[i])
	}

	fmt.Println("\nSeluruh isi array:")
	for i := 0; i < n; i++ {
		fmt.Printf("%d ", arr[i])
	}

	fmt.Println("\n\nElemen dengan indeks ganjil:")
	for i := 1; i < n; i += 2 {
		fmt.Printf("%d ", arr[i])
	}
	fmt.Println("\n\nElemen dengan indeks genap:")
	for i := 0; i < n; i += 2 {
		fmt.Printf("%d ", arr[i])
	}

	sum := 0
	for i := 0; i < n; i++ {
		sum += arr[i]
	}
	rata := float64(sum) / float64(n)
	fmt.Printf("\n\nRata-rata: %.2f\n", rata)

	var sumSquaredDiff float64
	for i := 0; i < n; i++ {
		diff := float64(arr[i]) - rata
		sumSquaredDiff += diff * diff
	}
	stdDev := math.Sqrt(sumSquaredDiff / float64(n))
	fmt.Printf("Standar deviasi: %.2f\n", stdDev)
}
