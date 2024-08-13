package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)
var wr = bufio.NewWriter(os.Stdout)

func out(x ...interface{}) {
	fmt.Fprintln(wr, x...)
}

func getI() int {
	sc.Scan()
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return i
}

func getF() float64 {
	sc.Scan()
	i, e := strconv.ParseFloat(sc.Text(), 64)
	if e != nil {
		panic(e)
	}
	return i
}

func getInts(N int) []int {
	ret := make([]int, N)
	for i := 0; i < N; i++ {
		ret[i] = getI()
	}
	return ret
}

func getS() string {
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

// min for n entry
func nmin(a ...int) int {
	ret := a[0]
	for _, e := range a {
		ret = min(ret, e)
	}
	return ret
}

// max for n entry
func nmax(a ...int) int {
	ret := a[0]
	for _, e := range a {
		ret = max(ret, e)
	}
	return ret
}

func chmin(a *int, b int) bool {
	if *a < b {
		return false
	}
	*a = b
	return true
}

func chmax(a *int, b int) bool {
	if *a > b {
		return false
	}
	*a = b
	return true
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

func get_sum(Lx, Rx, Ly, Ry, Lz, Rz int) int {
	total := S[Rx][Ry][Rz]
	total -= S[Lx-1][Ry][Rz]
	total -= S[Rx][Ly-1][Rz]
	total -= S[Rx][Ry][Lz-1]
	total += S[Lx-1][Ly-1][Rz]
	total += S[Lx-1][Ry][Lz-1]
	total += S[Rx][Ly-1][Lz-1]
	total -= S[Lx-1][Ly-1][Lz-1]
	return total
}

var S [][][]int

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N := getI()

	a := make([][][]int, N)
	for i := 0; i < N; i++ {
		a[i] = make([][]int, N)
		for j := 0; j < N; j++ {
			a[i][j] = getInts(N)
		}
	}

	S = make([][][]int, N+1)
	for i := 0; i <= N; i++ {
		S[i] = make([][]int, N+1)
		for j := 0; j <= N; j++ {
			S[i][j] = make([]int, N+1)
		}
	}

	for x := 0; x < N; x++ {
		for y := 0; y < N; y++ {
			for z := 0; z < N; z++ {
				S[x+1][y+1][z+1] += S[x+1][y+1][z] + a[x][y][z]
			}
		}
	}

	for x := 0; x <= N; x++ {
		for y := 0; y < N; y++ {
			for z := 0; z <= N; z++ {
				S[x][y+1][z] += S[x][y][z]
			}
		}
	}

	for x := 0; x < N; x++ {
		for y := 0; y <= N; y++ {
			for z := 0; z <= N; z++ {
				S[x+1][y][z] += S[x][y][z]
			}
		}
	}

	Q := getI()
	for qi := 0; qi < Q; qi++ {
		Lx, Rx, Ly, Ry, Lz, Rz := getI(), getI(), getI(), getI(), getI(), getI()
		ans := get_sum(Lx, Rx, Ly, Ry, Lz, Rz)
		out(ans)
	}

}
