package main

import (
	"bufio"
	"fmt"
	"os"
)

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
			// fmt.Println("update", idx, i-from, last-i)
		}
		for idx+i < n && str[idx] == str[idx+i] {
			idx++
		}
		a[i] = idx
		if last < i+idx {
			last = i + idx
			from = i
		}
		// fmt.Println(i, idx, from, last, str[i:])
	}
	return a
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)

	sc.Scan()
	s := sc.Text()

	ret := zalgo(s)

	fmt.Println(ret)
}
