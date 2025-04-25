//JESIKA METANIA RAHMA ARIFIN
// 103112400080

package main

import "fmt"

func main() {
	var klubA, klubB string
	fmt.Print("Klub A: ")
	fmt.Scan(&klubA)
	fmt.Print("Klub B: ")
	fmt.Scan(&klubB)

	const jumlahPertandingan = 9

	for i := 1; i <= jumlahPertandingan; i++ {
		var skorA, skorB int
		fmt.Printf("Pertandingan %d: ", i)
		fmt.Scan(&skorA, &skorB)

		if skorA < 0 || skorB < 0 {
			fmt.Printf("Pertandingan %d: Skor tidak valid, dilewati.\n", i)
			continue
		}
		if skorA > skorB {
			fmt.Println(klubA)
		} else if skorB > skorA {
			fmt.Println(klubB)
		} else {
			fmt.Println("Draw")
		}
	}

	fmt.Println("Pertandingan selesai")
}
