package main
//Mohammad Reyhan Aretha Fatin 
//103112400078
import "fmt"
func tampilkanHasil(pemenang []string) {
	fmt.Println("\n=== Hasil Pertandingan ===")
	for i, hasil := range pemenang {
		fmt.Printf("Hasil %d: %s\n", i+1, hasil)
	}
	fmt.Println("Pertandingan selesai.")
}

func main() {
	var clubA, clubB string
	var skorA, skorB int
	var hasilPertandingan []string
	nomorPertandingan := 1

	fmt.Print("Klub A: ")
	fmt.Scan(&clubA)
	fmt.Print("Klub B: ")
	fmt.Scan(&clubB)

	for {
		fmt.Printf("Pertandingan ke-%d (format: skorA skorB): ", nomorPertandingan)
		fmt.Scan(&skorA, &skorB)

		if skorA < 0 || skorB < 0 {
			break
		}

		if skorA > skorB {
			hasilPertandingan = append(hasilPertandingan, clubA)
		} else if skorB > skorA {
			hasilPertandingan = append(hasilPertandingan, clubB)
		} else {
			hasilPertandingan = append(hasilPertandingan, "Draw")
		}
		nomorPertandingan++
	}
	tampilkanHasil(hasilPertandingan)
}
