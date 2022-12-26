package main

import (
	"bufio"
	"fmt"
	"math/bits"
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

func check(s string, fr, fb, r, b int) bool {
	cntR := 0
	cntB := 0
	str := make([]byte, 0)

	for i := 0; i < len(s); i++ {
		if s[i] == 'W' {
			str = append(str, s[i])
		} else if s[i] == 'R' {
			if (fr>>cntR)&1 == 1 {
				str = append(str, s[i])
			}
			cntR++
		} else if s[i] == 'B' {
			if (fb>>cntB)&1 == 1 {
				str = append(str, s[i])
			}
			cntB++
		}
	}
	// fmt.Printf("%10.10b %10.10b ", fr, fb)
	// out(s, string(str))
	for i := 0; i < len(str); i++ {
		if str[i] == 'W' {
			continue
		}
		if i+r < len(str) {
			if str[i] == 'R' && str[i+r] == 'R' {
				return false
			}
		}
		if i+b < len(str) {
			if str[i] == 'B' && str[i+b] == 'B' {
				return false
			}
		}
	}
	return true
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1000000)
	Kr, Kb := getInt(), getInt()
	s := getString()
	ans := 0
	t := 1 << 10
	for r := 0; r < t; r++ {
		for b := 0; b < t; b++ {
			c := check(s, r, b, Kr, Kb)
			// out(c)
			if c == true {
				ans = max(ans,
					bits.OnesCount(uint(r))+
						bits.OnesCount(uint(b))+10)
			}
		}
	}
	out(ans)
}
