package main

import (
	"bufio"
	"fmt"
	"math"
	"math/bits"
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

// Mod constants.
const (
	Mod1000000007 = 1000000007
	Mod998244353  = 998244353
)

var (
	mod  Mint
	fmod func(Mint) Mint
)

// Mint treats the modular arithmetic
type Mint int64

// SetMod sets the mod. It must be called first.
func SetMod(newmod Mint) {
	switch newmod {
	case Mod1000000007:
		fmod = staticMod1000000007
	case Mod998244353:
		fmod = staticMod998244353
	default:
		mod = newmod
		fmod = dynamicMod
	}
}
func dynamicMod(m Mint) Mint {
	m %= mod
	if m < 0 {
		return m + mod
	}
	return m
}
func staticMod1000000007(m Mint) Mint {
	m %= Mod1000000007
	if m < 0 {
		return m + Mod1000000007
	}
	return m
}
func staticMod998244353(m Mint) Mint {
	m %= Mod998244353
	if m < 0 {
		return m + Mod998244353
	}
	return m
}

// Mod returns m % mod.
func (m Mint) Mod() Mint {
	return fmod(m)
}

// Inv returns modular multiplicative inverse
func (m Mint) Inv() Mint {
	return m.Pow(Mint(0).Sub(2))
}

// Pow returns m^n
func (m Mint) Pow(n Mint) Mint {
	p := Mint(1)
	for n > 0 {
		if n&1 == 1 {
			p.MulAs(m)
		}
		m.MulAs(m)
		n >>= 1
	}
	return p
}

// Add returns m+x
func (m Mint) Add(x Mint) Mint {
	return (m + x).Mod()
}

// Sub returns m-x
func (m Mint) Sub(x Mint) Mint {
	return (m - x).Mod()
}

// Mul returns m*x
func (m Mint) Mul(x Mint) Mint {
	return (m * x).Mod()
}

// Div returns m/x
func (m Mint) Div(x Mint) Mint {
	return m.Mul(x.Inv())
}

// AddAs assigns *m + x to *m and returns m
func (m *Mint) AddAs(x Mint) *Mint {
	*m = m.Add(x)
	return m
}

// SubAs assigns *m - x to *m and returns m
func (m *Mint) SubAs(x Mint) *Mint {
	*m = m.Sub(x)
	return m
}

// MulAs assigns *m * x to *m and returns m
func (m *Mint) MulAs(x Mint) *Mint {
	*m = m.Mul(x)
	return m
}

// DivAs assigns *m / x to *m and returns m
func (m *Mint) DivAs(x Mint) *Mint {
	*m = m.Div(x)
	return m
}
func pow(a, b, mod int64) int64 {
	p := int64(1)
	for b > 0 {
		if b&1 == 1 {
			p *= a
			p %= mod
		}
		a *= a
		a %= mod
		b >>= 1
	}
	return p
}
func modOperation(mod int64) func(int64) int64 {
	switch mod {
	case 2:
		return func(a int64) int64 {
			a %= 2
			if a < 0 {
				a += 2
			}
			return a
		}
	case 167772161:
		return func(a int64) int64 {
			a %= 167772161
			if a < 0 {
				a += 167772161
			}
			return a
		}
	case 469762049:
		return func(a int64) int64 {
			a %= 469762049
			if a < 0 {
				a += 469762049
			}
			return a
		}
	case 998244353:
		return func(a int64) int64 {
			a %= 998244353
			if a < 0 {
				a += 998244353
			}
			return a
		}
	case 754974721:
		return func(a int64) int64 {
			a %= 754974721
			if a < 0 {
				a += 754974721
			}
			return a
		}
	}
	return func(a int64) int64 {
		a %= mod
		if a < 0 {
			a += mod
		}
		return a
	}
}
func primitiveRoot(mod int64) int64 {
	switch mod {
	case 2:
		return 1
	case 167772161, 469762049, 998244353:
		return 3
	case 754974721:
		return 11
	}
	var divs [20]int64
	divs[0] = 2
	cnt := 1
	x := (mod - 1) / 2
	for x&1 == 0 {
		x >>= 1
	}
	for i := int64(3); i*i < x+1; i += 2 {
		if x%i == 0 {
			divs[cnt] = i
			cnt++
			for x%i == 0 {
				x /= i
			}
		}
	}
	if x > 1 {
		divs[cnt] = x
		cnt++
	}
	for g := int64(2); ; g++ {
		ok := true
		for i := 0; i < cnt; i++ {
			powModConstExpr := func() int64 {
				if mod == 1 {
					return 0
				}
				r := int64(1)
				y := x % mod
				n := (mod - 1) / divs[i]
				for n > 0 {
					if n&1 == 1 {
						r = (r * y) % mod
					}
					y = (y * y) % mod
					n >>= 1
				}
				return r
			}
			if powModConstExpr() == 1 {
				ok = false
				break
			}
		}
		if ok {
			return g
		}
	}
}
func butterflyF(mod int64) func([]int64) {
	g := primitiveRoot(mod)
	f := modOperation(mod)
	first := true
	var sumE [30]int64
	return func(a []int64) {
		n := len(a)
		h := 0
		for i := 0; 1<<uint(i) < n; i++ {
			h++
		}
		if first {
			first = false
			var es, ies [30]int64
			cnt2 := bits.TrailingZeros64(uint64(mod - 1))
			e := pow(g, (mod-1)>>cnt2, mod)
			ie := pow(e, mod-2, mod)
			for i := cnt2; i >= 2; i-- {
				es[i-2] = e
				ies[i-2] = ie
				e = f(e * e)
				ie = f(ie * ie)
			}
			now := int64(1)
			for i := 0; i < cnt2-2; i++ {
				sumE[i] = f(es[i] * now)
				now = f(now * ies[i])
			}
		}
		for ph := 1; ph <= h; ph++ {
			w := 1 << uint(ph-1)
			p := 1 << uint(h-ph)
			now := int64(1)
			for s := 0; s < w; s++ {
				offset := s << uint(h-ph+1)
				for i := 0; i < p; i++ {
					l := a[i+offset]
					r := f(a[i+offset+p] * now)
					a[i+offset] = f(l + r)
					a[i+offset+p] = f(l - r)
				}
				now = f(now * sumE[bits.TrailingZeros32(^uint32(s))])
			}
		}
	}
}
func butterflyInvF(mod int64) func([]int64) {
	g := primitiveRoot(mod)
	f := modOperation(mod)
	first := true
	var sumIE [30]int64

	return func(a []int64) {
		n := len(a)
		h := 0
		for i := 0; 1<<uint(i) < n; i++ {
			h++
		}
		if first {
			first = false
			var es, ies [30]int64
			cnt2 := bits.TrailingZeros64(uint64(mod - 1))
			e := pow(g, (mod-1)>>cnt2, mod)
			ie := pow(e, mod-2, mod)
			for i := cnt2; i >= 2; i-- {
				es[i-2] = e
				ies[i-2] = ie
				e = f(e * e)
				ie = f(ie * ie)
			}
			now := int64(1)
			for i := 0; i < cnt2-2; i++ {
				sumIE[i] = f(ies[i] * now)
				now = f(now * es[i])
			}
		}
		for ph := h; ph >= 1; ph-- {
			w := 1 << uint(ph-1)
			p := 1 << uint(h-ph)
			inow := int64(1)
			for s := 0; s < w; s++ {
				offset := s << uint(h-ph+1)
				for i := 0; i < p; i++ {
					l := a[i+offset]
					r := a[i+offset+p]
					a[i+offset] = f(l + r)
					a[i+offset+p] = f((l - r) * inow)
				}
				inow = f(inow * sumIE[bits.TrailingZeros32(^uint32(s))])
			}
		}
	}
}
func Convolution(a, b []int64, mod int64) []int64 {
	f := modOperation(mod)
	n := len(a)
	m := len(b)
	if n == 0 || m == 0 {
		return []int64{}
	}
	if n < m {
		n, m = m, n
		a, b = b, a
	}
	if m <= 60 {
		ans := make([]int64, n+m-1)
		for i := 0; i < n; i++ {
			for j := 0; j < m; j++ {
				ans[i+j] = f(ans[i+j] + f(a[i]*b[j]))
			}
		}
		return ans
	}
	z := 1
	for z < n+m-1 {
		z <<= 1
	}
	aa := make([]int64, z)
	bb := make([]int64, z)
	for i := 0; i < n; i++ {
		aa[i] = a[i]
	}
	for i := 0; i < m; i++ {
		bb[i] = b[i]
	}
	butterfly := butterflyF(mod)
	butterfly(aa)
	butterfly(bb)
	for i := 0; i < z; i++ {
		aa[i] = f(aa[i] * bb[i])
	}
	butterflyInv := butterflyInvF(mod)
	butterflyInv(aa)
	iz := pow(int64(z), mod-2, mod)
	for i := 0; i < n+m-1; i++ {
		aa[i] = f(aa[i] * iz)
	}
	return aa[:n+m-1]
}

type pair struct {
	f, s int64
}

var es []pair
var ps [][]int64
var dp [][]int64

const modx = 998244353

func dfs(l, r int) {
	if r-l <= 1 {
		return
	}
	c := (l + r) / 2

	dfs(l, c)
	for ei := 0; ei < len(es); ei++ {
		a, b := es[ei].f, es[ei].s
		for ri := 0; ri < 2; ri++ {
			x := dp[a][l:c]
			y := ps[ei][:r-l-1]
			z := Convolution(x, y, modx)
			for i := c; i < r; i++ {
				dp[b][i] += z[i-(l+1)]
				dp[b][i] %= modx
			}
			a, b = b, a
		}
	}
	dfs(c, r)
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	n, m, t := getI(), getI(), getI()

	es = make([]pair, 0)
	ps = make([][]int64, 0)
	for i := 0; i < m; i++ {
		a, b := int64(getI()-1), int64(getI()-1)
		es = append(es, pair{a, b})
		p := make([]int64, t)
		for j := 0; j < t; j++ {
			x := int64(getI())
			p[j] = x
		}
		ps = append(ps, p)
	}

	dp = make([][]int64, n)
	for i := 0; i < n; i++ {
		dp[i] = make([]int64, t+1)
	}
	dp[0][0] = 1

	dfs(0, t+1)

	out(dp[0][t])
}
