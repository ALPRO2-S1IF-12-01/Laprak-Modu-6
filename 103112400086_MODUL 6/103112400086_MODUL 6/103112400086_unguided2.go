package main

import "fmt"

func main() {
	var n int
	fmt.Print("masukkan jumlah elemen : ")
	fmt.Scan(&n)

	data := make([]int, n)
	var total int

	for i := 0; i < n; i++ {
		fmt.Printf("masukkan elemen ke-%d : ", i+1)
		fmt.Scan(&data[i])
		total += data[i]
	}

	rataRata := float64(total) / float64(n)
	fmt.Printf("\nRata-rata: %.2f\n", rataRata)

	var cariAngka, frekuensi int
	fmt.Print("Masukkan angka yang ingin dicari frekuensinya : ")
	fmt.Scan(&cariAngka)

	for i := 0; i < n; i++ {
		if data[i] == cariAngka {
			frekuensi++
		}
	}

	fmt.Println("Frekuensi angka", cariAngka, ":", frekuensi)
}

// CHILA
