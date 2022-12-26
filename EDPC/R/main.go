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

const mod = 1000000007

func mulMod(A, B [][]int) [][]int {
	H := len(A)
	W := len(B[0])
	K := len(A[0])
	C := make([][]int, W)
	for i := 0; i < W; i++ {
		C[i] = make([]int, W)
	}

	for i := 0; i < H; i++ {
		for j := 0; j < W; j++ {
			for k := 0; k < K; k++ {
				C[i][j] += A[i][k] * B[k][j]
				C[i][j] %= mod
			}
		}
	}

	return C
}

func powModMatrix(A [][]int, p int) [][]int {
	N := len(A)
	ret := make([][]int, N)
	for i := 0; i < N; i++ {
		ret[i] = make([]int, N)
		ret[i][i] = 1
	}

	for p > 0 {
		if p&1 == 1 {
			ret = mulMod(ret, A)
		}
		A = mulMod(A, A)
		p >>= 1
	}

	return ret
}

// 思いつかなかったので写経
func main() {
	sc.Split(bufio.ScanWords)

	N, K := getInt(), getInt()
	A := make([][]int, N)
	for i := 0; i < N; i++ {
		A[i] = make([]int, N)
		for j := 0; j < N; j++ {
			A[i][j] = getInt()
		}
	}

	ret := powModMatrix(A, K)

	ans := 0
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			ans += ret[i][j]
			ans %= mod
		}
	}

	out(ans)
}
