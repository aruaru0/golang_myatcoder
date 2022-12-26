package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)
var wr = bufio.NewWriter(os.Stdout)

func out(x ...interface{}) {
	fmt.Fprintln(wr, x...)
}

func getI() int {
	sc.Scan()
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return i
}

func getF() float64 {
	sc.Scan()
	i, e := strconv.ParseFloat(sc.Text(), 64)
	if e != nil {
		panic(e)
	}
	return i
}

func getInts(N int) []int {
	ret := make([]int, N)
	for i := 0; i < N; i++ {
		ret[i] = getI()
	}
	return ret
}

func getS() string {
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

// min for n entry
func nmin(a ...int) int {
	ret := a[0]
	for _, e := range a {
		ret = min(ret, e)
	}
	return ret
}

// max for n entry
func nmax(a ...int) int {
	ret := a[0]
	for _, e := range a {
		ret = max(ret, e)
	}
	return ret
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

func rot(x [][]byte) [][]byte {
	h := len(x)
	w := len(x[0])
	ret := make([][]byte, w)
	for i := 0; i < w; i++ {
		ret[i] = make([]byte, h)
	}

	for i := 0; i < h; i++ {
		for j := 0; j < w; j++ {
			ret[j][i] = x[i][w-1-j]
		}
	}

	return ret
}

func f(y, x int, s, t [][]byte) int {
	H := len(s)
	W := len(s[0])
	h := len(t)
	w := len(t[0])

	// out("s:", x, y)
	// for i := y; i < H; i++ {
	// 	for j := x; j < W; j++ {
	// 		fmt.Fprint(wr, string(s[i][j]))
	// 	}
	// 	out()
	// }

	cnt := 0
	cy := 0
	for i := y; i < H; i++ {
		if cy >= h {
			break
		}
		cx := 0
		for j := x; j < W; j++ {
			if cx >= w {
				break
			}
			if t[cy][cx] == '#' && s[i][j] == '.' {
				cnt++
			}
			cx++
		}
		cy++
	}
	return cnt
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1000000)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	H, W := getI(), getI()
	s := make([][]byte, H)
	t := make([][]byte, H)
	for i := 0; i < H; i++ {
		s[i] = []byte(getS())
	}
	a := make([]int, H)
	b := make([]int, W)
	for i := 0; i < H; i++ {
		t[i] = []byte(getS())
	}
	tot := 0
	for i := 0; i < H; i++ {
		for j := 0; j < W; j++ {
			if t[i][j] == '#' {
				a[i] = 1
				b[j] = 1
				tot++
			}
		}
	}
	tx := make([][]byte, 0)
	for i := 0; i < H; i++ {
		if a[i] == 0 {
			continue
		}
		v := make([]byte, 0)
		for j := 0; j < W; j++ {
			if b[j] == 1 {
				v = append(v, t[i][j])
			}
		}
		tx = append(tx, v)
	}

	t = tx

	for k := 0; k < 4; k++ {
		// out("T")
		// for i := 0; i < len(t); i++ {
		// 	out(string(t[i]))
		// }
		for i := 0; i < H; i++ {
			for j := 0; j < W; j++ {
				ret := f(i, j, s, t)
				if ret == tot {
					out("Yes")
					return
				}
			}
		}
		t = rot(t)
		// out("----")
	}
	out("No")
}
