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

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	s := getS()
	t := getS()
	ns := make([]bool, 26)
	nt := make([]bool, 26)

	n := len(s)
	p := make([][]int, 26)
	for i := 0; i < 26; i++ {
		p[i] = make([]int, n)
		for j := 0; j < n; j++ {
			p[i][j] = -1
		}
	}
	for i := 0; i < len(s); i++ {
		x := int(s[i] - 'a')
		ns[x] = true
		p[x][i] = i
	}
	for i := 0; i < len(t); i++ {
		x := int(t[i] - 'a')
		nt[x] = true
	}
	ok := true
	for i := 0; i < 26; i++ {
		if nt[i] == true && ns[i] == false {
			ok = false
		}
	}
	if !ok {
		out(-1)
		return
	}

	for i := 0; i < 26; i++ {
		for j := len(s) - 2; j >= 0; j-- {
			if p[i][j] == -1 {
				p[i][j] = p[i][j+1]
			}
		}
		//	out(string(byte(i)+'a'), p[i])
	}

	tot := 0
	pos := 0
	for i := 0; i < len(t); i++ {
		d := int(t[i] - 'a')
		next := p[d][pos]
		if next == -1 {
			tot += len(s)
			next = p[d][0]
		}
		if next+1 == n {
			tot += n
		}
		pos = (next + 1) % n
	}
	// out(tot, pos)
	out(tot + pos)
}
