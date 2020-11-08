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

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1000000)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	H, W := getI(), getI()
	N, M := getI(), getI()
	s := make([][]byte, H)
	a := make([][]bool, H)
	for i := 0; i < H; i++ {
		s[i] = make([]byte, W)
		a[i] = make([]bool, W)
		for j := 0; j < W; j++ {
			s[i][j] = '.'
		}
	}
	for i := 0; i < N; i++ {
		a, b := getI()-1, getI()-1
		s[a][b] = 'o'
	}
	for i := 0; i < M; i++ {
		c, d := getI()-1, getI()-1
		s[c][d] = 'x'
	}

	for y := 0; y < H; y++ {
		flg := false
		for x := 0; x < W; x++ {
			if s[y][x] == 'o' {
				flg = true
			}
			if s[y][x] == 'x' {
				flg = false
			}
			a[y][x] = a[y][x] || flg
		}
		flg = false
		for x := W - 1; x >= 0; x-- {
			if s[y][x] == 'o' {
				flg = true
			}
			if s[y][x] == 'x' {
				flg = false
			}
			a[y][x] = a[y][x] || flg
		}
	}

	for x := 0; x < W; x++ {
		flg := false
		for y := 0; y < H; y++ {
			if s[y][x] == 'o' {
				flg = true
			}
			if s[y][x] == 'x' {
				flg = false
			}
			a[y][x] = a[y][x] || flg
		}
		flg = false
		for y := H - 1; y >= 0; y-- {
			if s[y][x] == 'o' {
				flg = true
			}
			if s[y][x] == 'x' {
				flg = false
			}
			a[y][x] = a[y][x] || flg
		}
	}

	// for i := 0; i < H; i++ {
	// 	out(string(s[i]))
	// }

	// for i := 0; i < H; i++ {
	// 	out(a[i])
	// }

	ans := 0
	for i := 0; i < H; i++ {
		for j := 0; j < W; j++ {
			if a[i][j] == true {
				ans++
			}
		}
	}
	out(ans)
}
