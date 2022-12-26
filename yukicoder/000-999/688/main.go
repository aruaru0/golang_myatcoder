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

var C [51][51]int // C[n][k] -> nCk

func combTable(N int) {
	C[0][0] = 1
	for i := 1; i <= N; i++ {
		for j := 0; j <= N; j++ {
			if j == 0 || j == i {
				C[i][j] = 1
			} else {
				C[i][j] = C[i-1][j-1] + C[i-1][j]
			}
		}
	}
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1000000)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	k := getI()

	if k == 0 {
		out(4)
		out("0 0 0 0")
		return
	}
	n := k
	cnt := 0
	for n > 0 {
		if n%2 == 1 {
			break
		}
		n /= 2
		cnt++
	}
	p := k / n
	// out(n, p)
	// out("-----")
	for i := 0; i <= cnt; i++ {
		for j := 2; j <= 30; j++ {
			x := j * (j - 1) / 2
			if x*p == k {
				// out(j, cnt)
				out(j + cnt)
				for k := 0; k < j; k++ {
					fmt.Fprint(wr, "1 ")
				}
				for k := 0; k < cnt; k++ {
					fmt.Fprint(wr, "0 ")
				}
				out()
				return
			}
			if x*p > k {
				break
			}
		}
		p /= 2
		cnt--
	}
}
