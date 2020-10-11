package main

import (
	"bufio"
	"fmt"
	"math/big"
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

func prime(n int) []bool {
	a := make([]bool, n+1)
	for i := 2; i*i <= n; i++ {
		if a[i] == true {
			continue
		}
		for j := 2 * i; j <= n; j += i {
			a[j] = true
		}
	}
	return a
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1000000)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	Y, X := getI(), getI()

	if X == 2 || Y == 2 {
		out("Second")
		return
	}

	if big.NewInt(int64(X)).ProbablyPrime(0) &&
		big.NewInt(int64(Y)).ProbablyPrime(0) {
		out("Second")
		return
	}

	x := X + 1
	for {
		if big.NewInt(int64(x)).ProbablyPrime(0) {
			break
		}
		x++
	}
	y := Y + 1
	for {
		if big.NewInt(int64(y)).ProbablyPrime(0) {
			break
		}
		y++
	}

	cntx := x - 1 - X
	cnty := y - 1 - Y
	// out(y, x, cnty, cntx)

	if (cntx+cnty)%2 == 0 {
		out("Second")
		return
	}
	out("First")
}
