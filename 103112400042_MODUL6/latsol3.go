// M.HANIF AL FAIZ
// 103112400042
package main

import "fmt"

func main() {

	var clubA, clubB string
	fmt.Print("Klub A : ")
	fmt.Scan(&clubA)
	fmt.Print("Klub B : ")
	fmt.Scan(&clubB)

	var winners []string

	for i := 1; ; i++ {
		var scoreA, scoreB int
		fmt.Printf("Pertandingan %d : ", i)
		_, err := fmt.Scan(&scoreA, &scoreB)

		if err != nil || scoreA < 0 || scoreB < 0 {
			fmt.Println("Pertandingan selesai")
			break
		}

		var result string
		switch {
		case scoreA > scoreB:
			result = clubA
		case scoreB > scoreA:
			result = clubB
		default:
			result = "Draw"
		}

		winners = append(winners, result)
		fmt.Printf("Hasil %d : %s\n", i, result)
	}

	fmt.Println("\nDaftar klub yang memenangkan pertandingan:")
	for i, winner := range winners {
		if winner != "Draw" {
			fmt.Printf("Pertandingan %d: %s\n", i+1, winner)
		}
	}
}
