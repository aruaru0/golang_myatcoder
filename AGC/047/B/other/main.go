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

// Trie :
type Trie struct {
	c     int
	child [26]int
	cnt   [26]int
	p     int
	fin   bool
}

func rev(str string) string {
	s := []byte(str)
	N := len(s)
	for i := 0; i < N/2; i++ {
		s[i], s[N-1-i] = s[N-1-i], s[i]
	}
	return string(s)
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1000000)

	trie := make([]Trie, 1100000)
	for i := 0; i < 1100000; i++ {
		for j := 0; j < 26; j++ {
			trie[i].child[j] = -1
		}
	}
	p := 1
	queries := make([]int, 0)

	n := getInt()
	for i := 0; i < n; i++ {
		str := rev(getString())
		now := 0
		for j := 0; j < len(str); j++ {
			c := int(str[j] - 'a')
			if trie[now].child[c] == -1 {
				trie[now].child[c] = p
				p++
				trie[trie[now].child[c]].p = now
			}
			now = trie[now].child[c]
			trie[now].c = c
		}
		trie[now].fin = true
		queries = append(queries, now)
		have := make([]bool, 26)
		for i := len(str) - 1; i >= 0; i-- {
			c := str[i] - 'a'
			have[c] = true
			now = trie[now].p
			for j := 0; j < 26; j++ {
				if have[j] {
					trie[now].cnt[j]++
				}
			}
		}
	}
	ans := 0
	for _, pos := range queries {
		p := trie[pos].p
		c := trie[pos].c
		ans += trie[p].cnt[c] - 1
	}
	out(ans)
}
