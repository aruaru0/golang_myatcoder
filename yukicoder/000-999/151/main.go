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

func calc(lr, y, t, N int) int {
	t = t % (2 * N)
	if lr == 0 {
		switch {
		case t <= y:
			return y - t
		case t <= N+y:
			return t - y - 1
		default:
			return N - (t - (N + y))
		}
	}
	switch {
	case t < N-y:
		return y + t
	case t < N+N-y:
		return N - 1 - (t - (N - y))
	default:
		return t - (N + (N - y))
	}
}

type fish struct {
	dir, pos, t, num int
}

func main() {
	// for i := 0; i < 100; i++ {
	// 	out("---->", i, calc(0, 0, i, 5))
	// }
	// return
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1000000)
	N, Q := getInt(), getInt()
	f := make([]fish, 0, Q)
	for k := 0; k < Q; k++ {
		x, y, z := getString(), getInt(), getInt()
		switch x {
		case "L":
			f = append(f, fish{0, y, k, z})
		case "R":
			f = append(f, fish{1, y, k, z})
		case "C":
			ans := 0
			for i := 0; i < len(f); i++ {
				pos := calc(f[i].dir, f[i].pos, k-f[i].t, N)
				// out(k-f[i].t, pos, f[i], y, z)
				if pos >= y && pos < z {
					ans += f[i].num
				}
			}
			out(ans)
		}
	}
}
