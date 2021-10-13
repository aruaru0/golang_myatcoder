package main

import (
	"bufio"
	"fmt"
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

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer(make([]byte, 10000), 1001001)

	n := getI()
	s := getInts(n)
	t := getInts(n)
	u := make([]uint64, n)
	v := make([]uint64, n)
	for i := 0; i < n; i++ {
		u[i], _ = strconv.ParseUint(getS(), 10, 64)
	}
	for i := 0; i < n; i++ {
		v[i], _ = strconv.ParseUint(getS(), 10, 64)
	}

	a := make([][]uint64, n)
	b := make([][]bool, n)
	ans := make([][]uint64, n)
	for i := 0; i < n; i++ {
		a[i] = make([]uint64, n)
		b[i] = make([]bool, n)
		ans[i] = make([]uint64, n)
	}

	r := make([][2]int, n)
	c := make([][2]int, n)

	for k := 0; k < 64; k++ {
		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				a[i][j] = 0
				b[i][j] = false
			}
		}

		for i := 0; i < n; i++ {
			for j := 0; j < 2; j++ {
				r[i][j] = 0
				c[i][j] = 0
			}
		}

		for i := 0; i < n; i++ {
			f := u[i] & 1
			if s[i] == int(f) {
				continue
			}
			for j := 0; j < n; j++ {
				a[i][j] = f
				b[i][j] = true
			}
		}

		for i := 0; i < n; i++ {
			f := v[i] & 1
			if t[i] == int(f) {
				continue
			}
			for j := 0; j < n; j++ {
				if b[j][i] && a[j][i] != f {
					fmt.Fprintln(wr, -1)
					return
				}
				a[j][i] = f
			}
		}

		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				r[i][a[i][j]]++
				c[j][a[i][j]]++
			}
		}

		for i := 0; i < n; i++ {
			if s[i] == 1 && u[i]&1 == 1 {
				if r[i][1] > 0 {
					continue
				}
				for j := 0; j < n; j++ {
					if t[j] == 0 && v[j]&1 == 0 && c[j][0] > 1 || t[j] == 1 && v[j]&1 == 1 {
						a[i][j] = 1
						r[i][0]--
						r[i][1]++
						c[j][0]--
						c[j][1]++
						break
					}
					if j == n-1 {
						fmt.Fprintln(wr, -1)
						return
					}
				}
			}
		}

		for i := 0; i < n; i++ {
			if t[i] == 1 && v[i]&1 == 1 {
				if c[i][1] > 0 {
					continue
				}
				for j := 0; j < n; j++ {
					if s[j] == 0 && u[j]&1 == 0 && r[j][0] > 1 || s[j] == 1 && u[j]&1 == 1 {
						a[j][i] = 1
						r[j][0]--
						r[j][1]++
						c[i][0]--
						c[i][1]++
						break
					}
					if j == n-1 {
						fmt.Fprintln(wr, -1)
						return
					}
				}
			}
		}

		for i := 0; i < n; i++ {
			d := u[i] & 1
			if s[i] == 0 && (d == 0 && r[i][0] == 0 || d == 1 && r[i][1] != n) {
				fmt.Fprintln(wr, -1)
				return
			}
			if s[i] == 1 && (d == 0 && r[i][0] != n || d == 1 && r[i][1] == 0) {
				fmt.Fprintln(wr, -1)
				return
			}

			f := v[i] & 1
			if t[i] == 0 && (f == 0 && c[i][0] == 0 || f == 1 && c[i][1] != n) {
				fmt.Fprintln(wr, -1)
				return
			}
			if t[i] == 1 && (f == 0 && c[i][0] != n || f == 1 && c[i][1] == 0) {
				fmt.Fprintln(wr, -1)
				return
			}
		}

		for i := 0; i < n; i++ {
			for j := 0; j < n; j++ {
				ans[i][j] |= a[i][j] << k
			}
		}

		for i := 0; i < n; i++ {
			u[i] >>= 1
			v[i] >>= 1
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Fprint(wr, ans[i][j], " ")
		}
		fmt.Fprintln(wr)
	}
}
