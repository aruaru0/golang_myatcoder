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

var N, M int
var node [][]int
var color []int

func dfs(v, c, bit int) bool {
	// out(v, c, bit)
	color[v] = c
	for _, e := range node[v] {
		if color[e] == c {
			return false
		}
		if (bit>>v)%2 == 0 {
			continue
		}
		if color[e] == 0 {
			ret := dfs(e, -c, bit)
			if ret == false {
				return false
			}
		}
	}
	return true
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N, M = getI(), getI()
	node = make([][]int, N)
	for i := 0; i < M; i++ {
		from, to := getI()-1, getI()-1
		node[from] = append(node[from], to)
		node[to] = append(node[to], from)
	}

	ans := 0
	// １色を決めたパターンを全探索
	for bit := 0; bit < 1<<N; bit++ {
		color = make([]int, N)
		for i := 0; i < N; i++ {
			if (bit>>i)%2 == 1 {
				color[i] = 2
			}
		}

		// 決めた色が隣接している場合はNG
		ok := true
	L0:
		for i := 0; i < N; i++ {
			if color[i] == 2 {
				for _, e := range node[i] {
					if color[e] == 2 {
						ok = false
						break L0
					}
				}
			}
		}
		if !ok {
			continue
		}

		// 残りの頂点が２部グラフになっているかチェック
		tot := 1
		for i := 0; i < N; i++ {
			if (bit>>i)%2 == 1 {
				continue
			}
			if color[i] == 0 {
				// out("----")
				// 部分が２部グラフになっていれば２パターンを追加
				if dfs(i, 1, 1<<N-1-bit) {
					tot *= 2
				} else {
					tot *= 0
				}
			}
		}
		// out(bit, color, tot)
		ans += tot
	}
	out(ans)
}
