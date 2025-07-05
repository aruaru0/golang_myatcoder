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

func outSlice[T any](s []T) {
	if len(s) == 0 {
		return
	}
	for i := 0; i < len(s)-1; i++ {
		fmt.Fprint(wr, s[i], " ")
	}
	fmt.Fprintln(wr, s[len(s)-1])
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

	N, K := getI(), getI()
	S := getS()

	// sFreq[i][c] は S[i:] の文字 c の出現回数
	cnt := make([][26]int, N+1)
	for i := N - 1; i >= 0; i-- {
		for j := 0; j < 26; j++ {
			cnt[i][j] = cnt[i+1][j]
		}
		c := S[i]
		idx := int(c - 'a')
		cnt[i][idx]++
	}

	// available[c] は使用可能な文字 c の数
	available := make([]int, 26)
	for _, c := range S {
		idx := int(c - 'a')
		available[idx]++
	}

	currentDiff := 0
	result := make([]byte, 0, N)

	for i := 0; i < N; i++ {
		// a~z順に試す
		for c := 0; c < 26; c++ {
			if available[c] == 0 {
				continue
			}

			// 文字cをi番目に置いたと仮定する
			available[c]--

			newDiff := currentDiff
			if byte(c+'a') != S[i] {
				newDiff++
			}

			// 残りの文字で最大何文字一致させられるか
			m := N - i - 1
			var sRemaining [26]int
			if i+1 < len(cnt) {
				sRemaining = cnt[i+1]
			}

			maxMatch := 0
			for j := 0; j < 26; j++ {
				maxMatch += min(sRemaining[j], available[j])
			}

			// 必要な変更回数 = 現在の変更回数 + (残りの文字数 - 最大一致数)
			required := newDiff + (m - maxMatch)

			if required <= K {
				result = append(result, byte(c+'a'))
				currentDiff = newDiff
				break // この文字cで確定
			} else {
				// 仮定を元に戻す
				available[c]++
			}
		}
	}

	out(string(result))
}
