package main

import (
	"fmt"
	"sort"
)

//
// バグ修正（copyを忘れていた)
// 問題なければこちらの組み合わせ列挙を使う
//
func generateComb(index []int, s, r int, ch chan []int) {
	if r != 0 {
		if s < 0 {
			return
		}
		generateComb(index, s-1, r, ch)
		index[r-1] = s
		generateComb(index, s-1, r-1, ch)
	} else {
		out := make([]int, len(index))
		copy(out, index)
		ch <- out
	}
	return
}

func foreachComb(n, k int, ch chan []int) {
	index := make([]int, k)
	generateComb(index, n-1, k, ch)
	close(ch)
}

func main() {
	a := []byte{'a', 'b', 'c', 'd', 'e'}

	// goを使う
	comb := make(chan []int, 1)
	go foreachComb(5, 3, comb)

	for ch := range comb {
		for _, v := range ch {
			fmt.Printf("%c ", a[v])
		}
		fmt.Println()
	}

	b := []int{4, 7, 89, 23, 6, 7, 4}
	sort.Ints(b)
	fmt.Println(b)
}
