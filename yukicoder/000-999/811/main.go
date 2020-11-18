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

func psf(n int) (map[int]int, int) {
	ret := make(map[int]int)
	cnt := 0
	for i := 2; i*i <= n; i++ {
		for n%i == 0 {
			ret[i]++
			cnt++
			n /= i
		}
	}
	if n != 1 {
		ret[n]++
		cnt++
	}
	return ret, cnt
}

const inf = int(1e10)

func match(p, q map[int]int) int {
	ret := 0
	for n, c := range p {
		ret += min(c, q[n])
	}
	return ret
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1000000)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N, K := getI(), getI()
	p, _ := psf(N)

	ans := 0
	ma := 0
	for i := 1; i < N; i++ {
		x, _ := psf(i)
		k := match(p, x)
		if k < K {
			continue
		}
		// out(i, p, x, k)
		tot := 1
		for _, e := range x {
			tot *= e + 1
		}
		if tot > ma {
			ma = tot
			ans = i
		}
	}
	out(ans)
}
