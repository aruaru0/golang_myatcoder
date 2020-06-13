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

// 先頭からの文字列と各文字から一致する文字列の長さを調べる
// 戻り値は各文字からの一致文字数（０は文字列の長さと一致）
func zalgo(str string) []int {
	n := len(str)
	a := make([]int, n)
	from, last := -1, -1
	a[0] = n
	for i := 1; i < n; i++ {
		idx := a[i]
		if from != -1 {
			idx = min(a[i-from], last-i)
			idx = max(0, idx)
		}
		for idx+i < n && str[idx] == str[idx+i] {
			idx++
		}
		a[i] = idx
		if last < i+idx {
			last = i + idx
			from = i
		}
	}
	return a
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1000000)
	N := getInt()
	S := getString()

	ans := 0
	for n := 0; n < N; n++ {
		// out(S[n:])
		ret := zalgo(S[n:])
		for i := 1; i < len(ret); i++ {
			l := min(i, ret[i])
			ans = max(l, ans)
		}
	}
	out(ans)
}
