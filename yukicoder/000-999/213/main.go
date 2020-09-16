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

var N, P, C int
var A []int = []int{2, 3, 5, 7, 11, 13}
var B []int = []int{4, 6, 8, 9, 10, 12}

const mod = int(1e9 + 7)

var dp [2][400][4000]int
var dp2 [8000]int
var single [8010]int

const MAT = 130

type mat struct {
	v [MAT][MAT]int
}

func mulmat(a, b mat, n int) mat {
	var r mat
	for x := 0; x < n; x++ {
		for y := 0; y < n; y++ {
			r.v[x][y] = 0
		}
	}
	for x := 0; x < n; x++ {
		for z := 0; z < n; z++ {
			for y := 0; y < n; y++ {
				r.v[x][y] += (a.v[x][z] * b.v[z][y]) % mod
			}
		}
	}
	for x := 0; x < n; x++ {
		for y := 0; y < n; y++ {
			r.v[x][y] %= mod
		}
	}
	return r
}

func powmat(p int, a mat, n int) mat {
	var r mat
	for x := 0; x < n; x++ {
		for y := 0; y < n; y++ {
			r.v[x][y] = 0
		}
	}
	for i := 0; i < n; i++ {
		r.v[i][i] = 1
	}
	for p != 0 {
		if p%2 == 1 {
			r = mulmat(r, a, n)
		}
		a = mulmat(a, a, n)
		p >>= 1
	}
	return r
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1000000)
	N, P, C = getInt(), getInt(), getInt()

	dp[0][0][0] = 1
	dp[1][0][0] = 1

	for x := 0; x < 6; x++ {
		for y := 0; y < P; y++ {
			for i := 0; i < y*13+1; i++ {
				if dp[0][y][i] != 0 {
					dp[0][y+1][i+A[x]] += dp[0][y][i]
					dp[0][y+1][i+A[x]] %= mod
				}
			}
		}
	}

	for x := 0; x < 6; x++ {
		for y := 0; y < C; y++ {
			for i := 0; i < y*12+1; i++ {
				if dp[1][y][i] != 0 {
					dp[1][y+1][i+B[x]] += dp[1][y][i]
					dp[1][y+1][i+B[x]] %= mod
				}
			}
		}
	}

	for x := 0; x < 651; x++ {
		for y := 0; y < 601; y++ {
			single[x+y] += dp[0][P][x] * dp[1][C][y]
			single[x+y] %= mod
		}
	}

	M := P*13 + C*12
	var mm mat
	for i := 0; i < M-1; i++ {
		mm.v[i][i+1] = 1
	}
	for i := 0; i <= M; i++ {
		mm.v[M-1][M-i] = single[i]
	}
	m2 := powmat(N+M-1, mm, M)

	tot := 0
	for i := 0; i < M; i++ {
		tot += m2.v[0][i]
	}
	out(tot % mod)
}
