package main

// Muhammad Faris Rachmadi
// 103112400079

import (
	"fmt"
)

func main() {
	var klubA, klubB string
	var skorA, skorB int
	var pertandingan int
	var hasil []string

	fmt.Print("Masukkan nama klub A: ")
	fmt.Scan(&klubA)
	fmt.Print("Masukkan nama klub B: ")
	fmt.Scan(&klubB)

	fmt.Println("\nPertandingan dimulai!")

	pertandingan = 1
	for {
		fmt.Printf("Pertandingan %d: ", pertandingan)
		fmt.Scan(&skorA, &skorB)

		if skorA < 0 || skorB < 0 {
			fmt.Println("Pertandingan selesai")
			break
		}

		if skorA > skorB {
			hasil = append(hasil, klubA)
		} else if skorB > skorA {
			hasil = append(hasil, klubB)
		} else {
			hasil = append(hasil, "Draw")
		}
		pertandingan++
	}

	fmt.Println("\nHasil pertandingan:")
	for i := 0; i < len(hasil); i++ {
		fmt.Printf("Hasil %d: %s\n", i+1, hasil[i])
	}
}
