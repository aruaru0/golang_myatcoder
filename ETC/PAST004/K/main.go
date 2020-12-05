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

const mod = int(1e9)

type BIT struct {
	v []int
}

func newBIT(n int) *BIT {
	b := new(BIT)
	b.v = make([]int, n)
	return b
}
func (b BIT) sum(a int) int {
	ret := 0
	for i := a + 1; i > 0; i -= i & -i {
		ret += b.v[i-1]
		ret %= mod
	}
	return ret
}
func (b BIT) rangeSum(x, y int) int {
	if y == 0 {
		return 0
	}
	y--
	if x == 0 {
		return b.sum(y)
	} else {
		ret := b.sum(y) - b.sum(x-1)
		ret %= mod
		if ret < 0 {
			ret += mod
		}
		return ret
	}
}
func (b BIT) add(a, w int) {
	n := len(b.v)
	for i := a + 1; i <= n; i += i & -i {
		b.v[i-1] += w
		b.v[i-1] %= mod
	}
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1000000)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	K := getI()
	a := make([][]int, K)

	for i := 0; i < K; i++ {
		n := getI()
		a[i] = getInts(n)
	}

	n := make([][22]int, K)
	t := make([][22]int, K)
	c := make([]int, K)
	for i := 0; i < K; i++ {
		bit := newBIT(22)
		cnt := 0
		for j := 0; j < len(a[i]); j++ {
			n[i][a[i][j]]++
			cnt += bit.rangeSum(a[i][j]+1, 21)
			bit.add(a[i][j], 1)
		}
		c[i] = cnt
		t[i][0] = n[i][0]
		for j := 1; j < 20; j++ {
			t[i][j] = t[i][j-1] + n[i][j]
		}
	}

	Q := getI()
	tot := make([]int, 22)
	ans := 0
	for i := 0; i < Q; i++ {
		b := getI() - 1
		x := 0
		cnt := 0
		for j := 0; j < 20; j++ {
			x += n[b][j]
			cnt += x * tot[j+1]
			cnt %= mod
		}
		// out(n[b])
		// out(tot)
		// out(cnt)
		for j := 0; j < 21; j++ {
			tot[j] += n[b][j]
		}
		ans += c[b] + cnt
		ans %= mod
	}
	out(ans)
}
