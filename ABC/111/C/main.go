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
	X, V int
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
	return p[i].X > p[j].X
}

func f(a map[int]int) Datas {
	x := make(Datas, 0)
	for i, v := range a {
		x = append(x, Data{v, i})
	}
	sort.Sort(x)
	return x
}

func main() {
	sc.Split(bufio.ScanWords)
	n := getInt()
	x := make(map[int]int)
	y := make(map[int]int)
	for i := 0; i < n; i++ {
		a := getInt()
		if i%2 == 0 {
			x[a]++
		} else {
			y[a]++
		}
	}
	// out(x, y)

	xd := f(x)
	yd := f(y)
	// out(xd[:3])
	// out(yd[:3])
	ans := n - xd[0].X - yd[0].X

	if xd[0].V == yd[0].V {
		if xd[0].X == yd[0].X {
			if len(x) == 1 {
				ans = n - xd[0].X
			} else if xd[1].X > yd[1].X {
				ans = n - xd[0].X - xd[1].X
			} else {
				ans = n - xd[0].X - yd[1].X
			}
		} else if xd[0].X > yd[0].X {
			ans = n - xd[0].X - yd[1].X
		} else {
			ans = n - yd[0].X - xd[1].X
		}
	}

	out(ans)
}
