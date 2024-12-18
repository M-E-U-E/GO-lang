package main
import "fmt"
func main() {
	var t int;
	fmt.Scan(&t)
    for j := 0; j < t; j++ {
        var n int
		fmt.Scan(&n)
		fmt.Printf("%d\n",n-1)
    }
}