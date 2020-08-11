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

type pair struct {
	x, y int
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1000000)
	N := 10
	s := make([][]byte, N)
	tot := 0
	for i := 0; i < N; i++ {
		s[i] = []byte(getString())
		for j := 0; j < N; j++ {
			if s[i][j] == 'o' {
				tot++
			}
		}
	}

	dx := []int{-1, 1, 0, 0}
	dy := []int{0, 0, -1, 1}
	var bsf = func(x, y int) int {
		q := make([]pair, 0)
		q = append(q, pair{x, y})
		a := make([][]bool, N)
		for i := 0; i < N; i++ {
			a[i] = make([]bool, N)
		}
		ret := 0
		for len(q) != 0 {
			cx := q[0].x
			cy := q[0].y
			q = q[1:]
			for i := 0; i < 4; i++ {
				px := cx + dx[i]
				py := cy + dy[i]
				if px < 0 || px >= N || py < 0 || py >= N {
					continue
				}
				if s[py][px] == 'x' {
					continue
				}
				if a[py][px] == true {
					continue
				}
				ret++
				a[py][px] = true
				q = append(q, pair{px, py})
			}
		}
		return ret
	}

	for y := 0; y < N; y++ {
		for x := 0; x < N; x++ {
			if bsf(y, x) == tot {
				out("YES")
				return
			}
		}
	}
	out("NO")
}
