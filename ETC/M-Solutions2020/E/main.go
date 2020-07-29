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

type town struct {
	x, y, p int
}

const inf = int(1e15)

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1000000)
	N := getInt()
	X := make([]int, N)
	Y := make([]int, N)
	P := make([]int, N)
	for i := 0; i < N; i++ {
		X[i], Y[i], P[i] = getInt(), getInt(), getInt()
	}

	dist := make([]int, N+1)
	n := 1
	for i := 0; i < N; i++ {
		n *= 3
		dist[i] = inf
	}

	lx := make([]int, 0, N)
	ly := make([]int, 0, N)
	l := make([]int, 0, N)

	for i := 0; i < n; i++ {
		// if i%100000 == 0 {
		// 	out(i)
		// }
		lx = lx[:0]
		ly = ly[:0]
		l = l[:0]
		p := i
		for j := 0; j < N; j++ {
			switch p % 3 {
			case 0: // nothing
				l = append(l, j)
			case 1: // add x
				lx = append(lx, X[j])
			case 2: // add y
				ly = append(ly, Y[j])
			}
			p /= 3
		}
		sum := 0
		for _, x := range l {
			d := min(abs(X[x]), abs(Y[x]))
			for _, px := range lx {
				d = min(d, abs(px-X[x]))
			}
			for _, py := range ly {
				d = min(d, abs(py-Y[x]))
			}
			sum += d * P[x]
		}
		pos := len(lx) + len(ly)
		dist[pos] = min(dist[pos], sum)
	}
	for i := 0; i <= N; i++ {
		out(dist[i])
	}
}
