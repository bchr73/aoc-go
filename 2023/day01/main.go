package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func check(err error) {
    if err != nil {
        log.Fatal(err)
    }
}

func toInt(b byte) *int {
    i, _ := strconv.Atoi(string(b))
    return &i
}

func main(){
    bytes_chan := make(chan byte, 1)

    go func() {
        f, err := os.Open("input")
        check(err)
        defer f.Close()
        for {
            b1 := make([]byte, 1)
            n1, err := f.Read(b1)
            check(err)
            if n1 < 1 {
                return
            }
            bytes_chan <- b1[0]
        }
    }()

    sum := 0
    var n1, n2 *int

    for b := range bytes_chan {
        if b == '\n' {
            if n2 == nil {
                i := *n1
                n2 = &i
            }
            fmt.Printf("%d,%d\n", *n1, *n2)
            cali := (*n1 * 10) + *n2
            fmt.Printf("%d + %d\n", sum, cali)
            sum += cali
            n1, n2 = nil, nil
            continue
        }
        if (int(b) < 48 || int(b) > 57) {
            continue
        }
        if n1 == nil {
            n1 = toInt(b)
            continue
        } 
        n2 = toInt(b)
    }
    fmt.Printf("sum: %d\n", sum)
}
