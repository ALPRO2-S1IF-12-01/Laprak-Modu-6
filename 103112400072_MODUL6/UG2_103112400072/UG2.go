package main

import (
	"fmt"
	"math"
)

func main() {

	var n int
	fmt.Print("Masukkan jumlah elemen array (N): ")
	fmt.Scan(&n)

	arr := make([]int, n)
	fmt.Println("Masukkan elemen-elemen array:")
	for i := 0; i < n; i++ {
		fmt.Printf("Data ke-%d: ", i)
		fmt.Scan(&arr[i])
	}

	originalArr := make([]int, len(arr))
	copy(originalArr, arr)

	fmt.Println("\n===== HASIL SEMUA OPERASI =====")

	fmt.Println("\na. Semua elemen array:")
	fmt.Println("Isi array:", originalArr)

	fmt.Println("\nb. Elemen dengan indeks ganjil:")
	for i := 1; i < len(originalArr); i += 2 {
		fmt.Printf("Index %d: %d\n", i, originalArr[i])
	}

	fmt.Println("\nc. Elemen dengan indeks genap:")
	for i := 0; i < len(originalArr); i += 2 {
		fmt.Printf("Index %d: %d\n", i, originalArr[i])
	}

	x := 3
	fmt.Printf("\nd. Elemen dengan indeks kelipatan %d:\n", x)
	for i := 0; i < len(originalArr); i++ {
		if i%x == 0 {
			fmt.Printf("Index %d: %d\n", i, originalArr[i])
		}
	}

	indexToDelete := 1
	fmt.Printf("\ne. Hapus elemen pada indeks %d:\n", indexToDelete)
	fmt.Println("Array sebelum dihapus:", originalArr)

	arrAfterDelete := make([]int, len(originalArr))
	copy(arrAfterDelete, originalArr)

	if indexToDelete >= 0 && indexToDelete < len(arrAfterDelete) {
		arrAfterDelete = append(arrAfterDelete[:indexToDelete], arrAfterDelete[indexToDelete+1:]...)
		fmt.Println("Array setelah dihapus:", arrAfterDelete)
	} else {
		fmt.Println("Indeks tidak valid.")
	}

	sum := 0
	for _, v := range originalArr {
		sum += v
	}
	rata := float64(sum) / float64(len(originalArr))
	fmt.Printf("\nf. Rata-rata elemen: %.2f\n", rata)

	mean := 0.0
	for _, v := range originalArr {
		mean += float64(v)
	}
	mean /= float64(len(originalArr))

	variance := 0.0
	for _, v := range originalArr {
		variance += math.Pow(float64(v)-mean, 2)
	}
	stddev := math.Sqrt(variance / float64(len(originalArr)))
	fmt.Printf("\ng. Standar deviasi elemen: %.2f\n", stddev)

	fmt.Println("\nh. Frekuensi setiap bilangan dalam array:")

	freqMap := make(map[int]int)
	for _, v := range originalArr {
		freqMap[v]++
	}

	for num, count := range freqMap {
		fmt.Printf("Bilangan %d muncul %d kali\n", num, count)
	}

	fmt.Println("\nProgram selesai.")
}
