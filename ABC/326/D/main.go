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

//Combination generator for int slice
func combinations(list []int, choose, buf int) (c chan []int) {
	c = make(chan []int, buf)
	go func() {
		defer close(c)
		switch {
		case choose == 0:
			c <- []int{}
		case choose == len(list):
			c <- list
		case len(list) < choose:
			return
		default:
			for i := 0; i < len(list); i++ {
				for subComb := range combinations(list[i+1:], choose-1, buf) {
					c <- append([]int{list[i]}, subComb...)
				}
			}
		}
	}()
	return
}

func NextPermutation(x sort.Interface) bool {
	n := x.Len() - 1
	if n < 1 {
		return false
	}
	j := n - 1
	for ; !x.Less(j, j+1); j-- {
		if j == 0 {
			return false
		}
	}
	l := n
	for !x.Less(j, l) {
		l--
	}
	x.Swap(j, l)
	for k, l := j+1, n; k < l; {
		x.Swap(k, l)
		k++
		l--
	}
	return true
}

func check(a, b, c []int) bool {
	for i := 0; i < N; i++ {
		if a[i] == b[i] || a[i] == c[i] || b[i] == c[i] {
			return false
		}
	}
	return true
}

func check2(s [][]byte) (string, string) {
	r := []byte(".....")[:N]
	c := []byte(".....")[:N]
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			if s[i][j] != '.' {
				r[i] = s[i][j]
				break
			}
		}
	}
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			if s[j][i] != '.' {
				c[i] = s[j][i]
				break
			}
		}
	}

	// out("----")
	// for i := 0; i < N; i++ {
	// 	out(string(s[i]))
	// }
	// out("=====")
	// out(string(r))
	// out(string(c))

	return string(r), string(c)
}

var N int

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N = getI()
	R := getS()
	C := getS()

	a := make([]int, N)
	for i := 0; i < N; i++ {
		a[i] = i
	}
	pat := make([][]int, 0)

	for {
		b := make([]int, N)
		copy(b, a)
		pat = append(pat, b)
		if !NextPermutation(sort.IntSlice(a)) {
			break
		}
	}

	for _, a := range pat {
		for _, b := range pat {
			for _, c := range pat {
				if !check(a, b, c) {
					continue
				}
				s := make([][]byte, N)
				for i := 0; i < N; i++ {
					s[i] = make([]byte, N)
					for j := 0; j < N; j++ {
						s[i][j] = '.'
					}
				}
				for i, j := range a {
					s[i][j] = 'A'
				}
				for i, j := range b {
					s[i][j] = 'B'
				}
				for i, j := range c {
					s[i][j] = 'C'
				}
				// out("-----")
				// out(a, b, c)
				rr, cc := check2(s)
				// out(rr, cc, R, C)
				if rr[:len(R)] == R && cc[:len(C)] == C {
					out("Yes")
					for i := 0; i < N; i++ {
						for j := 0; j < N; j++ {
							fmt.Fprint(wr, string(s[i][j]))
						}
						out()
					}
					return
				}
			}
		}
	}

	out("No")
}
