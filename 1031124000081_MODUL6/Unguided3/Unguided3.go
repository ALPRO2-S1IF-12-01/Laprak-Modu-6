//RYAN AKEYLA NOVIANTO WIDODO
//103112400081

package main

import "fmt"

type pertandingan struct {
	klubA    string
	klubB    string
	skorA    int
	skorB    int
	pemenang string
}

func main() {
	var klubA, klubB string
	var hasil []pertandingan

	fmt.Print("klub A:")
	fmt.Scanln(&klubA)
	fmt.Print("klub B:")
	fmt.Scanln(&klubB)

	for {
		var pertandingan pertandingan
		pertandingan.klubA = klubA
		pertandingan.klubB = klubB
		fmt.Printf("Pertandingan %d: ", len(hasil)+1)
		fmt.Scan(&pertandingan.skorA, &pertandingan.skorB)

		if pertandingan.skorA < 0 || pertandingan.skorB < 0 {
			break
		}

		if pertandingan.skorA > pertandingan.skorB {
			pertandingan.pemenang = klubA
		} else if pertandingan.skorB > pertandingan.skorA {
			pertandingan.pemenang = klubB
		} else {
			pertandingan.pemenang = "Draw"
		}
		hasil = append(hasil, pertandingan)
	}

	fmt.Println("\nDaftar klub yang memenangkan pertandingan:")
	for i, h := range hasil {
		fmt.Printf("Hasil %d: %s (%d - %d)\n", i+1, h.pemenang, h.skorA, h.skorB)
	}
	fmt.Println("Pertandingan selesai")
}
