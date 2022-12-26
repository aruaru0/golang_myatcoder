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

func rev(s []byte) string {
	N := len(s)
	for i := 0; i < N/2; i++ {
		s[i], s[N-1-i] = s[N-1-i], s[i]
	}
	return string(s)
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1000000)
	N := getInt()
	s := make([]string, N)
	for i := 0; i < N; i++ {
		s[i] = rev([]byte(getString()))
	}
	sort.Slice(s, func(i, j int) bool {
		return len(s[i]) < len(s[j])
	})
	ans := 0
	mp := make(map[int]int)
	off := 1009
	for _, x := range s {
		hash := make([]int, len(x))
		hash[0] = int(x[0])
		for i := 1; i < len(x); i++ {
			hash[i] = hash[i-1]*off + int(x[i])
		}
		cnt := make([]int, 26)
		for i := 0; i < len(x); i++ {
			cnt[int(x[i]-'a')]++
		}
		for c := 0; c < 26; c++ {
			if cnt[c] == 0 {
				continue
			}
			h := c + 'a'
			ans += mp[h]
		}
		for i := 0; i < len(x)-1; i++ {
			h := hash[i] * off
			cnt[x[i]-'a']--
			for c := 0; c < 26; c++ {
				if cnt[c] == 0 {
					continue
				}
				h += c + 'a'
				ans += mp[h]
				h -= c + 'a'
			}
		}
		mp[hash[len(x)-1]]++
	}
	out(ans)
}
