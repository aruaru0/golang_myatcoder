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
	a, b int
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
	if p[i].b == p[j].b {
		return p[i].a < p[j].a
	}
	return p[i].b < p[j].b
}

func main() {
	sc.Split(bufio.ScanWords)
	_, M := getInt(), getInt()
	n := make(Datas, M)
	for i := 0; i < M; i++ {
		a, b := getInt(), getInt()
		n[i] = Data{a, b}
	}

	sort.Sort(n)
	// out(n, N)

	cut := -1
	ans := 0
	for i := 0; i < M; i++ {
		// out(n[i], cut)
		if n[i].a > cut {
			// out("cut", n[i].b)
			cut = n[i].b - 1
			ans++
		}
	}
	out(ans)
}
