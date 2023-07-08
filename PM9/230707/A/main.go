package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
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

func f() bool {
	t0 := getS()
	if t0 == "0" {
		return false
	}
	t1 := getS()
	t := t0 + " " + t1
	b := getS()
	// out(t0, t1, b)
	start, _ := time.Parse("2006/01/02 15:04:05", t)
	out(t, start)
	cnt := 0
	for i := 0; i < len(b); i++ {
		cnt += 1 << i
	}
	end := start.Add(time.Second * time.Duration(cnt))
	out(end.Format("2006/01/02 15:04:05"))
	return true
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()

	for {
		tmp1 := getS()
		if tmp1 == "0" {
			break
		}
		tmp2 := strings.Split(tmp1, "/")
		A := tmp2[0]
		a, _ := strconv.Atoi(A)
		B := tmp2[1]
		b, _ := strconv.Atoi(B)
		C := tmp2[2]
		c, _ := strconv.Atoi(C)
		tmp3 := getS()
		tmp4 := strings.Split(tmp3, ":")
		D := tmp4[0]
		d, _ := strconv.Atoi(D)
		E := tmp4[1]
		e, _ := strconv.Atoi(E)
		F := tmp4[2]
		f, _ := strconv.Atoi(F)
		n := getS()
		ti := 0
		for i := 0; i < len(n); i++ {
			ti *= 2
			if n[i] == '1' {
				ti++
			}
		}
		day := ti / 86400
		ti %= 86400
		f += ti % 60
		ti /= 60
		if f >= 60 {
			e++
			f -= 60
		}
		e += ti % 60
		ti /= 60
		if e >= 60 {
			d++
			e -= 60
		}
		d += ti % 24
		if d >= 24 {
			day++
			d -= 24
		}
		for day > 0 {
			day--
			c++
			if c > 31 || ((b == 4 || b == 6 || b == 9 || b == 11) && c > 30) || (b == 2 && c > 29) || (b == 2 && c > 28 && (a%4 != 0 || (a%400 != 0 && a%100 == 0))) {
				b++
				c = 1
			}
			if b > 12 {
				a++
				b = 1
			}
		}
		fmt.Fprintf(wr, "%d/%02d/%02d %02d:%02d:%02d\n", a, b, c, d, e, f)
	}
}
