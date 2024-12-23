package main

import "fmt"

func main() {

    var a [5]int
    fmt.Println("emp:", a)
	for i := 0; i<5; i++ {
		fmt.Scan(&a[i])
	}
	fmt.Println("New:", a)
    a[4] = 100
    // fmt.Println("set:", a)
    fmt.Println("get:", a[4])

    fmt.Println("len:", len(a))

    b := [5]int{1, 2, 3, 4, 5}
    fmt.Println("dcl:", b)

    b = [...]int{11, 22, 33, 44, 55}
    fmt.Println("dcl:", b)

    b = [...]int{100, 2: 400, 500, 0}
    fmt.Println("idx:", b)

    var twoD [2][3]int
    for i := 0; i < 2; i++ {
        for j := 0; j < 3; j++ {
            twoD[i][j] = i + 1 + j
        }
    }
    fmt.Println("2d: ", twoD)

    twoD = [2][3]int{
        {1, 2, 3},
        {1, 2, 3},
    }
    fmt.Println("2d: ", twoD)
}