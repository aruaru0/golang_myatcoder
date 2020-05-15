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

const inf = 1001001001001

// Data :
type Data struct {
	x, num int
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
	return p[i].num > p[j].num
}
func main() {
	sc.Split(bufio.ScanWords)
	n := getInt()
	a0 := make(map[int]int)
	a1 := make(map[int]int)
	c0 := 0
	c1 := 0
	for i := 0; i < n; i++ {
		v := getInt()
		if i%2 == 0 {
			a0[v]++
			c0++
		} else {
			a1[v]++
			c1++
		}
	}

	b0 := make(Datas, 0)
	b1 := make(Datas, 0)
	for i, v := range a0 {
		b0 = append(b0, Data{i, v})
	}
	b0 = append(b0, Data{inf, 0})
	for i, v := range a1 {
		b1 = append(b1, Data{i, v})
	}
	b1 = append(b1, Data{inf, 0})

	sort.Sort(b0)
	sort.Sort(b1)

	ans := 0
	if b0[0].x == b1[0].x {
		ans = min(
			c0-b0[0].num+c1-b1[1].num,
			c0-b0[1].num+c1-b1[0].num)
	} else {
		ans = c0 - b0[0].num + c1 - b1[0].num
	}

	out(ans)
}
