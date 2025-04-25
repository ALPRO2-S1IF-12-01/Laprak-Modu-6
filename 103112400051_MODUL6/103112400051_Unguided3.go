package main

import "fmt"

func main() {
    var klub [2]string
    var pemenang [100]string
    var skorA, skorB int
    var i int

    fmt.Print("Klub A : ")
    fmt.Scan(&klub[0])

    fmt.Print("Klub B : ")
    fmt.Scan(&klub[1])

    for {
        fmt.Printf("Pertandingan %d : ", i+1)
        fmt.Scan(&skorA, &skorB)

        if skorA < 0 || skorB < 0 {
            break
        }

        if skorA > skorB {
            pemenang[i] = klub[0]
        } else if skorB > skorA {
            pemenang[i] = klub[1]
        } else {
            pemenang[i] = "Draw"
        }
        i++
    }

    fmt.Println()
    for j := 0; j < i; j++ {
        fmt.Printf("Hasil %d : %s\n", j+1, pemenang[j])
    }
    fmt.Println("Pertandingan selesai")
}
