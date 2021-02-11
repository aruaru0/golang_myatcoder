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

func f() {
	N := getI()
	N++
	a := make([]int, N)
	b := make([]int, N)
	for i := 1; i < N; i++ {
		x := getI()
		a[i] = x + a[i-1]
		b[i] = x ^ b[i-1]
	}
	cnt := 0
	for i := 1; i < N; i++ {
		for j := i; j < N; j++ {
			sum := a[j] - a[i-1]
			xor := b[j] ^ b[i-1]
			if sum == xor {
				// out(i, j, a[j], a[i-1], b[j], b[i-1])
				cnt++
			}
		}
	}
	out(cnt)
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N := getI()
	N++
	a := make([]int, N)
	for i := 1; i < N; i++ {
		a[i] = getI()
	}
	bit := make([][]int, 20)
	for j := 0; j < 20; j++ {
		bit[j] = make([]int, N)
	}
	for i := 0; i < N; i++ {
		for j := 0; j < 20; j++ {
			bit[j][i] = (a[i] >> j) % 2
		}
	}
	for i := 0; i < 20; i++ {
		for j := 1; j < N; j++ {
			bit[i][j] += bit[i][j-1]
		}
	}

	r := 0
	cnt := 0
	for l := 1; l < N; l++ {
		for r < N {
			ok := true
			for k := 0; k < 20; k++ {
				if bit[k][r]-bit[k][l-1] > 1 {
					ok = false
					break
				}
			}
			if !ok {
				break
			}
			r++
		}
		cnt += r - l
	}
	out(cnt)
}
