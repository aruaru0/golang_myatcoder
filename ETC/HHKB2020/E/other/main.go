package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func out(x ...interface{}) {
	fmt.Println(x...)
}

var sc = bufio.NewScanner(os.Stdin)

func getInt() int {
	sc.Scan()
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return i
}

func getInts(N int) []int {
	ret := make([]int, N)
	for i := 0; i < N; i++ {
		ret[i] = getInt()
	}
	return ret
}

func getString() string {
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

const mod = 1000000007

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1000000)
	H, W := getInt(), getInt()
	s := make([]string, H)
	a := make([][]int, H)
	for i := 0; i < H; i++ {
		s[i] = getString()
		a[i] = make([]int, W)
	}

	k := 0
	for y := 0; y < H; y++ {
		from := 0
		for from < W {
			cnt := 0
			to := from
			for to < W {
				if s[y][to] == '#' {
					break
				}
				cnt++
				to++
				k++
			}
			for i := from; i < to; i++ {
				a[y][i] += cnt
			}
			from = to + 1
		}
	}

	for x := 0; x < W; x++ {
		from := 0
		for from < H {
			cnt := 0
			to := from
			for to < H {
				if s[to][x] == '#' {
					break
				}
				cnt++
				to++
			}
			for i := from; i < to; i++ {
				a[i][x] += cnt - 1
			}
			from = to + 1
		}
	}

	m := make(map[int]int)
	for y := 0; y < H; y++ {
		for x := 0; x < W; x++ {
			if a[y][x] == 0 {
				continue
			}
			m[a[y][x]]++
		}
	}

	// out(m)
	md := newModint(mod)
	ans := 0
	for t, n := range m {
		t = md.mul(md.sub(md.pow(2, t), 1), md.pow(2, k-t))
		ans = md.add(ans, md.mul(t, n))
	}

	out(ans)
}

//----------------------------------------
// modint
//----------------------------------------
type modint struct {
	mod int
}

func newModint(m int) *modint {
	var ret modint
	ret.mod = m
	return &ret
}

func (m *modint) add(a, b int) int {
	ret := (a + b) % m.mod
	if ret < 0 {
		ret += m.mod
	}
	return ret
}

func (m *modint) sub(a, b int) int {
	ret := (a - b) % m.mod
	if ret < 0 {
		ret += m.mod
	}
	return ret
}

func (m *modint) mul(a, b int) int {
	ret := a * b % m.mod
	if ret < 0 {
		ret += m.mod
	}
	return ret
}

func (m *modint) div(a, b int) int {
	ret := a * m.modinv(b)
	ret %= m.mod
	return ret
}

func (m *modint) pow(p, n int) int {
	ret := 1
	x := p
	for n != 0 {
		if n%2 == 1 {
			ret *= x
			ret %= m.mod
		}
		n /= 2
		x = x * x % m.mod
	}
	return ret
}

// 逆元を使った割り算（MOD）
// mod. m での a の逆元 a^{-1} を計算する
func (m *modint) modinv(a int) int {
	b := m.mod
	u := 1
	v := 0
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
