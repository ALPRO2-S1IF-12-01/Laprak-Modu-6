package main

import (
	"fmt"
)

type Klub struct {
	nama string
}

type Pertandingan struct {
	klubA Klub
	klubB Klub
	skorA int
	skorB int
	pemenang string
}

func main() {
	var namaA, namaB string
	fmt.Print("Klub A: ")
	fmt.Scan(&namaA)
	fmt.Print("Klub B: ")
	fmt.Scan(&namaB)

	klubA := Klub{nama: namaA}
	klubB := Klub{nama: namaB}

	var daftarPertandingan []Pertandingan

	for {
		var skorA, skorB int
		fmt.Printf("Masukkan skor %s dan %s (format: A B): ", klubA.nama, klubB.nama)
		fmt.Scan(&skorA, &skorB)

		if skorA < 0 || skorB < 0 {
			break
		}

		p := Pertandingan{
			klubA: klubA,
			klubB: klubB,
			skorA: skorA,
			skorB: skorB,
		}

		if skorA > skorB {
			p.pemenang = klubA.nama
			fmt.Printf("Hasil: %s\n", klubA.nama)
		} else if skorB > skorA {
			p.pemenang = klubB.nama
			fmt.Printf("Hasil: %s\n", klubB.nama)
		} else {
			p.pemenang = "Draw"
			fmt.Println("Hasil: Draw")
		}

		daftarPertandingan = append(daftarPertandingan, p)
	}

	fmt.Println("Pertandingan selesai")
	fmt.Println("Daftar klub yang memenangkan pertandingan:")
	for i, p := range daftarPertandingan {
		fmt.Printf("Hasil %d : %s\n", i+1, p.pemenang)
	}
}
