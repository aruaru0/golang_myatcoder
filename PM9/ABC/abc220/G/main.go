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

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

type pair struct {
	f, s int
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	n := getI()
	x, y, c := make([]int, n), make([]int, n), make([]int, n)
	for i := 0; i < n; i++ {
		x[i], y[i], c[i] = getI(), getI(), getI()
	}
	// ２点を結ぶ線分を傾き毎に分類する
	mp := make(map[pair][]pair)
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			dx, dy := x[i]-x[j], y[i]-y[j]
			if dx == 0 {
				dy = 1
			} else {
				if dx < 0 {
					dx = -dx
					dy = -dy
				}
				g := gcd(dx, abs(dy))
				dx /= g
				dy /= g
			}
			mp[pair{dx, dy}] = append(mp[pair{dx, dy}], pair{i, j})
		}
	}

	ans := -1
	for a, ps := range mp {
		// 中点を通る直線上に乗るかどうかで分類
		mp2 := make(map[int][]pair)
		for _, e := range ps {
			i, j := e.f, e.s
			mx := x[i] + x[j]
			my := y[i] + y[j]
			mp2[mx*a.f+my*a.s] = append(mp2[mx*a.f+my*a.s], pair{i, j})
		}
		// 中点を通過する直線２本で得られるＣを計算
		for _, ps2 := range mp2 {
			mp3 := make(map[pair]int)
			for _, e := range ps2 {
				i, j := e.f, e.s
				mx := x[i] + x[j]
				my := y[i] + y[j]
				mp3[pair{mx, my}] = max(mp3[pair{mx, my}], c[i]+c[j])
			}
			if len(mp3) < 2 {
				continue
			}
			//　最大値から２つを選択して値を計算
			cs := make([]int, 0)
			for _, nc := range mp3 {
				cs = append(cs, nc)
			}
			sort.Slice(cs, func(i, j int) bool {
				return cs[i] > cs[j]
			})
			ans = max(ans, cs[0]+cs[1])
		}
	}
	out(ans)
}
