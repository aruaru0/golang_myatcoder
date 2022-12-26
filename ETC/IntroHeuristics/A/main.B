package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func out(x ...interface{}) {
	fmt.Println(x...)
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

func calcScore(day int, score, c, last []int) int {
	loss := make([]int, 26)
	for i := 0; i < 26; i++ {
		loss[i] = c[i] * (day - last[i])
	}

	ret := 0
	for j := 0; j < 26; j++ {
		score[j] -= loss[j]
		ret += score[j]
	}

	return ret
}

func main() {
	sc.Split(bufio.ScanWords)
	D := getInt()
	c := getInts(26)
	s := make([][]int, D)
	for i := 0; i < D; i++ {
		s[i] = getInts(26)
	}

	last := make([]int, 26)
	score := make([]int, 26)

	for i := 0; i < D; i++ {
		t := getInt() - 1
		score[t] += s[i][t]
		last[t] = i + 1
		ans := calcScore(i+1, score, c, last)
		out(ans)
		// out(score)
	}

	// for i := 0; i < D; i++ {
	// 	maxS := 0
	// 	idxS := 0
	// 	for j := 0; j < 26; j++ {
	// 		if maxS < s[i][j] {
	// 			maxS = s[i][j]
	// 			idxS = j
	// 		}
	// 	}
	// 	out(idxS + 1)
	// }
}
