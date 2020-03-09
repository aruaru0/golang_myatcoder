package main

import "fmt"

func generateComb(index []int, s, r int, ch chan []int) {
	if r != 0 {
		if s < 0 {
			return
		}
		generateComb(index, s-1, r, ch)
		index[r-1] = s
		generateComb(index, s-1, r-1, ch)
	} else {
		ch <- index
	}
	return
}

func Combinations(n, k int, ch chan []int) {
	index := make([]int, k)
	generateComb(index, n-1, k, ch)
	close(ch)
}

func main() {
	a := []byte{'a', 'b', 'c', 'd', 'e'}

	// goを使う
	comb := make(chan []int)
	go Combinations(5, 3, comb)

	for ch := range comb {
		for _, v := range ch {
			fmt.Printf("%c ", a[v])
		}
		fmt.Println()
	}
}
