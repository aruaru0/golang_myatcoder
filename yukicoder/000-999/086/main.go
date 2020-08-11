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

var N, M int
var s []string
var tot int
var sx, sy int
var dx = [4]int{1, 0, -1, 0}
var dy = [4]int{0, -1, 0, 1}

func solve(sx, sy, dir int) bool {
	if s[sy][sx] == '#' {
		return false
	}
	used := make([][]bool, N)
	for i := 0; i < N; i++ {
		used[i] = make([]bool, M)
	}
	x := sx
	y := sy
	c := 0
	cnt := 0
	for {
		if c == 2 {
			break
		}
		px := x + dx[dir]
		py := y + dy[dir]
		if px < 0 || py < 0 || px >= M || py >= N {
			dir = (dir + 1) % 4
			c++
			continue
		}
		if s[py][px] == '#' {
			dir = (dir + 1) % 4
			c++
			continue
		}
		// out(px, py)
		if used[py][px] {
			break
		}
		c = 0
		used[py][px] = true
		x, y = px, py
		cnt++
	}
	// out(cnt, x, y, tot)
	if x == sx && y == sy && cnt == tot {
		return true
	}
	return false
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1000000)
	N, M = getInt(), getInt()
	s = make([]string, N)
	for i := 0; i < N; i++ {
		s[i] = getString()
	}
	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			if s[i][j] != '#' {
				tot++
			}
		}
	}

	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			for k := 0; k < 4; k++ {
				ret := solve(j, i, k)
				if ret {
					// out(j, i, k)
					out("YES")
					return
				}
			}
		}
	}
	out("NO")
}
