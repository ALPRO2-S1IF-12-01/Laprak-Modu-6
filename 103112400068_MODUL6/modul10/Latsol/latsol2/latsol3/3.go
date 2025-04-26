package main

import "fmt"

func main() {
	var klubA, klubB string

	fmt.Print("Klub A : ")
	fmt.Scan(&klubA)
	fmt.Print("Klub B : ")
	fmt.Scan(&klubB)

	var skorA, skorB int
	var pertandingan int = 0
	var menangA, menangB, seri int = 0, 0, 0

	for {
		pertandingan++
		fmt.Printf("Pertandingan %d : ", pertandingan)
		fmt.Scan(&skorA, &skorB)

		if skorA < 0 || skorB < 0 {
			pertandingan--
			break
		}

		if skorA > skorB {
			fmt.Printf("Hasil %d : %s\n", pertandingan, klubA)
			menangA++
		} else if skorB > skorA {
			fmt.Printf("Hasil %d : %s\n", pertandingan, klubB)
			menangB++
		} else {
			fmt.Printf("Hasil %d : Draw\n", pertandingan)
			seri++
		}
	}

	fmt.Println("Pertandingan selesai")
}
