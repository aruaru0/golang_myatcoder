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

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N := getI()
	n := N + 2
	s := getS() + ".."
	t := getS() + ".."

	dist := make(map[string]int)
	q := make([]string, 0)
	q = append(q, s)
	for len(q) != 0 {
		cur := q[0]
		q = q[1:]
		// search empty
		if cur == t {
			out(dist[cur])
			return
		}
		ePos := 0
		for i := 0; i < n; i++ {
			if cur[i] == '.' {
				ePos = i
				break
			}
		}
		for i := 0; i < n-1; i++ {
			if cur[i] == '.' || cur[i+1] == '.' {
				continue
			}
			tmp := make([]byte, n)
			copy(tmp, []byte(cur))
			// out("copy", cur, string(tmp))
			tmp[ePos], tmp[ePos+1] = cur[i], cur[i+1]
			tmp[i], tmp[i+1] = '.', '.'
			nxt := string(tmp)
			if _, ok := dist[nxt]; ok == false {
				dist[nxt] = dist[cur] + 1
				q = append(q, nxt)
			}
		}

	}
	out(-1)
}
