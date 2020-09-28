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

var id, N, K int

type pair struct {
	n, col int
}

var dx []int = []int{-1, 1, 0, 0}
var dy []int = []int{0, 0, -1, 1}

type pos struct {
	x, y int
	col  int
}

func paint(x, y, tgt int) {
	q := make([]pos, 0)
	q = append(q, pos{x, y, tgt})
	c := a[y][x]
	t := byte(tgt + '0')
	if c == t {
		return
	}
	a[y][x] = t
	for len(q) != 0 {
		cx := q[0].x
		cy := q[0].y
		q = q[1:]
		for i := 0; i < 4; i++ {
			px := cx + dx[i]
			py := cy + dy[i]
			if px < 0 || py < 0 || px >= N || py >= N {
				continue
			}
			if a[py][px] != c {
				continue
			}
			a[py][px] = t
			q = append(q, pos{px, py, tgt})
		}
	}
}

var a [][]byte

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1000000)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()
	id, N, K = getInt(), getInt(), getInt()

	a = make([][]byte, N)
	for i := 0; i < N; i++ {
		a[i] = []byte(getString())
	}

	tgt := 1
	ans := make([]pos, 0)
	frame := 0
	for {
		cnt := 0
		for y := 0; y < N; y++ {
			for x := 0; x < N; x++ {
				if int(a[y][x]-'0') == tgt {
					cnt++
				}
			}
		}
		if cnt == 10000 {
			break
		}
		tgt = (tgt+1)%9 + 1
		//out(frame, tgt, cnt)
		frame++
		paint(50, 50, tgt)
		ans = append(ans, pos{51, 51, tgt})
	}

	fmt.Fprintln(w, len(ans))
	for _, e := range ans {
		fmt.Fprintln(w, e.y, e.x, e.col)
	}
}
