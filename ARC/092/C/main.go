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

type point struct {
	x, y int
}

type points []point

var sel = false

func (p points) Len() int {
	return len(p)
}

func (p points) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p points) Less(i, j int) bool {
	if sel == true {
		return p[i].y > p[j].y
	}
	return p[i].x < p[j].x
}

func main() {
	sc.Split(bufio.ScanWords)
	N := getInt()
	ab := make(points, N)
	cd := make(points, N)
	for i := 0; i < N; i++ {
		ab[i].x, ab[i].y = getInt(), getInt()
	}
	for i := 0; i < N; i++ {
		cd[i].x, cd[i].y = getInt(), getInt()
	}

	sort.Sort(cd)
	sel = true
	sort.Sort(ab)
	ans := 0
	// out(ab)
	// out(cd)
	used := make([]bool, N)
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			// out(ab[j], cd[i])
			if used[j] {
				continue
			}
			if cd[i].x > ab[j].x && cd[i].y > ab[j].y {
				// out("ok")
				used[j] = true
				ans++
				break
			}
		}
	}
	out(ans)
}
