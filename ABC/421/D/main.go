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
	if len(s) == 0 {
		return
	}
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

func getStrings(N int) []string {
	ret := make([]string, N)
	for i := 0; i < N; i++ {
		ret[i] = getS()
	}
	return ret
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

// 値を圧縮した配列を返す
func compressArray(a []int) []int {
	m := make(map[int]int)
	for _, e := range a {
		m[e] = 1
	}
	b := make([]int, 0)
	for e := range m {
		b = append(b, e)
	}
	sort.Ints(b)
	for i, e := range b {
		m[e] = i
	}

	ret := make([]int, len(a))
	for i, e := range a {
		ret[i] = m[e]
	}
	return ret
}

type Pair struct {
	d byte
	n int
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	rt, ct, ra, ca := getI(), getI(), getI(), getI()
	_, M, L := getI(), getI(), getI()

	m := make(map[int]bool)
	s := make([]Pair, M)
	tot := 0
	m[tot] = true
	for i := 0; i < M; i++ {
		s[i] = Pair{getS()[0], getI()}
		tot += s[i].n
		m[tot] = true
	}
	t := make([]Pair, L)
	tot = 0
	for i := 0; i < L; i++ {
		t[i] = Pair{getS()[0], getI()}
		tot += t[i].n
		m[tot] = true
	}

	// --- 累積時間の計算 ---
	s_cum_times := make([]int, M+1)
	for i := 0; i < M; i++ {
		s_cum_times[i+1] = s_cum_times[i] + s[i].n
	}
	t_cum_times := make([]int, L+1)
	for i := 0; i < L; i++ {
		t_cum_times[i+1] = t_cum_times[i] + t[i].n
	}

	tm := make([]int, 0)
	for e := range m {
		tm = append(tm, e)
	}
	sort.Ints(tm)

	// --- sn, tn の事前生成 ---
	nm := len(tm) - 1
	sn := make([]byte, nm)
	tn := make([]byte, nm)

	s_ptr := 0
	t_ptr := 0

	for i := 0; i < nm; i++ {
		start_time := tm[i]

		for s_cum_times[s_ptr+1] <= start_time {
			s_ptr++
		}
		sn[i] = s[s_ptr].d

		for t_cum_times[t_ptr+1] <= start_time {
			t_ptr++
		}
		tn[i] = t[t_ptr].d
	}

	// --- 衝突判定と座標更新 ---
	dir := map[byte][2]int{
		'U': {-1, 0}, 'D': {1, 0}, 'L': {0, -1}, 'R': {0, 1},
	}
	ans := 0

	for i := 0; i < nm; i++ {
		duration := tm[i+1] - tm[i]
		vec_s := dir[sn[i]]
		vec_t := dir[tn[i]]

		// 相対位置と相対速度を計算
		rel_r, rel_c := ra-rt, ca-ct
		rel_dr, rel_dc := vec_s[0]-vec_t[0], vec_s[1]-vec_t[1]

		// 衝突判定
		if rel_dr == 0 && rel_dc == 0 {
			if rel_r == 0 && rel_c == 0 {
				ans += duration
			}
		} else if rel_dr == 0 {
			if rel_r == 0 && rel_dc != 0 && rel_c%rel_dc == 0 {
				k := rel_c / rel_dc
				if k > 0 && k <= duration {
					ans++
				}
			}
		} else if rel_dc == 0 {
			if rel_c == 0 && rel_dr != 0 && rel_r%rel_dr == 0 {
				k := rel_r / rel_dr
				if k > 0 && k <= duration {
					ans++
				}
			}
		} else {
			if rel_r%rel_dr == 0 && rel_c%rel_dc == 0 {
				k1, k2 := rel_r/rel_dr, rel_c/rel_dc
				if k1 == k2 && k1 > 0 && k1 <= duration {
					ans++
				}
			}
		}

		// 座標更新
		rt += vec_s[0] * duration
		ct += vec_s[1] * duration
		ra += vec_t[0] * duration
		ca += vec_t[1] * duration
	}
	out(ans)
}
