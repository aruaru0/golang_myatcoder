package main

import (
	"bufio"
	"fmt"
	"math"
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

const inf = math.MaxInt64 / 2

func main() {
	sc.Split(bufio.ScanWords)
	n := getInt()
	a := make([]int, n)
	for i := 0; i < n; i++ {
		a[i] = getInt()
	}

	ans := inf
	for i := 0; i < 2; i++ {
		sign := true
		if i == 1 {
			sign = false
		}
		sum := 0
		cnt := 0
		// out("-----")
		for j := 0; j < n; j++ {
			sum += a[j]
			x := a[j]
			if sign {
				if sum <= 0 {
					cnt += -sum + 1
					x += sum + 1
					sum = 1
				}
				sign = false
			} else {
				if sum >= 0 {
					cnt += sum + 1
					x -= sum + 1
					sum = -1
				}
				sign = true
			}
			// out(x, sum)
		}
		ans = min(ans, cnt)
	}
	out(ans)
}
