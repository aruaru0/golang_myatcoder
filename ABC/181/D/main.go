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

type pair struct {
	v [10]int
}

func keta2(s string) {
	v, _ := strconv.Atoi(s)
	if v%8 == 0 {
		out("Yes")
		return
	}
	x := string(s[1]) + string(s[0])
	v, _ = strconv.Atoi(x)
	if v%8 == 0 {
		out("Yes")
		return
	}
	out("No")
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1000000)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()

	s := getS()
	if len(s) == 1 {
		if s == "8" {
			out("Yes")
			return
		}
		out("No")
		return
	}
	if len(s) == 2 {
		keta2(s)
		return
	}

	a := make([]pair, 0)
	for i := 100; i < 1000; i++ {
		if i%8 == 0 {
			n := i
			var x [10]int
			for j := 0; j < 3; j++ {
				x[n%10]++
				n /= 10
			}
			a = append(a, pair{x})
		}
	}

	m := make(map[int]int)
	for i := 0; i < len(s); i++ {
		m[int(s[i]-'0')]++
	}

	for _, e := range a {
		flg := true
		for i, v := range e.v {
			if m[i] < v {
				flg = false
				break
			}
		}
		if flg {
			out("Yes")
			return
		}
	}
	out("No")
}
