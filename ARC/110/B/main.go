package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
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
	sc.Buffer([]byte{}, 1000000)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N := getI()
	T := getS()
	// N := len(T)
	// 110 10 0 ... 1 11 110 pat
	if T == "1" {
		out(int(1e10 * 2))
		return
	}

	st := 0
	if strings.HasPrefix(T, "10") {
		st = 1
	}
	if strings.HasPrefix(T, "0") {
		st = 2
	}

	flg := 0
	cnt := 0
	if st != 0 {
		flg++
		cnt--
	}
	for i := 0; i < N; i++ {
		if st == 0 {
			if T[i] != '1' {
				out(0)
				return
			}
			st = 1
			continue
		}
		if st == 1 {
			if T[i] != '1' {
				out(0)
				return
			}
			st = 2
			continue
		}
		if st == 2 {
			if T[i] != '0' {
				out(0)
				return
			}
			st = 0
			cnt++
		}
	}
	if st != 0 {
		flg++
	}
	// out("ok", st, cnt, flg)
	ans := int(1e10) - (cnt - 1) - flg
	out(ans)
}
