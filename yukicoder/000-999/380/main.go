package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"unicode"
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

func check0(s string, stop bool) int {
	cnt := 0
	for i := len(s) - 1; i >= 0; i-- {
		r := rune(s[i])
		if !unicode.IsDigit(r) && !unicode.IsLower(r) && !unicode.IsUpper(r) {
			cnt++
		} else if stop == true {
			break
		}
	}
	return cnt
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanLines)
	sc.Buffer([]byte{}, 1000000)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	for {
		tok := sc.Scan()
		if !tok {
			break
		}
		s := sc.Text()
		a := strings.Split(s, " ")
		name := a[0]
		last := ""
		for i := 1; i < len(a); i++ {
			last += a[i]
		}
		last = strings.ToLower(last)
		ok := true
		if len(a) <= 1 {
			ok = false
		}
		switch name {
		case "digi":
			c := check0(last, true)
			if c > 3 {
				ok = false
			}
			l := max(0, len(last)-c-3)
			r := max(0, len(last)-c)
			last = last[l:r]
			if last != "nyo" {
				ok = false
			}
		case "petit":
			c := check0(last, true)
			if c > 3 {
				ok = false
			}
			l := max(0, len(last)-c-3)
			r := max(0, len(last)-c)
			last = last[l:r]
			if last != "nyu" {
				ok = false
			}
		case "rabi":
			tot := 0
			cnt := 0
			for i := 1; i < len(a); i++ {
				tot += len(a[i])
				cnt += check0(a[i], false)
			}
			if tot == cnt {
				ok = false
			}
		case "gema":
			c := check0(last, true)
			if c > 3 {
				ok = false
			}
			l := max(0, len(last)-c-4)
			r := max(0, len(last)-c)
			last = last[l:r]
			if last != "gema" {
				ok = false
			}
		case "piyo":
			c := check0(last, true)
			if c > 3 {
				ok = false
			}
			l := max(0, len(last)-c-3)
			r := max(0, len(last)-c)
			last = last[l:r]
			if last != "pyo" {
				ok = false
			}
		default:
			ok = false
		}

		if ok {
			out("CORRECT (maybe)")
		} else {
			out("WRONG!")
		}
		// out(name, last, ok)
	}
}
