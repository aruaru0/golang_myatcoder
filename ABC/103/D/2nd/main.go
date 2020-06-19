package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func out(x ...interface{}) {
	fmt.Println(x...)
}

var sc = bufio.NewScanner(os.Stdin)

func getInt() int {
	sc.Scan()
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return i
}

func getString() string {
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

// Data :
type Data struct {
	f, t int
}

// Datas :
type Datas []Data

func (p Datas) Len() int {
	return len(p)
}

func (p Datas) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p Datas) Less(i, j int) bool {
	return p[i].f < p[j].f
}

func main() {
	sc.Split(bufio.ScanWords)
	_, M := getInt(), getInt()
	p := make(Datas, M)
	for i := 0; i < M; i++ {
		a, b := getInt(), getInt()
		p[i] = Data{a, b}
	}
	sort.Sort(p)
	// out(p, N)
	l := p[0].f
	r := p[0].t
	ans := 0
	for i := 1; i < M; i++ {
		if r <= p[i].f {
			l = p[i].f
			r = p[i].t
			// out("cut")
			ans++
		} else {
			l = max(l, p[i].f)
			r = min(r, p[i].t)
		}
		// out(l, r)
	}
	if l != r {
		ans++
	}
	out(ans)
}
