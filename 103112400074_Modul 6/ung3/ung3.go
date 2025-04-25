//Ahmad Ruba'i
//103112400074 
package main

import "fmt"

func main() {
	var klubA, klubB string

	fmt.Print("Klub A : ")
	fmt.Scan(&klubA)
	fmt.Print("Klub B : ")
	fmt.Scan(&klubB)

	var hasilSemua []string
	nomorPertandingan := 1

	for {
		var skorA, skorB int
		
		fmt.Printf("Pertandingan %d : ", nomorPertandingan)
		fmt.Scan(&skorA, &skorB)

		if skorA < 0 || skorB < 0 {
			break
		}

		var pemenang string
		if skorA > skorB {
			pemenang = klubA
		} else if skorB > skorA {
			pemenang = klubB
		} else {
			pemenang = "Draw"
		}

		hasilSemua = append(hasilSemua, pemenang)


		nomorPertandingan++
	}

	for i, hasil := range hasilSemua {
		fmt.Printf("Hasil %d : %s\n", i+1, hasil)
	}
	fmt.Println("Pertandingan selesai")
}