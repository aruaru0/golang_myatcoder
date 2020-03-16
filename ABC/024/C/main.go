package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func out(x ...interface{}) {
	fmt.Println(x...)
}

var sc = bufio.NewScanner(os.Stdin)

func getInt() int {
	sc.Scan()
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return i
}

func getString() string {
	sc.Scan()
	return sc.Text()
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func asub(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

func main() {
	sc.Split(bufio.ScanWords)

	_, D, K := getInt(), getInt(), getInt()
	L := make([]int, D)
	R := make([]int, D)
	for i := 0; i < D; i++ {
		L[i], R[i] = getInt(), getInt()
	}
	S := make([]int, K)
	T := make([]int, K)
	for i := 0; i < K; i++ {
		S[i], T[i] = getInt(), getInt()
	}

	arrived := make([]int, K)
	for d := 0; d < D; d++ {
		l := L[d]
		r := R[d]
		for k := 0; k < K; k++ {
			if arrived[k] != 0 {
				continue
			}
			if l <= S[k] && S[k] <= r {
				if l <= T[k] && T[k] <= r {
					S[k] = T[k]
					arrived[k] = d + 1
				} else if T[k] < S[k] {
					S[k] = l
				} else {
					S[k] = r
				}
			}
		}
	}
	for i := 0; i < K; i++ {
		out(arrived[i])
	}

}
