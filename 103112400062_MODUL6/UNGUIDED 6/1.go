// M. DAVI ILYAS RENALDO_103112400062
package main

import "fmt"

type Titik struct {x,y int}
type Lingkaran struct {Pusat Titik; Radius int}

func main() {
    var l1, l2 Lingkaran
    var p Titik

    fmt.Scan(&l1.Pusat.x, &l1.Pusat.y, &l1.Radius)
    
    fmt.Scan(&l2.Pusat.x, &l2.Pusat.y, &l2.Radius)
    
    fmt.Scan(&p.x, &p.y)

    cekLingkaran := func(p Titik, l Lingkaran) bool {
        dx := p.x - l.Pusat.x  
        dy := p.y - l.Pusat.y  
        return dx*dx + dy*dy <= l.Radius*l.Radius
    }

    inL1 := cekLingkaran(p, l1)  
    inL2 := cekLingkaran(p, l2) 

    switch {
    case inL1 && inL2:
        fmt.Println("titik didalam lingkaran 1 dan 2")
    case inL1:
        fmt.Println("titik didalam lingkaran 1")
    case inL2:
        fmt.Println("titik didalam lingkaran 2")
    default:
        fmt.Println("titik diluar lingkaran 1 dan 2")
    }
}