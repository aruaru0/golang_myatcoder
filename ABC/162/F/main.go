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
	X   int
	pos int
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

func solve(a []int, N int) { // 偶数
	m := 1001001001001
	ans1 := 0
	cnt := 0
	for i := 0; i < N; i += 2 {
		ans1 += a[i]
		// out(ans1)
		cnt++
		m = min(m, a[i])
	}
	ans2 := 0
	cnt = 0
	for i := 1; i < N; i += 2 {
		ans2 += a[i]
		// out(ans2, a[i])
		cnt++
		m = min(m, a[i])
	}
	out(max(ans1, ans2))
}

func solve2(a []int, N int) {
	x := make([]int, N)
	sum := 0
	m := a[0]
	for i := 0; i < N; i += 2 {
		sum += a[i]
		x[i] = sum
		m = min(m, a[i])
		// if i != N-1 {
		// 	x[i+1] = sum
		// }
	}
	y := make([]int, N)
	sum = 0
	for i := 1; i < N; i += 2 {
		sum += a[i]
		y[i] = sum
		// if i != N-1 {
		// 	y[i+1] = sum
		// }
	}

	ans := max(y[N-1], x[N-1]-m)
	for i := 0; i < N-3; i += 2 {
		val := x[i] + y[N-2] - y[i+1]
		out(i, x[i], y[N-2], y[i+1], y[N-2]-y[i+1], val)
		ans = max(ans, val)
	}

	out(m)
	out(x)
	out(y)
	out(ans)
}

func main() {
	sc.Split(bufio.ScanWords)

	N := getInt()
	a := make([]int, N)
	x := make(Datas, N)
	for i := 0; i < N; i++ {
		a[i] = getInt()
		x[i].X = a[i]
		x[i].pos = i
	}

	if N%2 == 0 {
		solve(a, N)
	} else {
		solve2(a, N)
	}

	out("----")
	sort.Sort(x)
	used := make([]bool, N)
	cnt := 0
	ans := 0
	idx := 0
	for cnt < N/2 {
		// out("idx=", idx)
		pos := x[idx].pos
		ok := true
		if pos > 0 {
			if used[pos-1] == true {
				ok = false
			}
		}
		if pos < N-1 {
			if used[pos+1] == true {
				ok = false
			}
		}
		if ok == true {
			out(a[pos])
			ans += a[pos]
			used[pos] = true
			cnt++
		}
		idx++
	}
	out(ans)
}
