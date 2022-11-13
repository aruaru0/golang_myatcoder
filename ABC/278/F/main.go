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

var N int
var s []string
var memo [][1 << 16]*bool

func rec(x, bit int) bool {
	if memo[x][bit] != nil {
		return *memo[x][bit]
	}
	// cnt := bits.OnesCount(uint(bit))
	last := s[x][len(s[x])-1]
	// 答えられなかったら負け
	ok := false
	win := true
	for i := 0; i < N; i++ {
		// もし使ってなくて、つながるなら
		if (bit>>i)%2 == 0 && s[i][0] == last {
			ok = true
			// 再帰呼び出し
			ret := rec(i, bit|1<<i)
			// 勝ち負け判定
			//  相手の勝ち筋が１つでもあれば負け
			if ret == true {
				win = false
			}
		}
	}

	win = win && ok

	memo[x][bit] = &win

	return win
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N = getI()
	s = make([]string, N)
	for i := 0; i < N; i++ {
		s[i] = getS()
	}

	memo = make([][1 << 16]*bool, N)

	ans := false
	for i := 0; i < N; i++ {
		ret := rec(i, 1<<i)
		if ret == false {
			ans = true
		}
	}

	if ans {
		out("First")
	} else {
		out("Second")
	}
}
