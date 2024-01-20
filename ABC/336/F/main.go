package main

import (
	"bufio"
	"fmt"
	"math"
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

func chmin(a *int, b int) bool {
	if *a < b {
		return false
	}
	*a = b
	return true
}

func chmax(a *int, b int) bool {
	if *a > b {
		return false
	}
	*a = b
	return true
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

func rot(s [][]int, sy, sx int) {
	nh, nw := H-1, W-1
	n := (H - 1) * (W - 1) / 2
	cnt := 0
	for y := 0; y < nh; y++ {
		for x := 0; x < nw; x++ {
			s[sy+y][sx+x], s[sy+nh-1-y][sx+nw-1-x] = s[sy+nh-1-y][sx+nw-1-x], s[sy+y][sx+x]
			cnt++
			if cnt == n {
				return
			}
		}
	}
}

func hash(cur [][]int) [64]byte {
	var x [64]byte
	cnt := 0
	for h := 0; h < H; h++ {
		for w := 0; w < W; w++ {
			x[cnt] = byte(cur[h][w])
			cnt++
		}
	}
	return x
}

func bsf(s [][]int, N int) map[[64]byte]int {
	m := make(map[[64]byte]int)

	q := make([][][]int, 0)
	q = append(q, s)
	m[hash(s)] = 0
	for n := 0; n < N; n++ {
		nxt := make([][][]int, 0)
		for len(q) != 0 {
			cur := q[0]
			hcur := hash(cur)
			q = q[1:]
			for h := 0; h < 2; h++ {
				for w := 0; w < 2; w++ {
					tmp := make([][]int, H)
					for i := 0; i < H; i++ {
						tmp[i] = make([]int, W)
						for j := 0; j < W; j++ {
							tmp[i][j] = cur[i][j]
						}
					}
					rot(tmp, h, w)
					hashed := hash(tmp)
					if _, ok := m[hashed]; ok == true {
						continue
					}
					m[hashed] = m[hcur] + 1
					nxt = append(nxt, tmp)
				}
			}
		}
		q = nxt
	}

	return m
}

var H, W int
var tot int
var memo map[[64]byte]bool

const inf = int(1e18)

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	H, W = getI(), getI()
	s := make([][]int, H)
	for i := 0; i < H; i++ {
		s[i] = getInts(W)
	}
	t := make([][]int, H)
	for i := 0; i < H; i++ {
		t[i] = make([]int, W)
		for j := 0; j < W; j++ {
			t[i][j] = i*W + j + 1
		}
	}

	m0 := bsf(s, 10)
	m1 := bsf(t, 10)

	ans := inf
	for e := range m0 {
		if _, ok := m1[e]; ok {
			ans = min(ans, m0[e]+m1[e])
		}
	}
	if ans == inf {
		out(-1)
	} else {
		out(ans)
	}
}
