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

type position struct {
	x int
	y int
}

type status struct {
	turn  int
	money int
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	n := getI()
	ps := make([][]int, n)
	for i := 0; i < n; i++ {
		ps[i] = getInts(n)
	}
	rs := make([][]int, n)
	for i := 0; i < n; i++ {
		rs[i] = getInts(n - 1)
	}
	ds := make([][]int, n-1)
	for i := 0; i < n-1; i++ {
		ds[i] = getInts(n)
	}

	dp := make(map[position]map[int]status)
	ini := make(map[int]status)
	ini[0] = status{0, 0}
	dp[position{0, 0}] = ini

	for x := 0; x < n; x++ {
		for y := 0; y < n; y++ {
			pos := position{x, y}

			stmap := dp[pos]
			for maxpay, st := range stmap {
				maxpay = max(maxpay, ps[pos.x][pos.y]) // 経路中の最大のPを求める

				nextArr := make([]position, 0)
				costArr := make([]int, 0)

				// 次に進める場所の場所とコストを計算
				if pos.y < n-1 {
					nextArr = append(nextArr, position{pos.x, pos.y + 1})
					costArr = append(costArr, rs[pos.x][pos.y])
				}
				if pos.x < n-1 {
					nextArr = append(nextArr, position{pos.x + 1, pos.y})
					costArr = append(costArr, ds[pos.x][pos.y])
				}

				// 行ける場所について
				for idx, nextPos := range nextArr {
					nextCost := costArr[idx]
					nextSt := status{money: st.money, turn: st.turn}
					// 次へのお金とコストを計算
					if nextSt.money < nextCost {
						cnt := (nextCost - nextSt.money) / maxpay
						if (nextCost-nextSt.money)%maxpay > 0 {
							cnt++
						}
						nextSt.money += maxpay * cnt
						nextSt.turn += cnt
					}
					nextSt.turn++
					nextSt.money -= nextCost
					oldStMap, ok := dp[nextPos]
					if !ok {
						dp[nextPos] = make(map[int]status)
					} else {
						oldSt, ok := oldStMap[maxpay]
						if ok {
							if oldSt.turn < nextSt.turn {
								continue
							}
							if oldSt.turn == nextSt.turn && oldSt.money >= nextSt.money {
								continue
							}
						}
					}
					dp[nextPos][maxpay] = nextSt
				}
			}
		}
	}

	ans := -1
	for _, v := range dp[position{n - 1, n - 1}] {
		if ans == -1 || ans > v.turn {
			ans = v.turn
		}
	}

	out(ans)
}
