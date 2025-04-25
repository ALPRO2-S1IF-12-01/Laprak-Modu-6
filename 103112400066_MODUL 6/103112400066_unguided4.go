// DWI OKTA SURYANINGRUM
// 103112400066

package main

import "fmt"

const NMAX int = 127 // Konstanta untuk ukuran maksimum array

type tabel [NMAX]rune // Deklarasi tipe array rune dengan ukuran NMAX

// membuat fungsi untuk mengisi array dengan input dari user
func isiArray(t *tabel, n *int) {
    var input rune
    *n = 0 // Inisialisasi panjang array
    fmt.Print("Teks : ")
    
    // Loop untuk membaca input hingga menemukan titik atau mencapai kapasitas maksimum
    for {
        fmt.Scanf("%c", &input) // membaca input per karakter

        // melewati spasi dan newline
        if input == ' ' || input == '\n' {
            continue
        }

        // berhenti jika menemukan titik atau melebihi kapasitas array
        if input == '.' || *n >= NMAX {
            break
        }

        t[*n] = input // menyimpan karakter ke array
        (*n)++        // menambah counter panjang array
    }
}

// membuat fungsi untuk mencetak isi array ke layar
func cetakArray(t tabel, n int) {
    // mencetak setiap karakter dipisahkan spasi
    for i := 0; i < n; i++ {
        fmt.Printf("%c ", t[i])
    }
    fmt.Println() // membuat baris baru setelah selesai
}

// membuat fungsi untuk membalikkan urutan elemen dalam array
func balikanArray(t *tabel, n int) {
    // membalik array dengan menukar elemen dari depan dan belakang
    for i := 0; i < n/2; i++ {
        temp := t[i]          // menyimoan sementara elemen depan
        t[i] = t[n-1-i]       // mengganti elemen depan dengan belakang
        t[n-1-i] = temp       // mengganti elemen belakang dengan depan
    }
}

// membuat fungsi untuk memeriksa apakah array adalah palindrom
func palindrom(t tabel, n int) bool {
    // membandingkan elemen dari kedua ujung
    for i := 0; i < n/2; i++ {
        if t[i] != t[n-1-i] { // jika ada yang tidak sama
            return false       // bukan palindrom
        }
    }
    return true // jika semua sama, palindrom
}

func main() {
    var tab tabel // array untuk menyimpan karakter
    var m int     // variabel untuk menyimpan panjang array

    isiArray(&tab, &m) // memanggil fungsi untuk mengisi array
    tabCopy := tab // membuat salinan array untuk pengecekan palindrom

    // membalik array dan mencetak hasilnya
    fmt.Print("Reverse teks : ")
    balikanArray(&tab, m)
    cetakArray(tab, m)

    // mencetak dan menampilkan apakah palindrom
    fmt.Print("Palindrom ? ")
    if palindrom(tabCopy, m) {
        fmt.Println("true")
    } else {
        fmt.Println("false")
    }
}