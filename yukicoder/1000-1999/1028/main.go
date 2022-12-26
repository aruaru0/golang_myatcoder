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
	sc.Buffer([]byte{}, 1000000)

	N := getInt()
	a := make([][]int, N)
	b := make([]int, N)
	c := make([]int, N)
	d := make([][2]int, N)
	for i := 0; i < N; i++ {
		a[i] = make([]int, N)
		for j := 0; j < N; j++ {
			a[i][j] = getInt() - 1
			b[a[i][j]] += i + j - min(i, j)
			if j-i < 0 {
				d[a[i][j]][1]++
			}
		}
	}
	for i := 0; i < N; i++ {
		c[i] = b[i]
	}
	for i := 0; i < N; i++ {
		for j := 0; i-j >= 0; j++ {
			d[a[i-j][j]][0]++
		}
		for j := 0; j < N; j++ {
			b[j] += d[j][0] - d[j][1]
			c[j] = min(c[j], b[j])
		}
		for j := 0; i+j+1 < N; j++ {
			d[a[i+j+1][j]][1]--
		}
	}
	ans := 0
	for i := 0; i < N; i++ {
		ans += c[i]
	}
	out(ans)
}
