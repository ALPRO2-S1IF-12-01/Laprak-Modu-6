// Felix Pedrosa V 

package main

import (
	"fmt"
)

func main() {
	var teamA, teamB string
	var scoreA, scoreB int
	var results []string

	// Input nama klub
	fmt.Print("Masukkan nama Klub A: ")
	fmt.Scanln(&teamA)
	fmt.Print("Masukkan nama Klub B: ")
	fmt.Scanln(&teamB)

	matchNumber := 1
	for {
		fmt.Printf("Skor Pertandingan %d: ", matchNumber)
		fmt.Scan(&scoreA, &scoreB)

		// Menghentikan input jika skor negatif
		if scoreA < 0 || scoreB < 0 {
			break
		}

		// Menentukan hasil pertandingan
		if scoreA > scoreB {
			results = append(results, teamA)
		} else if scoreB > scoreA {
			results = append(results, teamB)
		} else {
			results = append(results, "Draw")
		}
		matchNumber++
	}

	// Menampilkan hasil pertandingan
	for i, result := range results {
		fmt.Printf("Hasil Pertandingan %d: %s\n", i+1, result)
	}
	fmt.Println("Pertandingan selesai")
}