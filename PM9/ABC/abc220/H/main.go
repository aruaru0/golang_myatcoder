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

// コード書いてもいまいち理解できない。
// 現状スキルでは厳しい感じ
// 半分全列挙で行けるとは分かったが、実装ができなかった。
func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	n, m := getI(), getI()
	l := n / 2
	r := n - l
	// ビット表現で接続を管理する
	gl := make([]int, l)
	gr := make([]int, r)
	g := make([]int, l)
	for i := 0; i < m; i++ {
		a, b := getI()-1, getI()-1
		if a > b { // a < bになるようにする
			a, b = b, a
		}
		if b < l { // bがl以下の部分（a,b < l)
			gl[a] |= 1 << b
			gl[b] |= 1 << a
		} else if l <= a { // a,b >= l
			a -= l
			b -= l
			gr[a] |= 1 << b
			gr[b] |= 1 << a
		} else { // 2つを結ぶ部分
			b -= l
			g[a] |= 1 << b
		}
	}

	// 小さいほうのブロックの処理
	dl := make([][]int, 2)
	for i := 0; i < 2; i++ {
		dl[i] = make([]int, 1<<r)
	}
	for s := 0; s < 1<<l; s++ {
		deg, one := 0, 0
		x := 0
		for i := 0; i < l; i++ {
			if (s>>i)%2 != 0 {
				deg += bits.OnesCount(uint(gl[i]))
				deg += bits.OnesCount(uint(g[i]))
				one += bits.OnesCount(uint(gl[i] & s))
			} else {
				x ^= g[i]
			}
		}
		e := (deg - one/2) % 2
		dl[e][x]++
	}
	// 大きいほうのブロックの処理
	dr := make([][]int, 2)
	for i := 0; i < 2; i++ {
		dr[i] = make([]int, 1<<r)
	}
	for s := 0; s < 1<<r; s++ {
		deg, one := 0, 0
		for i := 0; i < r; i++ {
			if (s>>i)%2 != 0 {
				deg += bits.OnesCount(uint(gr[i]))
				one += bits.OnesCount(uint(gr[i] & s))
			}
		}
		e := (deg - one/2) % 2
		dr[e][s]++
	}

	// ２つのブロックの間の処理
	calc := func(x, y []int) []int {
		a := make([]int, 1<<r)
		b := make([]int, 1<<r)
		copy(a, x)
		copy(b, y)
		for i := 0; i < r; i++ {
			for j := 0; j < 1<<r; j++ {
				if (j>>i)%2 != 0 {
					a[j^(1<<i)] += a[j]
				}
			}
		}
		for i := 0; i < r; i++ {
			for j := 0; j < 1<<r; j++ {
				if (j>>i)%2 != 0 {
					b[j^(1<<i)] += b[j]
				}
			}
		}
		for i := 0; i < 1<<r; i++ {
			a[i] *= b[i]
		}
		for i := 0; i < r; i++ {
			for j := 0; j < 1<<r; j++ {
				if (j>>i)%2 != 0 {
					a[j^(1<<i)] -= a[j]
				}
			}
		}
		return a
	}

	ans := 0
	for el := 0; el < 2; el++ {
		for er := 0; er < 2; er++ {
			d := calc(dl[el], dr[er])
			for i := 0; i < 1<<r; i++ {
				if (bits.OnesCount(uint(i))+el+er)%2 == 0 {
					ans += d[i]
				}
			}
		}
	}
	out(ans)
}

// gl:0  gr:2 1  g:1
// dl
// 0 1 0 0
// 1 0 0 0
// dr
// 1 0 0 0
// 0 1 1 1
// 6
// 1 0 0 0
// 1 2 0 0
// 1 0 0 0
// 3 0 0 0
// 6
