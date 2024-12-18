package main

import "fmt"

func main() {
	var t int
	fmt.Scan(&t)
	for i := 0; i<t; i++ {
		var n int
		fmt.Scan(&n)
		var ans int
		ans += n%10
		ans += n/10
		fmt.Println(ans)
	}
}