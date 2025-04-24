//Feros Pedrosa

package main

import "fmt"

func main() {
	var klubA, klubB string
	var skorA, skorB int
	var hasil []string

	fmt.Print("Klub A: ")
	fmt.Scanln(&klubA)
	fmt.Print("Klub B: ")
	fmt.Scanln(&klubB)

	for {
		fmt.Printf("Pertandingan %d: ", len(hasil)+1)
		fmt.Scan(&skorA, &skorB)

		if skorA < 0 || skorB < 0 {
			break
		}

		if skorA > skorB {
			hasil = append(hasil, klubA)
		} else if skorB > skorA {
			hasil = append(hasil, klubB)
		} else {
			hasil = append(hasil, "Draw")
		}
	}

	fmt.Println("\nDaftar klub yang memenangkan pertandingan:")
	for i, h := range hasil {
		fmt.Printf("Hasil %d: %s\n", i+1, h)
	}
	fmt.Println("Pertandingan selesai")
}
