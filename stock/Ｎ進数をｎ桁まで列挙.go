package main

import "fmt"

// N進数でn桁を列挙
func nNum(s string, n, N int) {
	//	fmt.Println(len(s))
	if len(s) >= n {
		fmt.Println(s)
		return
	}
	for i := 0; i < N; i++ {
		nNum(s+string('0'+i), n, N)
	}
}

func main() {
	nNum("", 4, 3)
}
