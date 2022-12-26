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

var H, W int
var s [][]byte

const mod = int(1e9 + 7)

var dp [][2]int

func pat() int {
	dp := make([][2]int, 101)
	dp[0][0] = 1
	for x := 0; x < W; x++ {
		for i := 0; i < 2; i++ {
			for j := 0; j < 2; j++ {
				ng := 0
				for y := 0; y < H; y++ {
					if s[y][x] != '?' && s[y][x] != byte('0'+(i+j+y)%2) {
						ng++
					}
				}
				if ng == 0 {
					dp[x+1][(i+j)%2] += dp[x][i]
					dp[x+1][(i+j)%2] %= mod
				}
			}
		}
	}
	return dp[W][0] + dp[W][1]
}

func ichimatu() int {
	ok1, ok2 := 1, 1
	for y := 0; y < H; y++ {
		for x := 0; x < W; x++ {
			if s[y][x] != '?' {
				if int(s[y][x]-'0') != ((y + x) % 2) {
					ok1 = 0
				}
				if int(s[y][x]-'0') == ((y + x) % 2) {
					ok2 = 0
				}
			}
		}
	}
	return ok1 + ok2
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1000000)
	H, W = getInt(), getInt()
	s = make([][]byte, H)
	for i := 0; i < H; i++ {
		s[i] = []byte(getString())
	}

	ret := mod - ichimatu()

	for i := 0; i < 2; i++ {
		ret += pat()
		t := make([][]byte, W)
		for i := 0; i < W; i++ {
			t[i] = make([]byte, H)
		}
		for x := 0; x < W; x++ {
			for y := 0; y < H; y++ {
				t[x][y] += s[y][x]
			}
		}
		s = t
		W, H = H, W
	}
	out(ret % mod)
}
