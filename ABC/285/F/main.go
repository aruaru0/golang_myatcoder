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
		return b.sum(y) - b.sum(x-1)
	}
}
func (b BIT) add(a, w int) {
	n := len(b.v)
	for i := a + 1; i <= n; i += i & -i {
		b.v[i-1] += w
	}
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N := getI()
	s := []byte(getS())

	sorted := newBIT(N)
	cnt := make([]*BIT, 26)
	for i := 0; i < 26; i++ {
		cnt[i] = newBIT(N)
	}

	for i := 0; i < N; i++ {
		pos := int(s[i] - 'a')
		cnt[pos].add(i, 1)
	}
	for i := 0; i < N-1; i++ {
		if s[i] > s[i+1] {
			sorted.add(i, 1)
		}
	}

	add := func(i, x int) {
		pos := int(s[i] - 'a')
		cnt[pos].add(i, x)
		if 1 <= i && s[i-1] > s[i] {
			sorted.add(i-1, x)
		}
		if i+1 < N && s[i] > s[i+1] {
			sorted.add(i, x)
		}
	}

	Q := getI()
	for qi := 0; qi < Q; qi++ {
		t := getI()
		switch t {
		case 1:
			i, c := getI()-1, getS()
			add(i, -1)
			s[i] = byte(c[0])
			add(i, 1)
		case 2:
			l, r := getI()-1, getI()
			ok := true
			if sorted.rangeSum(l, r-1) != 0 {
				ok = false
			}
			for i := s[l] + 1; i < s[r-1]; i++ {
				pos := int(i - 'a')
				if cnt[pos].rangeSum(l, r) != cnt[pos].rangeSum(0, N) {
					ok = false
				}
			}
			if ok {
				out("Yes")
			} else {
				out("No")
			}
		}
	}
}
