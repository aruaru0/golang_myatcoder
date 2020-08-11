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

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1000000)
	// 2014年7月23日 水
	N := getInt()
	cnt := 0
	ans := 0
	n := min(N, 2400)
	for i := 2015; i <= n; i++ {
		cnt += 365
		if i%4 == 0 {
			cnt++
			if i%100 == 0 {
				cnt--
				if i%400 == 0 {
					cnt++
				}
			}
		}
		if cnt%7 == 0 {
			ans++
		}
		cnt %= 7
	}
	if n == N {
		out(ans)
		return
	}
	x := (N - 2400) / 400
	ans += x * 57
	for i := N - N%400 + 1; i <= N; i++ {
		cnt += 365
		if i%4 == 0 {
			cnt++
			if i%100 == 0 {
				cnt--
				if i%400 == 0 {
					cnt++
				}
			}
		}
		if cnt%7 == 0 {
			ans++
		}
		cnt %= 7
	}
	out(ans)
}
