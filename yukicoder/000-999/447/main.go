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

type user struct {
	score [27]int
	total int
	last  int
	name  string
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1000000)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N := getI()
	l := getInts(N)
	a := make([]int, N)
	T := getI()
	player := make(map[string]user)
	for i := 0; i < T; i++ {
		name, p := getS(), getS()
		num := int(p[0] - 'A')

		a[num]++ // count anser
		score := 50*l[num] + 500*l[num]/(8+2*a[num])
		v, _ := player[name]
		v.last = i
		v.score[num] = score
		v.total += score
		v.name = name
		player[name] = v
	}

	ans := make([]user, 0)
	for _, e := range player {
		ans = append(ans, e)
	}
	sort.Slice(ans, func(i, j int) bool {
		if ans[i].total == ans[j].total {
			return ans[i].last < ans[j].last
		}
		return ans[i].total > ans[j].total
	})

	for i, e := range ans {
		fmt.Fprint(wr, i+1, " ", e.name, " ")
		for j := 0; j < N; j++ {
			fmt.Fprint(wr, e.score[j], " ")
		}
		out(e.total)
	}
}
