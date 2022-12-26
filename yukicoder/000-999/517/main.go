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
	a := make([]string, N)
	for i := 0; i < N; i++ {
		a[i] = getS()
	}
	M := getI()
	b := make([]string, M)
	for i := 0; i < M; i++ {
		b[i] = getS()
	}

	if M == 1 && N == 1 {
		if a[0] == b[0] {
			out(a[0])
		} else {
			out(-1)
		}
		return
	}

	s := ""
	for i := 0; i < N; i++ {
		if len(s) < len(a[i]) {
			s = a[i]
		}
		for j := 0; j < M; j++ {
			if a[i] == b[j] {
				out(-1)
				return
			}
			if len(s) < len(b[j]) {
				s = b[j]
			}
		}
	}

	usedN := make([]bool, N)
	usedM := make([]bool, M)
	cntN, cntM := 0, 0
	for {
		change := false
		var ok bool
		for i := 0; i < N; i++ {
			if usedN[i] == true {
				continue
			}
			s, ok = match(s, a[i])
			if ok {
				cntN++
				change = true
				usedN[i] = true
			}
			// out(s, a[i], ok)
		}
		for j := 0; j < M; j++ {
			if usedM[j] == true {
				continue
			}
			s, ok = match(s, b[j])
			// out(s, b[j], ok)
			if ok {
				cntM++
				change = true
				usedM[j] = true
			}
		}
		if !change {
			break
		}
	}

	if cntN == N && cntM == M {
		out(s)
		return
	}
	out(-1)
}

func match(s, x string) (string, bool) {
	if s == x {
		return s, true
	}
	if strings.Contains(s, x) {
		return s, true
	}
	for i := 1; i < len(s); i++ {
		if strings.HasSuffix(x, s[:i]) {
			return x + s[i:], true
		}
		if strings.HasPrefix(x, s[i:]) {
			return s[:i] + x, true
		}
	}
	return s, false
}
