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

type ints []int

const d = 15

func main() {
	sc.Split(bufio.ScanWords)
	wr := bufio.NewWriter(os.Stdout)
	defer wr.Flush()

	N, S := getInt(), getInt()
	p := getInts(N)

	n := min(N, d)
	bit := 1 << n
	sum0 := make([]int, bit)
	for i := 0; i < bit; i++ {
		for j := 0; j < n; j++ {
			if (i>>j)%2 == 1 {
				sum0[i] += p[j]
			}
		}
	}
	m := N - d
	ans := make([][]int, 0)
	if m <= 0 {
		for i, v := range sum0 {
			if v == S {
				pat := make([]int, 0)
				for j := 0; j < n; j++ {
					if (i>>j)%2 == 1 {
						pat = append(pat, j+1)
					}
				}
				ans = append(ans, pat)
			}
		}
	} else {
		bit = 1 << m
		sum1 := make([]int, bit)
		for i := 0; i < bit; i++ {
			for j := 0; j < m; j++ {
				if (i>>j)%2 == 1 {
					sum1[i] += p[n+j]
				}
			}
		}
		mp := make(map[int]ints, bit)
		for i := 0; i < bit; i++ {
			mp[sum1[i]] = append(mp[sum1[i]], i)
		}

		for i, v := range sum0 {
			e, ok := mp[S-v]
			if ok {
				for _, k := range e {
					pat := make([]int, 0)
					for j := 0; j < n; j++ {
						if (i>>j)%2 == 1 {
							pat = append(pat, j+1)
						}
					}
					for j := 0; j < m; j++ {
						if (k>>j)%2 == 1 {
							pat = append(pat, n+j+1)
						}
					}
					ans = append(ans, pat)
				}
			}
		}
	}

	sort.Slice(ans, func(i, j int) bool {
		l := min(len(ans[i]), len(ans[j]))
		for k := 0; k < l; k++ {
			if ans[i][k] != ans[j][k] {
				return ans[i][k] < ans[j][k]
			}
		}
		if len(ans[i]) > len(ans[j]) {
			return false
		}
		return true
	})

	for _, e := range ans {
		for _, v := range e {
			fmt.Fprint(wr, v, " ")
		}
		fmt.Fprintln(wr)
	}

	return
}
