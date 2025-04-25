// MUHAMMAD GAMEL AL GHIFARI
// // 103112400028
package main

import "fmt"

func main() {
	var clubA, clubB string
	fmt.Print("Masukkan nama club A: ")
	fmt.Scan(&clubA)
	fmt.Print("Masukkan nama club B: ")
	fmt.Scan(&clubB)

	var winners []string

	for i := 1; ; i++ {
		var scoreA, scoreB int
		fmt.Printf("Pertandingan %d (masukkan skor A dan B): ", i)
		fmt.Scan(&scoreA, &scoreB)
		if scoreA < 0 || scoreB < 0 {
			fmt.Println("Pertandingan selesai")
			break
		}
		if scoreA > scoreB {
			winners = append(winners, clubA)
		} else if scoreB > scoreA {
			winners = append(winners, clubB)
		} else {
			winners = append(winners, "Draw")
		}
	}

	fmt.Println("\nHasil pertandingan:")
	for i, winner := range winners {
		fmt.Printf("Hasil %d: %s\n", i+1, winner)
	}
}
