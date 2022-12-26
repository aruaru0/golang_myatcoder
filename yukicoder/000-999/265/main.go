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

func f(s string, flg bool) ([]int, int) {
	// out(s)
	ret := make([]int, d+1)
	x := 0
	a := 0
	pos := 0
	for i := 0; i < len(s); i++ {
		switch s[i] {
		case '+':
			if x != 0 && a == 0 {
				a++
			}
			ret[x] += a
			x = 0
			a = 0
		case 'x':
			x++
		case '*':
		case 'd':
			r, skip := f(s[i+1:], false)
			for j := 1; j < len(r); j++ {
				ret[j-1] += r[j] * j
			}
			i += skip
			pos += skip
			if i == len(s) {
				flg = false
			}
		case '}':
			if x != 0 && a == 0 {
				a++
			}
			ret[x] += a
			x = 0
			a = 0
			pos++
			// out(s[:pos], s)
			return ret, pos
		case '{':
		default:
			a = int(s[i] - '0')
		}
		pos++
	}
	if flg == true {
		if x != 0 && a == 0 {
			a++
		}
		ret[x] += a
	}
	return ret, pos
}

var N, d int

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1000000)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N = getI()
	d = getI()
	s := getS()

	ret, _ := f(s, true)
	for _, e := range ret {
		fmt.Print(e, " ")
	}
	fmt.Println()
}
