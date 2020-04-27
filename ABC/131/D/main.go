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
	spend, limit int
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
	return p[i].limit < p[j].limit
}

func main() {
	sc.Split(bufio.ScanWords)
	N := getInt()
	d := make(Datas, N)
	for i := 0; i < N; i++ {
		a, b := getInt(), getInt()
		d[i].limit = b
		d[i].spend = a
	}
	sort.Sort(d)

	// out(d)

	ans := true
	start := 0
	for i := 0; i < N; i++ {
		// out(start, i, "limit", d[i].limit, "time", d[i].spend)
		start += d[i].spend
		// out(start)
		if d[i].limit < start {
			ans = false
			break
		}
	}

	if ans {
		out("Yes")
	} else {
		out("No")
	}

}
