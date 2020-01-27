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

type Data struct {
	a int
	b int
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
	if p[i].a == p[j].a {
		return p[i].b < p[j].b
	}
	return p[i].a < p[j].a
}

func main() {
	sc.Split(bufio.ScanWords)
	h := getInt()
	w := getInt()
	n := getInt()
	P := make(Datas, n*9)
	for i := 0; i < n; i++ {
		P[i].b = getInt() - 1
		P[i].a = getInt() - 1
	}

	x := [8]int{-1, 0, 1, -1, 1, -1, 0, 1}
	y := [8]int{-1, -1, -1, 0, 0, 1, 1, 1}
	idx := n
	for i := 0; i < n; i++ {
		for j := 0; j < 8; j++ {
			px := P[i].a + x[j]
			py := P[i].b + y[j]
			if (px < w) && (px >= 0) && (py < h) && (py >= 0) {
				P[idx].a = px
				P[idx].b = py
				idx++
			}
		}
	}
	P = P[:idx]

	sort.Sort(P)
	//out(P)
	ans := make([]int, 10)

	xx := -1
	yy := -1
	cnt := 0
	for i := 0; i < len(P); {
		if (P[i].a == 0) || (P[i].a == w-1) || (P[i].b == 0) || (P[i].b == h-1) {
			i++
			continue
		}
		xx = P[i].a
		yy = P[i].b
		cnt = 0
		for ; ; i++ {
			if i >= len(P) {
				break
			}
			if (xx == P[i].a) && (yy == P[i].b) {
				cnt++
			} else {
				break
			}
		}
		ans[cnt]++
	}

	sum := 0
	for i := 1; i < 10; i++ {
		sum += ans[i]
	}
	ans[0] = (w-2)*(h-2) - sum

	for i := 0; i < 10; i++ {
		out(ans[i])
	}
}
