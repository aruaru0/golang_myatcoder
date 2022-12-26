package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

func getInts(N int) []int {
	ret := make([]int, N)
	for i := 0; i < N; i++ {
		ret[i] = getInt()
	}
	return ret
}

func getString() string {
	sc.Scan()
	return sc.Text()
}

// min, max, asub, absなど基本関数
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

func abs(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}

func lowerBound(a []int, x int) int {
	idx := sort.Search(len(a), func(i int) bool {
		return a[i] >= x
	})
	return idx
}

func upperBound(a []int, x int) int {
	idx := sort.Search(len(a), func(i int) bool {
		return a[i] > x
	})
	return idx
}

const mod = int(1e9 + 7)

func solve1(N, K int, a []int) {
	K = min(1000000, K)
	sum := 0
	for i := 0; i < N; i++ {
		sum += a[i]
	}
	f := sum
	s := sum
	for i := N; i < K; i++ {
		f = sum
		s += sum
		s %= mod
		// out(f, s)
		x := a[i%N]
		a[i%N] = sum
		sum += sum - x
		sum %= mod
		if sum < 0 {
			sum += mod
		}
	}
	out(f, s)
}

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

func solve2(N, K int, x []int) {
	a := make([][]int, N)
	for i := 0; i < N; i++ {
		a[i] = make([]int, N)
	}
	for i := 0; i < N; i++ {
		a[0][i] = 1
	}
	for i := 1; i < N; i++ {
		a[i][i-1] = 1
	}
	a = powModMatrix(a, K-N)
	f := 0
	for i := 0; i < N; i++ {
		f += x[N-1-i] * a[0][i]
		f %= mod
	}

	a = make([][]int, N+1)
	for i := 0; i < N+1; i++ {
		a[i] = make([]int, N+1)
	}
	a[0][0] = 2
	a[0][N] = -1
	for i := 1; i < N+1; i++ {
		a[i][i-1] = 1
	}
	// for i := 0; i < N+1; i++ {
	// 	out(a[i])
	// }
	a = powModMatrix(a, K-N)
	y := make([]int, N+1)
	y[0] = 0
	for i := 1; i < N+1; i++ {
		y[i] = y[i-1] + x[i-1]
	}
	// out(y)
	s := 0
	for i := 0; i < N+1; i++ {
		s += y[N-i] * a[0][i]
		s %= mod
		if s < 0 {
			s += mod
		}
	}

	out(f, s)
}

func main() {
	sc.Split(bufio.ScanWords)
	N, K := getInt(), getInt()
	a := getInts(N)
	if K > 1000000 {
		solve2(N, K, a)
	} else {
		solve1(N, K, a)
	}
}
