package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)
var wr = bufio.NewWriter(os.Stdout)

func out(x ...interface{}) {
	fmt.Fprintln(wr, x...)
}

func getI() int {
	sc.Scan()
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return i
}

func getF() float64 {
	sc.Scan()
	i, e := strconv.ParseFloat(sc.Text(), 64)
	if e != nil {
		panic(e)
	}
	return i
}

func getInts(N int) []int {
	ret := make([]int, N)
	for i := 0; i < N; i++ {
		ret[i] = getI()
	}
	return ret
}

func getS() string {
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

func move(a, b [][]int) ([][]int, int) {
	cx, cy := 0, 0
	flg := true
	for y := 0; y < 4; y++ {
		for x := 0; x < 4; x++ {
			if a[y][x] != b[y][x] {
				flg = false
			}
			if a[y][x] == 0 {
				cx, cy = x, y
			}
		}
	}

	if flg == true {
		return a, 1
	}

	dx := []int{-1, 1, 0, 0}
	dy := []int{0, 0, -1, 1}
	for i := 0; i < 4; i++ {
		px := cx + dx[i]
		py := cy + dy[i]
		if px < 0 || py < 0 || px >= 4 || py >= 4 {
			continue
		}
		if b[cy][cx] == a[py][px] {
			a[py][px], a[cy][cx] = a[cy][cx], a[py][px]
			return a, 0
		}
	}
	return a, -1
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1000000)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	a := make([][]int, 4)
	b := make([][]int, 4)
	for i := 0; i < 4; i++ {
		a[i] = getInts(4)
	}
	b[0] = []int{1, 2, 3, 4}
	b[1] = []int{5, 6, 7, 8}
	b[2] = []int{9, 10, 11, 12}
	b[3] = []int{13, 14, 15, 0}

	for {
		var ok int
		b, ok = move(b, a)
		if ok == 1 {
			out("Yes")
			break
		}
		if ok == -1 {
			out("No")
			break
		}
	}
}
