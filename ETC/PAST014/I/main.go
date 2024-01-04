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

func combinations(list []int, k, buf int) (c chan []int) {
	c = make(chan []int, buf)
	n := len(list)

	pattern := make([]int, k)

	var body func(pos, begin int)
	body = func(pos, begin int) {
		if pos == k {
			t := make([]int, k)
			copy(t, pattern)
			c <- t
			return
		}

		for num := begin; num < n+pos-k+1; num++ {
			pattern[pos] = list[num]
			body(pos+1, num+1)
		}
	}
	go func() {
		defer close(c)
		body(0, 0)
	}()

	return
}

type Score struct {
	n  int
	c  byte
	id int
}

func sel(x string, y []string) (int, byte) {
	a := []int{0, 1, 2, 3}

	ret_n, ret_c := 0, byte(0)

	for _, z := range y {
		for b := range combinations(a, 2, 2) {
			m := make(map[byte]int)
			for i := 0; i < len(z); i++ {
				m[z[i]]++
			}
			for _, e := range b {
				m[x[e]]++
			}
			max_n, max_c := 0, byte(0)
			for e := range m {
				if m[e] > max_n {
					max_n = m[e]
					max_c = e
				} else if m[e] == max_n && e < max_c {
					max_n = m[e]
					max_c = e
				}
			}
			if max_n > ret_n {
				ret_n = max_n
				ret_c = max_c
			} else if max_n == ret_n && max_c < ret_c {
				ret_n = max_n
				ret_c = max_c
			}
		}
	}

	return ret_n, ret_c
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N := getI()
	x := getS()

	// n , c, id
	// p := make([]Score, N)

	a := []int{0, 1, 2, 3, 4}
	pat := []string{}
	for b := range combinations(a, 3, 2) {
		t := []byte{}
		for _, e := range b {
			t = append(t, x[e])
		}
		pat = append(pat, string(t))
	}

	p := make([]Score, N)
	for i := 0; i < N; i++ {
		s := getS()
		max_n, max_c := sel(s, pat)
		p[i] = Score{max_n, max_c, i + 1}
	}

	sort.Slice(p, func(i, j int) bool {
		if p[i].n == p[j].n {
			if p[i].c == p[j].c {
				return p[i].id < p[j].id
			}
			return p[i].c < p[j].c
		}
		return p[i].n > p[j].n
	})
	out(p[0].id)
}
