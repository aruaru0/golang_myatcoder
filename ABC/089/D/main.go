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

// Data :
type Data struct {
	n, x, y int
	cost    int
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
	return p[i].n < p[j].n
}

func main() {
	sc.Split(bufio.ScanWords)
	H, W, D := getInt(), getInt(), getInt()
	N := H * W
	n := make(Datas, N)
	idx := 0
	for i := 0; i < H; i++ {
		for j := 0; j < W; j++ {
			a := getInt()
			n[idx] = Data{a, j, i, 0}
			idx++
		}
	}
	sort.Sort(n)
	// out(n, D)

	for i := D; i < N; i++ {
		p := i - D
		d := asub(n[i].x, n[p].x) + asub(n[i].y, n[p].y)
		// out(n[p].x, n[p].y, "-", n[i].x, n[i].y, d)
		n[i].cost = n[p].cost + d
	}
	// out(n, D)

	Q := getInt()
	for i := 0; i < Q; i++ {
		l, r := getInt()-1, getInt()-1
		out(n[r].cost - n[l].cost)
	}

}
