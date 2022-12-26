package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func out(x ...interface{}) {
	// fmt.Println(x...)
}

var sc = bufio.NewScanner(os.Stdin)

func getInt() int {
	sc.Scan()
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return i
}

func getInts(N int) []int {
	ret := make([]int, N)
	for i := 0; i < N; i++ {
		ret[i] = getInt()
	}
	return ret
}

func getString() string {
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
	pos int
	dir byte
}

const inf = int(1e12)

func check(v []pair, s string) int {
	sort.Slice(v, func(i, j int) bool {
		return v[i].pos < v[j].pos
	})
	ans := inf
	for i := 0; i < len(v)-1; i++ {
		if v[i].dir == s[0] && v[i+1].dir == s[1] {
			ans = min(ans, v[i+1].pos-v[i].pos)
		}
	}
	return ans
}

func main() {
	sc.Split(bufio.ScanWords)
	N := getInt()
	X := make([]int, N)
	Y := make([]int, N)
	U := make([]byte, N)
	for i := 0; i < N; i++ {
		X[i] = getInt()
		Y[i] = getInt()
		U[i] = getString()[0]
	}
	out(X, Y, string(U), U)
	UD := make(map[int][]pair)
	LR := make(map[int][]pair)
	DR := make(map[int][]pair)
	DL := make(map[int][]pair)
	UR := make(map[int][]pair)
	UL := make(map[int][]pair)
	for i := 0; i < N; i++ {
		if U[i] == 'U' || U[i] == 'D' {
			UD[X[i]] = append(UD[X[i]], pair{Y[i], U[i]})
		}
		if U[i] == 'L' || U[i] == 'R' {
			LR[Y[i]] = append(LR[Y[i]], pair{X[i], U[i]})
		}
		if U[i] == 'D' || U[i] == 'R' {
			p := Y[i] - X[i]
			q := Y[i]
			DR[p] = append(DR[p], pair{q, U[i]})
		}
		if U[i] == 'D' || U[i] == 'L' {
			p := Y[i] + X[i]
			q := Y[i]
			DL[p] = append(DL[p], pair{q, U[i]})
		}
		if U[i] == 'U' || U[i] == 'R' {
			p := Y[i] + X[i]
			q := Y[i]
			UR[p] = append(UR[p], pair{q, U[i]})
		}
		if U[i] == 'U' || U[i] == 'L' {
			p := Y[i] - X[i]
			q := Y[i]
			UL[p] = append(UL[p], pair{q, U[i]})
		}
	}
	out(UD)
	out(LR)
	out(DR)
	out(DL)
	out(UR)
	out(UL)

	ans := inf
	for _, e := range UD {
		ret := check(e, "UD")
		if ret != inf {
			ans = min(ans, ret*5)
		}
	}
	out(ans)
	for _, e := range LR {
		ret := check(e, "RL")
		if ret != inf {
			ans = min(ans, ret*5)
		}
	}
	out(ans)
	for _, e := range DR {
		ret := check(e, "RD")
		if ret != inf {
			ans = min(ans, ret*10)
		}
	}
	out(ans)
	for _, e := range DL {
		ret := check(e, "LD")
		if ret != inf {
			ans = min(ans, ret*10)
		}
	}
	out(ans)
	for _, e := range UR {
		ret := check(e, "UR")
		if ret != inf {
			ans = min(ans, ret*10)
		}
	}
	out(ans)
	for _, e := range UL {
		ret := check(e, "UL")
		if ret != inf {
			ans = min(ans, ret*10)
		}
	}
	out(ans)

	if ans == inf {
		fmt.Println("SAFE")
		return
	}
	fmt.Println(ans)
}
