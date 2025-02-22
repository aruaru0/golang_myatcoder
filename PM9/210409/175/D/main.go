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

const inf = int(1e15)

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()

	// 輪を見つけて、ループすると得ならループさせる
	// 輪の部分で、最大になるところを見つける
	// (1)輪の部分で最大　または、(2)ループして、あまりで最大を稼ぐ
	// の最大が最大。部分最大の取り方を考えるのがポイント
	N, K := getI(), getI()
	p := make([]int, N)
	for i := 0; i < N; i++ {
		p[i] = getI() - 1
	}
	c := getInts(N)

	route := make([][]int, 0)
	used := make([]bool, N)
	for i := 0; i < N; i++ {
		if used[i] == true {
			continue
		}
		r := []int{i}
		next := p[i]
		used[i] = true
		for next != i {
			r = append(r, next)
			used[next] = true
			next = p[next]
		}
		route = append(route, r)
	}

	tot := make([]int, len(route))
	for i, e := range route {
		tot[i] = 0
		for _, v := range e {
			tot[i] += c[v]
		}
	}
	rmax := make([][]int, len(route))
	for i, e := range route {
		x := append(e, e...)
		x = append(x, e...)
		rmax[i] = make([]int, len(x)+1)
		for j := 1; j <= len(x); j++ {
			rmax[i][j] = -inf
		}
		for j := 0; j < len(x); j++ {
			sum := c[x[j]]
			cnt := 1
			chmax(&rmax[i][cnt], sum)
			for k := j + 1; k < len(x); k++ {
				sum += c[x[k]]
				cnt++
				chmax(&rmax[i][cnt], sum)
			}
		}
	}

	ans := -inf
	for i := 0; i < len(route); i++ {
		// out(tot[i], rmax[i])

		if tot[i] > 0 {
			n := len(route[i])
			// out(tot[i], n, "LOOP")
			loop := max(0, K/n-1)
			sum := max(0, tot[i]*loop)
			// out(sum, K%n)
			rest := 0
			for j := 0; j <= K-loop*n; j++ {
				rest = max(rest, rmax[i][j])
			}
			ans = max(ans, sum+rest)
		} else {
			n := min(K, len(rmax[i])-1)
			rest := -inf
			for j := 1; j <= n; j++ {
				// out(n, rmax[i])
				rest = max(rest, rmax[i][j])
			}
			// out(rest, "REST")
			ans = max(ans, rest)
		}
	}
	out(ans)
}
