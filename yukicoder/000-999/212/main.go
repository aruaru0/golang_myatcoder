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

func getInts(N int) []int {
	ret := make([]int, N)
	for i := 0; i < N; i++ {
		ret[i] = getInt()
	}
	return ret
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

func main() {
	sc.Split(bufio.ScanWords)
	P, C := getInt(), getInt()

	x := []int{2, 3, 5, 7, 11, 13}
	p0 := make(map[int]int)
	p0[1] = 1
	for i := 0; i < P; i++ {
		p1 := make(map[int]int)
		for j, e := range p0 {
			for _, v := range x {
				p1[v*j] += e
			}
		}
		p0 = p1
	}

	y := []int{4, 6, 8, 9, 10, 12}
	q0 := make(map[int]int)
	q0[1] = 1
	for i := 0; i < C; i++ {
		q1 := make(map[int]int)
		for j, e := range q0 {
			for _, v := range y {
				q1[v*j] += e
			}
		}
		q0 = q1
	}

	sum := 0
	cnt := 0
	for i, e := range p0 {
		for j, v := range q0 {
			cnt += e * v
			sum += i * j * e * v
		}
	}
	// out(p0, q0)
	out(float64(sum) / float64(cnt))
}
