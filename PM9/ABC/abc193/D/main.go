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

func f(s []byte) int {
	m := make([]int, 10)
	for _, e := range s {
		m[int(e-'0')]++
	}
	ret := 0
	for i := 1; i < 10; i++ {
		d := i
		for j := 0; j < m[i]; j++ {
			d *= 10
		}
		ret += d
	}
	return ret
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	K := getI()
	S, T := []byte(getS()), []byte(getS())

	rest := make([]int, 10)
	for i := 1; i < 10; i++ {
		rest[i] = K
	}
	for i := 0; i < 4; i++ {
		rest[S[i]-'0']--
		rest[T[i]-'0']--
	}

	tot := float64(9*K - 8)
	win := 0.0
	lose := 0.0
	for i := 1; i <= 9; i++ {
		for j := 1; j <= 9; j++ {
			if i == j && rest[i] < 2 {
				continue
			}
			if rest[i] < 1 || rest[j] < 1 {
				continue
			}
			x, y := rest[i], rest[j]
			if i == j {
				y = rest[j] - 1
			}

			S[4] = byte(i + '0')
			T[4] = byte(j + '0')
			s, t := f(S), f(T)
			// out(string(S), string(T), s, t, x, y)
			if s > t {
				// out(x, y, float64(x)/tot, float64(y)/tot)
				win += (float64(x) / tot) * (float64(y) / (tot - 1))
			} else {
				lose += (float64(x) / tot) * (float64(y) / (tot - 1))
			}
		}
	}
	out(win)
}
