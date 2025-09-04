package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
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

func chmaxf(x *float64, y float64) {
	if *x < y {
		*x = y
	}
}

// vector<int> をmapのキーとして使用するためのカスタム型とメソッド
type key []int

func (k key) String() string {
	// スライスをソートして正規化する
	sort.Ints(k)

	// 無限再帰を防ぐため、strings.Builderを使用して手動で文字列を構築する
	var sb strings.Builder
	sb.WriteString("[")
	for i, v := range k {
		if i > 0 {
			sb.WriteString(" ")
		}
		sb.WriteString(fmt.Sprintf("%d", v))
	}
	sb.WriteString("]")
	return sb.String()
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	n := 6
	a := getInts(n)

	memo := make([]map[string]float64, 3)
	for i := 0; i < 3; i++ {
		memo[i] = make(map[string]float64)
	}

	var f func(rem int, keep []int) float64
	f = func(rem int, keep []int) float64 {
		if rem == 0 {
			var res float64
			for i := 0; i < n; i++ {
				cnt := 0
				for _, x := range keep {
					if x == a[i] {
						cnt++
					}
				}
				chmaxf(&res, float64(cnt*a[i]))
			}
			return res
		}

		rem--
		keepKey := key(keep).String()
		if val, ok := memo[rem][keepKey]; ok {
			return val
		}

		var g func(num int, dice []int) float64
		g = func(num int, dice []int) float64 {
			if num == 0 {
				m := len(dice)
				var res float64
				for s := 0; s < 1<<m; s++ {
					kp := make([]int, len(keep))
					copy(kp, keep)
					for i := 0; i < m; i++ {
						if s>>i&1 == 1 {
							kp = append(kp, dice[i])
						}
					}
					chmaxf(&res, f(rem, kp))
				}
				return res
			}

			num--
			var res float64
			for i := 0; i < n; i++ {
				res += g(num, append(dice, a[i]))
			}
			return res / float64(n)
		}

		res := g(5-len(keep), []int{})
		memo[rem][keepKey] = res
		return res
	}

	ans := f(3, []int{})
	fmt.Printf("%.10f\n", ans)
}
