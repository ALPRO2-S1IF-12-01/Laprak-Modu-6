//Ahmad Ruba'i
//103112400074 
package main

import "fmt"

func main() {
	var total int
	fmt.Print("Masukkan jumlah elemen array: ")
	fmt.Scanln(&total)

	koleksiData := make([]int, total)
	for i := 0; i < total; i++ {
		fmt.Printf("Data ke-%d: ", i)
		fmt.Scanln(&koleksiData[i])
	}

	fmt.Println("\na. Seluruh isi array:", koleksiData)

	fmt.Print("b. Elemen dengan indeks ganjil:")
	for i := 1; i < total; i += 2 {
		fmt.Print(" ", koleksiData[i])
	}
	fmt.Println()

	fmt.Print("c. Elemen dengan indeks genap:")
	for i := 0; i < total; i += 2 {
		fmt.Print(" ", koleksiData[i])
	}
	fmt.Println()

	var x int
	fmt.Print("d. Masukkan bilangan x untuk kelipatan indeks: ")
	fmt.Scanln(&x)
	fmt.Printf("d. Elemen dengan indeks kelipatan %d:", x)
	if x > 0 {
		for i := x; i < total; i += x {
			fmt.Print(" ", koleksiData[i])
		}
		fmt.Println()
	} else {
		fmt.Println()
	}

	var index int
	fmt.Print("Masukkan indeks yang ingin dihapus: ")
	fmt.Scanln(&index)
	if index >= 0 && index < total {
		koleksiData = append(koleksiData[:index], koleksiData[index+1:]...)
		total--
		fmt.Print("e. Array setelah penghapusan:", koleksiData)
		fmt.Println()
	} else {
		fmt.Println("Indeks tidak valid.")
	}

	if total > 0 {
		jumlah := 0
		for _, v := range koleksiData {
			jumlah += v
		}
		rata := float64(jumlah) / float64(total)
		fmt.Printf("f. Rata-rata: %.2f\n", rata)

		var totalSelisihKuadrat float64
		for _, v := range koleksiData {
			selisih := float64(v) - rata
			totalSelisihKuadrat += selisih * selisih
		}
		varian := totalSelisihKuadrat / float64(total)

		sd := varian
		for i := 0; i < 10; i++ {
			sd = 0.5 * (sd + varian/sd)
		}
		fmt.Printf("g. Simpangan baku kuadrat: %.2f\n", sd)

		var cari int
		fmt.Print("h. Masukkan angka yang ingin dihitung frekuensinya: ")
		fmt.Scanln(&cari)
		frekuensi := 0
		for _, v := range koleksiData {
			if v == cari {
				frekuensi++
			}
		}
		fmt.Printf("h. Frekuensi kemunculan angka %d: %d kali\n", cari, frekuensi)
	}
}