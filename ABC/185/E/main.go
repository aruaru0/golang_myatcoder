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

var memo [1100][1100]int

func rec(i, j int) int {
	if i == N && j == M {
		return 0
	}
	if i == N && j != M {
		return M - j
	}
	if j == M && i != N {
		return N - i
	}
	if memo[i][j] != -1 {
		return memo[i][j]
	}
	ret := int(1e10)
	if a[i] == b[j] {
		ret = rec(i+1, j+1)
	} else {
		ret = rec(i+1, j+1) + 1
	}

	ret = min(ret, rec(i+1, j)+1)
	ret = min(ret, rec(i, j+1)+1)
	ret = min(ret, rec(i+1, j+1)+2)

	// ri := N - 1 - i
	// rj := M - 1 - j
	// if ri > rj {
	// ret = min(ret, rec(i+1, j)+1)
	// }
	// if ri < rj {
	// ret = min(ret, rec(i, j+1)+1)
	// }
	memo[i][j] = ret
	return ret
}

var N, M int
var a, b []int

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1000000)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N, M = getI(), getI()
	a = getInts(N)
	b = getInts(M)
	for i := 0; i < 1100; i++ {
		for j := 0; j < 1100; j++ {
			memo[i][j] = -1
		}
	}
	out(rec(0, 0))
}
