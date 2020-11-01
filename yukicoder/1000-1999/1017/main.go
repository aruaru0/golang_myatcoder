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

const maxV = 150000*22 + 1

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1000000)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N := getI()
	a := getInts(N)

	n := min(N, 22)
	m := make([][]int, maxV)
	nn := 1 << n
	for i := 0; i < nn; i++ {
		tot := 0
		for k := 0; k < n; k++ {
			if (i>>k)&1 == 1 {
				tot += a[k]
			}
		}
		m[tot] = append(m[tot], i)
	}

	for i := 0; i < maxV; i++ {
		if len(m[i]) > 1 {
			plus := m[i][0]
			minus := m[i][1]
			out("Yes")
			for k := 0; k < n; k++ {
				if (plus>>k)&1 == 1 {
					fmt.Fprint(wr, a[k], " ")
					continue
				}
				if (minus>>k)&1 == 1 {
					fmt.Fprint(wr, -a[k], " ")
					continue
				}
				fmt.Fprint(wr, 0, " ")
			}
			for k := n; k < N; k++ {
				fmt.Fprint(wr, 0, " ")
			}
			out()
			return
		}
	}
	out("No")
}
