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

func condAB(c bool, a, b int) int {
	if c {
		return a
	}
	return b
}

func main() {
	sc.Split(bufio.ScanWords)

	s := getString()
	k := getInt()

	nbits := len(s)
	d := make([]int, nbits)
	for i := 0; i < nbits; i++ {
		d[i] = int(s[i] - '0')
	}

	//  [digit][smaller][cond1]
	//  cond1 ... 3である
	var dp [101][2][5]int

	dp[0][0][0] = 1
	for i := 0; i < nbits; i++ {
		for m := 0; m < 2; m++ {
			for a := 0; a <= condAB(m != 0, 9, d[i]); a++ {
				mm := m
				if a < d[i] {
					mm = 1
				}
				// cond1
				if a != 0 {
					for k := 1; k < 5; k++ {
						dp[i+1][mm][k] += dp[i][m][k-1]
					}
				} else {
					for k := 0; k < 5; k++ {
						dp[i+1][mm][k] += dp[i][m][k]
					}
				}
			}
		}
	}

	// for i := 0; i <= nbits; i++ {
	// 	out(dp[i], k)
	// }
	out(dp[nbits][1][k] + dp[nbits][0][k])

}
