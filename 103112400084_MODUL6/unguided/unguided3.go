package main
// Nufail Alauddin Tsaqif
// 103112400084
import "fmt"
func tampilkanHasil(hasil []string) {
	fmt.Println("\n=== Hasil Pertandingan ===")
	for i, pemenang := range hasil {
		fmt.Printf("Pertandingan %d dimenangkan oleh: %s\n", i+1, pemenang)
	}
	fmt.Println("Pertandingan selesai.")
}

func main() {
	var klub1, klub2 string
	var skor1, skor2 int
	var hasil []string
	nomor := 1

	fmt.Print("Masukkan nama Klub 1: ")
	fmt.Scan(&klub1)

	fmt.Print("Masukkan nama Klub 2: ")
	fmt.Scan(&klub2)

	for {
		fmt.Printf("Pertandingan ke-%d (format input: skor1 skor2, negatif untuk berhenti): ", nomor)
		fmt.Scan(&skor1, &skor2)

		if skor1 < 0 || skor2 < 0 {
			break
		}

		switch {
		case skor1 > skor2:
			hasil = append(hasil, klub1)
		case skor2 > skor1:
			hasil = append(hasil, klub2)
		default:
			hasil = append(hasil, "Draw")
		}

		nomor++
	}

	tampilkanHasil(hasil)
}
