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

type pair struct{ x, y int }
type triple struct{ x, y, z int }

func dot2(a, b pair) int   { return a.x*b.x + a.y*b.y }           // 内積
func cross2(a, b pair) int { return a.x*b.y - a.y*b.x }           // 外積
func dot3(a, b triple) int { return a.x*b.x + a.y*b.y + a.x*b.z } // 内積
func cross3(a, b triple) triple { // 外積
	return triple{a.y*b.z - a.z*b.y, a.z*b.x - a.x*b.z, a.x*b.y - a.y*b.x}
}
func pt2sub(a, b pair) pair       { return pair{a.x - b.x, a.y - b.y} }              // 座標の減算
func pt2add(a, b pair) pair       { return pair{a.x + b.x, a.y + b.y} }              // 座標の加算
func pt2scale(n int, a pair) pair { return pair{n * a.x, n * a.y} }                  // 倍率をかける
func area2x(a, b, c pair) int     { return abs(cross2(pt2sub(b, a), pt2sub(c, a))) } // 面積の２倍（頂点aを原点に移動させて内積）

const MOD = int(1e9 + 7)

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()

	N := getI()
	X := make([]int, N)
	Y := make([]int, N)
	for i := 0; i < N; i++ {
		X[i] = getI()
		Y[i] = getI()
	}
	pts := make([]pair, N)
	for i := 0; i < N; i++ {
		pts[i] = pair{X[i], Y[i]}
	}
	// x順でソート
	sort.Slice(pts, func(i, j int) bool {
		return pts[i].x < pts[j].x || pts[i].x == pts[j].x && pts[i].y < pts[j].y
	})
	parity := [80][80][80]int{}
	numinside := [80][80][80]int{}
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			for k := 0; k < N; k++ {
				if i == j || i == k || j == k { // 同じ点は除く
					continue
				}
				a := area2x(pts[i], pts[j], pts[k]) // 面積を計算
				parity[i][j][k] = a & 1             // 面積が整数かどうかのフラグ
				for l := 0; l < N; l++ {
					if l == i || l == j || l == k { // 同じ点は除く
						continue
					}
					// lが、三角形i,j,kの内部かどうか（面積でチェックする）
					a2 := 0
					a2 += area2x(pts[i], pts[j], pts[l])
					a2 += area2x(pts[j], pts[k], pts[l])
					a2 += area2x(pts[k], pts[i], pts[l])
					if a2 == a {
						numinside[i][j][k]++
					}
				}
			}
		}
	}

	// 2のn乗を前計算
	pow2 := make([]int, N+1)
	pow2[0] = 1
	for i := 1; i <= N; i++ {
		pow2[i] = 2 * pow2[i-1]
		if pow2[i] >= MOD {
			pow2[i] -= MOD
		}
	}

	// DP　※解けなかったので、他の解答をコピー（いまいち、ピンとこない）
	upper := [80][80][2]int{}
	lower := [80][80][2]int{}
	ans := 0
	for leftmost := N - 1; leftmost >= 0; leftmost-- {
		for i := leftmost; i < N; i++ {
			for j := leftmost; j < N; j++ {
				for k := 0; k < 2; k++ {
					upper[i][j][k] = 0
					lower[i][j][k] = 0
				}
			}
		}
		for i := leftmost + 1; i < N; i++ {
			upper[leftmost][i][0] = 1
			lower[leftmost][i][0] = 1
		}
		for i := leftmost; i < N; i++ {
			for j := i + 1; j < N; j++ {
				for k := 0; k < 2; k++ {
					for l := j + 1; l < N; l++ {
						if cross2(pt2sub(pts[l], pts[j]), pt2sub(pts[j], pts[i])) > 0 {
							upper[j][l][k^parity[leftmost][j][l]] += upper[i][j][k] * pow2[numinside[leftmost][j][l]] % MOD
							upper[j][l][k^parity[leftmost][j][l]] %= MOD

						} else {
							lower[j][l][k^parity[leftmost][j][l]] += lower[i][j][k] * pow2[numinside[leftmost][j][l]] % MOD
							lower[j][l][k^parity[leftmost][j][l]] %= MOD
						}
					}

				}
			}
		}
		for j := leftmost + 1; j < N; j++ {
			for k := 0; k < 2; k++ {
				up, lo := 0, 0
				for i := leftmost; i < j; i++ {
					up += upper[i][j][k]
					if up >= MOD {
						up -= MOD
					}
					lo += lower[i][j][k]
					if lo >= MOD {
						lo -= MOD
					}
				}
				ans += up * lo % MOD
				ans %= MOD
			}
		}
	}
	ans = ans + MOD - (N)*(N-1)/2
	ans %= MOD
	out(ans)
}
