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

var mod *ModInt

func comb(n, r int) int {
	a, b := 1, 1
	for i := 0; i < r; i++ {
		a = mod.Mul(a, mod.Sub(n, i))
		b = mod.Mul(b, (i + 1))
	}
	return mod.Div(a, b)
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N := getI()
	p := getInts(N - 1)
	c := getInts(N)
	d := getInts(N)

	mod = NewModInt(998244353)

	node := make([][]int, N)
	for i, e := range p {
		node[e-1] = append(node[e-1], i+1)
	}

	ans := 1
	var f func(cur int) int
	f = func(cur int) int {
		num := c[cur]

		for _, e := range node[cur] {
			num += f(e)
		}
		ans = mod.Mul(ans, comb(num, d[cur]))
		num -= d[cur]
		return num
	}

	f(0)
	out(ans)

}

type ModInt struct {
	mod       int
	factMemo  []int
	ifactMemo []int
}

func NewModInt(mod int) *ModInt {
	return &ModInt{
		mod:       mod,
		factMemo:  []int{1, 1},
		ifactMemo: []int{1, 1},
	}
}

func (m *ModInt) Add(a, b int) int {
	ret := (a + b) % m.mod
	if ret < 0 {
		ret += m.mod
	}
	return ret
}

func (m *ModInt) Sub(a, b int) int {
	ret := (a - b) % m.mod
	if ret < 0 {
		ret += m.mod
	}
	return ret
}

func (m *ModInt) Mul(a, b int) int {
	return int(int64(a) * int64(b) % int64(m.mod))
}

func (m *ModInt) Div(a, b int) int {
	return m.Mul(a, m.ModInv(b))
}

func (m *ModInt) Pow(p, n int) int {
	ret := 1
	x := p % m.mod
	for n != 0 {
		if n%2 == 1 {
			ret = m.Mul(ret, x)
		}
		n /= 2
		x = m.Mul(x, x)
	}
	return ret
}

func (m *ModInt) ModInv(a int) int {
	b, u, v := m.mod, 1, 0
	for b != 0 {
		t := a / b
		a -= t * b
		a, b = b, a
		u -= t * v
		u, v = v, u
	}
	u %= m.mod
	if u < 0 {
		u += m.mod
	}
	return u
}

func (m *ModInt) InitComb(limit int) {
	if len(m.factMemo) > limit {
		return
	}
	fact := make([]int, limit+1)
	ifact := make([]int, limit+1)

	fact[0] = 1
	for i := 1; i <= limit; i++ {
		fact[i] = int(int64(fact[i-1]) * int64(i) % int64(m.mod))
	}

	ifact[limit] = m.ModInv(fact[limit])
	for i := limit - 1; i >= 0; i-- {
		ifact[i] = int(int64(ifact[i+1]) * int64(i+1) % int64(m.mod))
	}

	m.factMemo = fact
	m.ifactMemo = ifact
}

func (m *ModInt) Fact(n int) int {
	if len(m.factMemo) <= n {
		m.InitComb(n * 2)
	}
	return m.factMemo[n]
}

func (m *ModInt) IFact(n int) int {
	if len(m.ifactMemo) <= n {
		m.InitComb(n * 2)
	}
	return m.ifactMemo[n]
}

func (m *ModInt) NCr(n, r int) int {
	if n < r || r < 0 {
		return 0
	}
	if len(m.factMemo) <= n {
		m.InitComb(n * 2)
	}
	return m.Mul(m.Fact(n), m.Mul(m.IFact(r), m.IFact(n-r)))
}

func (m *ModInt) PowModMatrix(A [][]int, p int) [][]int {
	N := len(A)
	ret := make([][]int, N)
	for i := 0; i < N; i++ {
		ret[i] = make([]int, N)
		ret[i][i] = 1
	}

	for p > 0 {
		if p&1 == 1 {
			ret = m.MulMod(ret, A)
		}
		A = m.MulMod(A, A)
		p >>= 1
	}
	return ret
}

func (m *ModInt) MulMod(A, B [][]int) [][]int {
	H := len(A)
	W := len(B[0])
	K := len(A[0])

	C := make([][]int, H)
	for i := 0; i < H; i++ {
		C[i] = make([]int, W)
	}

	for i := 0; i < H; i++ {
		for j := 0; j < W; j++ {
			var sum int64
			for k := 0; k < K; k++ {
				sum += int64(A[i][k]) * int64(B[k][j])
			}
			C[i][j] = int(sum % int64(m.mod))
		}
	}
	return C
}
