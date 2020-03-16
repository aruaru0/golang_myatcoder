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
	SL := make([]int, K)
	SR := make([]int, K)
	T := make([]int, K)
	for i := 0; i < K; i++ {
		S[i], T[i] = getInt(), getInt()
		SL[i] = S[i]
		SR[i] = S[i]
	}

	arrived := make([]int, K)
	for d := 0; d < D; d++ {
		l := L[d]
		r := R[d]
		for k := 0; k < K; k++ {
			if arrived[k] != 0 {
				continue
			}
			//fmt.Print(k, ":", SL[k], "-", SR[k], "-->")
			if l <= SL[k] && SL[k] <= r {
				SL[k] = l
			}
			if l <= SR[k] && SR[k] <= r {
				SR[k] = r
			}
			//fmt.Print(SL[k], "-", SR[k], "\n")
			if SL[k] <= T[k] && T[k] <= SR[k] {
				arrived[k] = d + 1
			}
		}
	}
	for i := 0; i < K; i++ {
		out(arrived[i])
	}

}
