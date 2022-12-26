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

const inf = int(1e18)
const mod = int(1e9) + 1

func f(x int) bool {
	s := strconv.Itoa(x)
	for i := 0; i < len(s); i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}

func mk(n int) (int, int) {
	s := strconv.Itoa(n)
	s0 := s
	s1 := s
	for i := len(s) - 1; i >= 0; i-- {
		s0 += string(s[i])
		if i != len(s)-1 {
			s1 += string(s[i])
		}
	}
	r0, _ := strconv.Atoi(s0)
	r1, _ := strconv.Atoi(s1)
	return r0, r1
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1000000)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	cnt := 0
	m := make(map[int]int, 0)
	for {
		cnt++
		r0, r1 := mk(cnt)
		m[r0]++
		m[r1]++
		if cnt > int(1e5) {
			break
		}
	}
	a := make([]int, 0, len(m))
	for e := range m {
		a = append(a, e)
	}
	sort.Ints(a)

	N := getI()
	ans := 0
	for i := 0; i < len(a); i++ {
		x := a[i] * mod
		if x/mod != a[i] {
			break
		}
		if a[i]*mod > N {
			break
		}
		ans++
	}
	out(ans)
}
