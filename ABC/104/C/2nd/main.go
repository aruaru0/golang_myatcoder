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

func main() {
	sc.Split(bufio.ScanWords)
	D, G := getInt(), getInt()
	p, c := make([]int, D), make([]int, D)
	s := make([]int, D)
	for i := 0; i < D; i++ {
		p[i], c[i] = getInt(), getInt()
		s[i] = 100*(i+1)*p[i] + c[i]
	}

	N := 1 << uint(D)
	ans := math.MaxInt64
	for i := 0; i < N; i++ {
		sum := 0
		cnt := 0
		for j := 0; j < D; j++ {
			if (i>>uint(j))%2 == 1 {
				sum += s[j]
				cnt += p[j]
			}
		}
		if sum >= G {
			// out(sum, D)
			for j := 0; j < D; j++ {
				cnt2 := 0
				sum2 := sum
				if (i>>uint(j))%2 == 1 {
					sum2 -= c[j]
					if sum2 > G {
						for k := 0; k < p[j]; k++ {
							sum2 -= 100 * (j + 1)
							if sum2 < G {
								break
							}
							cnt2++
						}
					}
				}
				// out(strconv.FormatInt(int64(i), 2), "---")
				// out(cnt, sum, cnt2, sum2)
				ans = min(ans, cnt-cnt2)
			}
			// out(sum)
		}
	}
	out(ans)
}
