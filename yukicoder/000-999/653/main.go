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

var memo [21000][2][2][2]int

func rec(pos, carry, a1, a2 int) int {
	if pos < 0 {
		if carry == 0 {
			return 1
		}
		return 0
	}

	if memo[pos][carry][a1][a2] != -1 {
		return memo[pos][carry][a1][a2]
	}

	v := int(p[pos] - '0')

	ok := 0

	if a1 != 0 && a2 != 0 {
		x := carry
		if (x+7+7)%10 == v {
			ok |= rec(pos-1, 1, 1, 1)
		}
		if (x+7+6)%10 == v {
			ok |= rec(pos-1, 1, 1, 1)
		}
		if (x+6+6)%10 == v {
			ok |= rec(pos-1, 1, 1, 1)
		}
	}

	if a1 != 0 {
		x := carry
		if x+7 == v {
			ok |= rec(pos-1, 0, 1, 0)
		}
		if x+6 == v {
			ok |= rec(pos-1, 0, 1, 0)
		}
	}

	if a2 != 0 {
		x := carry
		if x+7 == v {
			ok |= rec(pos-1, 0, 0, 1)
		}
		if x+6 == v {
			ok |= rec(pos-1, 0, 0, 1)
		}
	}

	if v == carry {
		ok |= rec(pos-1, 0, 0, 0)
	}

	memo[pos][carry][a1][a2] = ok
	return ok
}

var p string

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1000000)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	p = getS()

	for i := 0; i < 21000; i++ {
		for j := 0; j < 2; j++ {
			for k := 0; k < 2; k++ {
				for l := 0; l < 2; l++ {
					memo[i][j][k][l] = -1
				}
			}
		}
	}

	x := p[len(p)-1]
	if x == '6' || x == '7' {
		out("No")
		return
	}

	ret := rec(len(p)-1, 0, 1, 1)

	if ret == 1 {
		out("Yes")
		return
	}
	out("No")
}
