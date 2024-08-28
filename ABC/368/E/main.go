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

func outSlice[T any](s []T) {
	for i := 0; i < len(s)-1; i++ {
		fmt.Fprint(wr, s[i], " ")
	}
	fmt.Fprintln(wr, s[len(s)-1])
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

type Event struct {
	time, flg, idx int
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N, M, X := getI(), getI(), getI()

	A := make([]int, M)
	B := make([]int, M)
	S := make([]int, M)
	T := make([]int, M)
	event := make([]Event, 0)
	for i := 0; i < M; i++ {
		A[i], B[i], S[i], T[i] = getI(), getI(), getI(), getI()
		event = append(event, Event{S[i], 1, i})
		event = append(event, Event{T[i], 0, i})
	}
	// 時間発生順にソート（出発が後）
	sort.Slice(event, func(i, j int) bool {
		if event[i].time == event[j].time {
			return event[i].flg < event[j].flg
		}
		return event[i].time < event[j].time
	})

	ans := make([]int, M)
	ans[0] = X
	station := make([]int, N+1)
	// 時間順に処理
	for _, e := range event {
		t, f, i := e.time, e.flg, e.idx
		if f == 1 { // 出発の場合
			if i != 0 { // X1でなければ、現在の到着時刻から予定時刻を引いた差分を記録
				ans[i] = max(0, station[A[i]]-t)
			}
		} else { // 到着の場合、到着時刻と到着時刻＋現在の遅れの大きな方を記憶
			station[B[i]] = max(station[B[i]], t+ans[i])
		}
	}
	outSlice(ans[1:])
}
