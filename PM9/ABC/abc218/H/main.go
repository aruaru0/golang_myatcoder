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

func sumarr(a []int) int {
	ans := 0
	for _, aa := range a {
		ans += aa
	}
	return ans
}

func solveBaseCase(A []int, l, r int) ([]int, []int, []int, []int) {
	N := r - l + 1
	if N == 3 {
		return []int{}, []int{}, []int{}, []int{A[l+1]}
	}
	if N == 4 {
		return []int{}, []int{A[l+2]}, []int{A[l+1]}, []int{max(A[l+1], A[l+2])}
	}
	rr1 := A[l+2]
	rb1 := max(A[l+2], A[l+3])
	br1 := max(A[l+1], A[l+2])
	bb1 := max(max(A[l+1], A[l+2]), A[l+3])
	bb2 := A[l+1] + A[l+3] - bb1
	return []int{rr1}, []int{rb1}, []int{br1}, []int{bb1, bb2}
}

func merge(a1, a2, a3 []int, N int) []int {
	best := make([]int, N+1)
	inc := make([]int, N)
	for i := 0; i <= N; i++ {
		best[i] = max(max(a1[i], a2[i]), a3[i])
	}
	for i := 0; i < N; i++ {
		inc[i] = best[i+1] - best[i]
	}
	return inc
}

func combine(a1, a2 []int, starter, N int) []int {
	inf := 1_000_000_000_000_000_000
	res := make([]int, N)
	for i := 0; i < N; i++ {
		res[i] = -inf
	}
	cnt, i1, i2, l1, l2 := 0, 0, 0, len(a1), len(a2)
	if starter == 0 {
		res[0] = 0
	} else {
		res[1], cnt = starter, 1
	}
	for cnt+1 < N {
		if i1 < l1 && (i2 == l2 || a1[i1] >= a2[i2]) {
			cnt++
			res[cnt] = res[cnt-1] + a1[i1]
			i1++
		} else if i2 < l2 {
			cnt++
			res[cnt] = res[cnt-1] + a2[i2]
			i2++
		} else {
			break
		}
	}
	return res
}

func solveit(A []int, l, r int) ([]int, []int, []int, []int) {
	if r-l+1 <= 5 {
		return solveBaseCase(A, l, r)
	}
	ressize := (r - l + 1 + 1) / 2
	m := (r + l) >> 1
	leftrr, leftrb, leftbr, leftbb := solveit(A, l, m)
	rightrr, rightrb, rightbr, rightbb := solveit(A, m+1, r)
	exp1a := combine(leftrr, rightbr, A[m], ressize+1)
	exp1b := combine(leftrb, rightrr, A[m+1], ressize+1)
	exp1c := combine(leftrb, rightbr, 0, ressize+1)
	rr1 := merge(exp1a, exp1b, exp1c, ressize)
	exp2a := combine(leftrr, rightbb, A[m], ressize+1)
	exp2b := combine(leftrb, rightbb, 0, ressize+1)
	exp2c := combine(leftrb, rightrb, A[m+1], ressize+1)
	rr2 := merge(exp2a, exp2b, exp2c, ressize)
	exp3a := combine(leftbb, rightrr, A[m+1], ressize+1)
	exp3b := combine(leftbr, rightbr, A[m], ressize+1)
	exp3c := combine(leftbb, rightbr, 0, ressize+1)
	rr3 := merge(exp3a, exp3b, exp3c, ressize)
	exp4a := combine(leftbb, rightbb, 0, ressize+1)
	exp4b := combine(leftbr, rightbb, A[m], ressize+1)
	exp4c := combine(leftbb, rightrb, A[m+1], ressize+1)
	rr4 := merge(exp4a, exp4b, exp4c, ressize)
	return rr1, rr2, rr3, rr4
}

// お手上げ！！
func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N, R := getI(), getI()
	A := getInts(N - 1)
	R = min(R, N-R)
	AA := make([]int, N)
	AA[0] = A[0]
	for i := 1; i < N-1; i++ {
		AA[i] = A[i] + A[i-1]
	}
	AA[N-1] = A[N-2]
	if N == 2 {
		fmt.Println(A[0])
		return
	}
	rr, rb, br, bb := solveit(AA, 0, N-1)
	cand1 := 0
	if R > 1 {
		cand1 = AA[0] + AA[N-1] + sumarr(rr[:R-2])
	}
	cand2 := AA[0] + sumarr(rb[:R-1])
	cand3 := AA[N-1] + sumarr(br[:R-1])
	cand4 := sumarr(bb[:R])
	ans := max(max(max(cand1, cand2), cand3), cand4)
	fmt.Println(ans)
}
