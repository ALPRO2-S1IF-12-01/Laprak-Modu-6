package main

import (
	"fmt"
	"math"
)

func main() {
	var ( 
		storage [10]int 
		temporaryStorage []int
		n, x, indexToDelete int
	)

	fmt.Print("Masukkan data ke dalam storage: ")
	n = len(storage)
	for i := 0; i < n; i++ {
		fmt.Scan(&storage[i])
	}

	// a. Menampilkan keseluruhan isi dari array
	fmt.Print("Berikut merupakan isi dari keseluruhan storage (array): ")
	fmt.Println(storage)

	// b. Menampilkan elemen-elemen array dengan indeks ganjil saja
	fmt.Print("Berikut merupakan isi dari storage (array) dengan indeks ganjil: ")
	for i := 0; i < n; i++ {
		if i % 2 == 1 {
			temporaryStorage = append(temporaryStorage, storage[i])
		}
	}
	fmt.Println(temporaryStorage)
	temporaryStorage = nil

	// c. Menampilkan elemen-elemen array dengan indeks genap saja (asumsi indek ke-0 adalah genap)
	fmt.Print("Berikut merupakan isi dari storage (array) dengan indeks genap: ")
	for i := 0; i < n; i++ {
		if i % 2 == 0 {
			temporaryStorage = append(temporaryStorage, storage[i])
		}
	}
	fmt.Println(temporaryStorage)
	temporaryStorage = nil

	// d. Menampilkan elemen-elemen array dengan indeks kelipatan bilangan x. x bisa diperoleh dari masukan pengguna
	fmt.Print("Berikut merupakan isi dari storage (array) dengan indeks kelipatan x: ")
	fmt.Scan(&x)
	for i, num := range storage {
		if i % x == 0 && i != 0 {
			temporaryStorage = append(temporaryStorage, storage[num])
		}
	}
	fmt.Println(temporaryStorage)
	
	// e. Menghapus elemen array pada indeks tertentu, asumsi indeks yang hapus selalu valid. Tampilkan keseluruhan isi dari arraynya, pastikan data yang dihapus tidak tampil
	temporaryStorage = storage[:]
	fmt.Print("Masukkan index yang akan dihapus: ")
	fmt.Scan(&indexToDelete)
	temporaryStorage = append(temporaryStorage[:indexToDelete], temporaryStorage[indexToDelete+1:]...)
	fmt.Println(temporaryStorage)
	temporaryStorage = nil

	// f. Menampilkan rata-rata dari bilangan yang ada di dalam array
	var rataRata, totalNum float64
	for _, num := range storage {
		totalNum += float64(storage[num])
	}
	rataRata = totalNum / float64(n)
	fmt.Print("Rata-rata dari bilangan dalam storage: ")
	fmt.Println(rataRata)

	// g. Menampilkan standar deviasi atau simpangan baku dari bilangan yang ada di dalam array tersebut
	var mean, sum, diff, squaredDiff, variance float64
	sum = 0.0
	for _, num := range storage {
		sum += float64(num)
	}
	mean = sum / float64(len(storage))
	squaredDiff = 0.0
	for _, num := range storage {
		diff = float64(num) - mean
		squaredDiff += diff * diff
	}
	variance = math.Sqrt(squaredDiff / float64(len(storage)-1))
	fmt.Println("Standar deviasi dari array storage:", variance)

	// h. Menampilkan frekuensi dari suatu bilangan tertentu di dalam array yang telah diisi tersebut
	fmt.Println("\n========== Frekuensi ==========")
	frequency := make(map[int]int)
	for _, num := range storage {
		frequency[num]++
	}
	for num, freq := range frequency {
		fmt.Printf("Frekuensi dari bilangan %d: %d\n", num, freq)
	}
	fmt.Println("========== ========= ==========")
}