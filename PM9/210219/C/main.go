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
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	s := getS()

	if len(s) == 1 {
		if s == "8" {
			out("Yes")
		} else {
			out("No")
		}
		return
	}
	if len(s) == 2 {
		a, b := int(s[0]-'0'), int(s[1]-'0')
		if (a*10+b)%8 == 0 {
			out("Yes")
			return
		}
		if (b*10+a)%8 == 0 {
			out("Yes")
			return
		}
		out("No")
		return
	}

	m := make([]int, 10)
	for _, e := range s {
		m[int(e-'0')]++
	}

	for i := 100; i < 999; i++ {
		n := make([]int, 10)
		copy(n, m)
		if i%8 == 0 {
			ok := true
			x := i
			for j := 0; j < 3; j++ {
				if n[x%10] == 0 {
					ok = false
					break
				}
				n[x%10]--
				x /= 10
			}
			if ok {
				out("Yes")
				return
			}
		}
	}
	out("No")
}
