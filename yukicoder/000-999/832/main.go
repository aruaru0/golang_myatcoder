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

func check(hand []int) bool {
	r := 0
	a := hand[0]
	b := hand[1]

	for i := 0; i < 7; i++ {
		r = a % 3
		if b >= r && hand[i+2] >= r {
			a = b - r
			b = hand[i+2] - r
		} else {
			return false
		}
	}
	if a%3 == 0 && b%3 == 0 {
		return true
	} else {
		return false
	}
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1000000)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	s := getS()
	x := make([]int, 9)
	for i := 0; i < len(s); i++ {
		n := int(s[i]-'0') - 1
		x[n]++
	}

	ret := make(map[int]int)
	for i := 0; i < 9; i++ {
		x[i]++
		ok := true
		for j := 0; j < 9; j++ {
			if x[j] != 0 && x[j] != 2 {
				ok = false
			}
		}
		if ok == true {
			ret[i+1]++
		}
		x[i]--
	}

	for j := 0; j < 9; j++ {
		if x[j] == 4 {
			continue
		}
		x[j]++
		// out("=====")
		// out(j+1, x)
		// out("-----")
		for k := 0; k < 9; k++ {
			if x[k] >= 2 {
				x[k] -= 2
				// out(x)
				if check(x) {
					ret[j+1]++
					x[k] += 2
					break
				}
				x[k] += 2
			}
		}
		x[j]--
	}

	ans := make([]int, 0)
	for e := range ret {
		ans = append(ans, e)
	}
	sort.Ints(ans)
	for _, e := range ans {
		out(e)
	}
}
