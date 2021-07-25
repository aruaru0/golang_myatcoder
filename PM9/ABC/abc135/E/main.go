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

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func lcm(a, b int) int {
	return a / gcd(a, b) * b
}

func calc2(X, Y, K int) [][2]int {
	// X , Y >= 0,  X > Y
	if (X+Y)%2 == 1 && K%2 == 0 {
		return nil
	}
	n := (X + Y + K - 1) / K

	if X+Y == K {
		return [][2]int{{X, Y}}
	}
	if n <= 1 {
		n = 2
	}

	if (X+Y)%2 != n*K%2 {
		n++
	}
	ans := make([][2]int, n)
	if n == 3 && X < K {
		ans = make([][2]int, 3)
		ans[0][0] = X
		ans[0][1] = -K + X

		overRight := (K + X - Y) / 2

		ans[1][0] = X + overRight
		ans[1][1] = Y - (K - overRight)

		ans[2][0] = X
		ans[2][1] = Y
		return ans
	} else {
		over := (K*n - (X + Y)) / 2

		for i := 0; i < n; i++ {
			dist := (i + 1) * K
			if dist <= Y+over {
				ans[i][0] = 0
				ans[i][1] = dist
			} else if dist <= Y+over+X {
				ans[i][0] = dist - Y - over
				ans[i][1] = Y + over
			} else {
				ans[i][0] = X
				ans[i][1] = Y + ((n - (i + 1)) * K)
			}
		}
	}
	return ans
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	K := getI()
	X, Y := getI(), getI()

	signX, signY := false, false
	swap := false
	if X < 0 {
		signX = true
		X = -X
	}
	if Y < 0 {
		signY = true
		Y = -Y
	}
	if X < Y {
		X, Y = Y, X
		swap = true
	}

	ans := calc2(X, Y, K)

	if ans == nil {
		out(-1)
		return
	}
	if swap {
		for i := 0; i < len(ans); i++ {
			ans[i][0], ans[i][1] = ans[i][1], ans[i][0]
		}
	}
	if signX {
		for i := 0; i < len(ans); i++ {
			ans[i][0] = -ans[i][0]
		}
	}
	if signY {
		for i := 0; i < len(ans); i++ {
			ans[i][1] = -ans[i][1]
		}
	}

	out(len(ans))
	for _, e := range ans {
		out(e[0], e[1])
	}

}
