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

type BIT struct {
	v []int
}

func newBIT(n int) *BIT {
	return &BIT{v: make([]int, n)}
}

func (b *BIT) sum(a int) int {
	ret := 0
	for i := a + 1; i > 0; i -= i & -i {
		ret += b.v[i-1]
	}
	return ret
}

func (b *BIT) rangeSum(x, y int) int {
	if y == 0 {
		return 0
	}
	y--
	if x == 0 {
		return b.sum(y)
	} else {
		return b.sum(y) - b.sum(x-1)
	}
}

func (b *BIT) add(a, w int) {
	n := len(b.v)
	for i := a + 1; i <= n; i += i & -i {
		b.v[i-1] += w
	}
}

func (b *BIT) lowerBound(k int) int {
	if k <= 0 {
		return 0
	}
	x := 0
	n := len(b.v)
	w := 1
	for w*2 <= n {
		w *= 2
	}
	for w > 0 {
		if x+w <= n && b.v[x+w-1] < k {
			k -= b.v[x+w-1]
			x += w
		}
		w /= 2
	}
	return x
}

type Query struct {
	idx int // 元のクエリのインデックス
	k   int // そのブロックの中で何番目か (1-indexed)
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)

	N, M := getI(), getI()
	A := make([]int, N)
	C := make([]int, M+1)
	for i := 0; i < N; i++ {
		A[i] = getI()
		C[A[i]]++
	}

	C_max := 0
	for i := 1; i <= M; i++ {
		C_max = max(C_max, C[i])
	}

	// 各出現回数の要素数をカウント
	countC := make([]int, C_max+1)
	for i := 1; i <= M; i++ {
		countC[C[i]]++
	}

	sizes := make([]int, C_max)
	acc := make([]int, C_max)
	currentSize := 0
	sumAcc := 0
	for c := 0; c < C_max; c++ {
		currentSize += countC[c]
		sizes[c] = currentSize
		sumAcc += sizes[c]
		acc[c] = sumAcc
	}

	Q := getI()
	ans := make([]int, Q)
	queries := make([][]Query, C_max)

	for q := 0; q < Q; q++ {
		X := getI()
		if X <= N {
			ans[q] = A[X-1]
		} else if C_max == 0 || X > N+acc[C_max-1] {
			rem := X - N
			if C_max > 0 {
				rem -= acc[C_max-1]
			}
			ans[q] = (rem-1)%M + 1
		} else {
			X_rem := X - N
			l, r := 0, C_max-1
			for l < r {
				mid := (l + r) / 2
				if acc[mid] >= X_rem {
					r = mid
				} else {
					l = mid + 1
				}
			}
			c := l
			k := X_rem
			if c > 0 {
				k -= acc[c-1]
			}
			queries[c] = append(queries[c], Query{idx: q, k: k})
		}
	}

	bit := newBIT(M + 1)

	cToI := make([][]int, C_max)
	for i := 1; i <= M; i++ {
		if C[i] < C_max {
			cToI[C[i]] = append(cToI[C[i]], i)
		}
	}

	for c := 0; c < C_max; c++ {
		for _, i := range cToI[c] {
			bit.add(i, 1)
		}
		for _, q := range queries[c] {
			ans[q.idx] = bit.lowerBound(q.k)
		}
	}

	for i := 0; i < Q; i++ {
		out(ans[i])
	}
}
