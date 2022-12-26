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
	N := getInt()
	a := getInts(N)

	r := 0
	for i := 0; i < 64; i++ {
		j := r
		for ; j < N; j++ {
			if (a[j]>>i)&1 == 1 {
				break
			}
		}
		if j != N {
			a[j], a[r] = a[r], a[j]
			for k := r + 1; k < N; k++ {
				if (a[k]>>i)&1 == 1 {
					a[k] ^= a[r]
				}
			}
			r++
		}
	}
	ans := 0
	for i := 0; i < N; i++ {
		if a[i] != 0 {
			ans++
		}
	}
	// for i := 0; i < N; i++ {
	// 	fmt.Printf("%60.60b\n", a[i])
	// }
	// out(a)
	out(1 << ans)
}
