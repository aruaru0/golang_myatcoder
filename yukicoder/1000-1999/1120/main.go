package main

import (
	"bufio"
	"fmt"
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

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1000000)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N := getI()
	a := getInts(N)
	b := getInts(N)

	if N == 2 {
		if a[0]+a[1] != b[0]+b[1] {
			out(-1)
			return
		}
		out(abs(a[0] - b[0]))
		return
	}

	totA, totB := 0, 0
	for i := 0; i < N; i++ {
		totA += a[i]
		totB += b[i]
	}
	if totA < totB || (totA-totB)%(N-2) != 0 {
		out(-1)
		return
	}

	ans := (totA - totB) / (N - 2)
	cnt := 0
	p := make([]int, N)
	for i := 0; i < N; i++ {
		p[i] = b[i] - (a[i] - ans)
	}
	for i := 0; i < N; i++ {
		if p[i]%2 == 1 || p[i] < 0 {
			out(-1)
			return
		}
		cnt += p[i]
	}
	if cnt != 2*ans {
		out(-1)
		return
	}
	out(ans)
}
