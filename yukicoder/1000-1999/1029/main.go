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

const inf = int(1e18)

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1000000)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N, K := getI(), getI()
	s := make([]string, N)
	c := make([]int, N)
	for i := 0; i < N; i++ {
		s[i] = getS()
		c[i] = getI()
	}

	cntj := make([]int, N)
	cnto := make([]int, N)
	cnti := make([]int, N)
	for i := 0; i < N; i++ {
		cntj[i] = strings.Count(s[i], "J")
		cnto[i] = strings.Count(s[i], "O")
		cnti[i] = strings.Count(s[i], "I")
	}

	cost := make([]int, K*3+1)
	for i := 0; i < len(cost); i++ {
		cost[i] += inf
	}
	cost[0] = 0
	for i := 0; i < K*3; i++ {
		{
			if cost[i] < cost[i+1] {
				for j := 0; j < N; j++ {
					if i+100 < K {
						cost[i+cntj[j]] = min(cost[i+cntj[j]], cost[i]+c[j])
					} else if i >= K && i+100 < K*2 {
						cost[i+cnto[j]] = min(cost[i+cnto[j]], cost[i]+c[j])
					} else if i >= K*2 && i+100 < K*3 {
						cost[i+cnti[j]] = min(cost[i+cnti[j]], cost[i]+c[j])
					} else {
						now := i
						for _, c := range s[j] {
							if now < K && c == 'J' {
								now++
							}
							if now >= K && now < K*2 && c == 'O' {
								now++
							}
							if now >= K*2 && now < K*3 && c == 'I' {
								now++
							}
						}
						cost[now] = min(cost[now], cost[i]+c[j])
					}
				}
			}
		}
	}
	if cost[len(cost)-1] == inf {
		out(-1)
		return
	}
	out(cost[len(cost)-1])
}
