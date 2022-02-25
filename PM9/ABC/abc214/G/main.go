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

const mod = int(1e9) + 7

func f(n int) []int {
	res := make([]int, n+1)
	if n == 1 {
		res[0] = 1
		res[1] = 1
		return res
	}
	for si := 0; si < 2; si++ {
		dp := make([][2]int, 1)
		dp[0][si] = 1
		for i := 0; i < n; i++ {
			p := make([][2]int, i+2)
			dp, p = p, dp
			for j := 0; j < len(p); j++ {
				for k := 0; k < 2; k++ {
					for x := 0; x < 3; x++ {
						nk := 0
						if x == 2 {
							nk = 1
						}
						if k == 1 && x == 1 {
							continue
						}
						nj := j
						if x != 0 {
							nj++
						}
						dp[nj][nk] += p[j][k]
						dp[nj][nk] %= mod
					}
				}
			}
		}
		for i := 0; i <= n; i++ {
			res[i] += dp[i][si]
			res[i] %= mod
		}
	}
	return res
}

// 公式解答をgoに変換しただけ
// 方針はわかるが、コードはいまいち理解できず
func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	n := getI()
	p := make([]int, n)
	q := make([]int, n)
	for i := 0; i < n; i++ {
		p[i] = getI() - 1
	}
	for i := 0; i < n; i++ {
		q[i] = getI() - 1
	}

	tmp := make([]int, n)
	copy(tmp, q)

	for i := 0; i < n; i++ {
		q[p[i]] = tmp[i]
	}
	used := make([]bool, n)

	ss := make([]int, 0)
	for i := 0; i < n; i++ {
		if used[i] {
			continue
		}
		v := i
		cnt := 0
		for used[v] == false {
			used[v] = true
			v = q[v]
			cnt++
		}
		ss = append(ss, cnt)
	}

	dp := make([]int, 1)
	dp[0] = 1
	for _, s := range ss {
		d := f(s)
		p := make([]int, len(dp)+len(d)-1)
		p, dp = dp, p
		for i := 0; i < len(p); i++ {
			for j := 0; j < len(d); j++ {
				dp[i+j] += p[i] * d[j]
				dp[i+j] %= mod
			}
		}
	}

	ans := 0
	fact := make([]int, n+1)
	fact[0] = 1
	for i := 0; i < n; i++ {
		fact[i+1] = fact[i] * (i + 1) % mod
	}
	for i := 0; i <= n; i++ {
		now := dp[i] * fact[n-i] % mod
		if i%2 != 0 {
			ans -= now
			ans %= mod
			if ans < 0 {
				ans += mod
			}
		} else {
			ans += now
			ans %= mod
		}
	}

	out(ans)
}
