package main

import (
	"bufio"
	"fmt"
	"math"
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

const inf = math.MaxUint64 >> 1

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1000000)
	N, K := getInt(), getInt()
	p := make([]int, N)
	for i := 0; i < N; i++ {
		p[i] = getInt() - 1
	}
	c := getInts(N)
	ans := -inf

	for i := 0; i < N; i++ {
		used := make([]bool, N)
		pos := p[i]
		sum := 0
		l := 0
		for used[pos] == false {
			used[pos] = true
			l++
			sum += c[pos]
			pos = p[pos]
		}
		n := K / l
		m := K % l
		if n > 0 {
			n--
			m += l
		}
		if sum > 0 {
			sum = sum * n
		} else {
			sum = 0
		}

		pos = p[i]
		ma := -inf
		for j := 0; j < m; j++ {
			sum += c[pos]
			pos = p[pos]
			ma = max(ma, sum)
		}
		ans = max(ans, ma)
	}

	out(ans)
}
