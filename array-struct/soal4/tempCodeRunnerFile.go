return true
}

func main() {
	var teks tabel
	var jumlah int
	isiArray(&teks, &jumlah)
	cetakArray(teks, jumlah)

	if palindrom(teks, jumlah) {
		fmt.Println("Palindrom? true")
	} else {
		fmt.Println("Palindrom? false")
	}
}
