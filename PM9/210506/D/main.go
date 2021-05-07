package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
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

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N := getI()

	sa := make([]string, 0)
	sb := make([]string, 0)
	sab := make([]string, 0)
	s := make([]string, 0)
	for i := 0; i < N; i++ {
		t := getS()
		if t[0] == 'B' && t[len(t)-1] == 'A' {
			sab = append(sab, t)
			continue
		}
		if t[0] == 'B' {
			sb = append(sb, t)
			continue
		}
		if t[len(t)-1] == 'A' {
			sa = append(sa, t)
			continue
		}
		s = append(s, t)
	}

	ans := ""
	cur := 0
	for {
		// out(sa, sb, sab, s, cur)
		if cur == 0 && len(sa) != 0 {
			ans += sa[0]
			sa = sa[1:]
			cur = 1
			continue
		}
		if cur == 1 && len(sab) != 0 {
			ans += sab[0]
			sab = sab[1:]
			cur = 1
			continue
		}
		if cur == 1 && len(sb) != 0 {
			ans += sb[0]
			sb = sb[1:]
			cur = 0
			continue
		}
		if len(sab) != 0 {
			ans += sab[0]
			sab = sab[1:]
			cur = 1
			continue
		}
		if len(sa) != 0 {
			ans += sa[0]
			sa = sa[1:]
			cur = 1
			continue
		}
		if len(sb) != 0 {
			ans += sb[0]
			sb = sb[1:]
			cur = 1
			continue
		}
		if len(s) != 0 {
			ans += s[0]
			s = s[1:]
			cur = 1
			continue
		}
		break
	}

	// out(ans)
	out(strings.Count(ans, "AB"))

}
