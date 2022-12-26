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
	sc.Buffer([]byte{}, 1000000)
	N := getInt()
	s := getString()

	var y, u, k, i, a_h, j, l_t, v_x, z int
	for x := 0; x < N; x++ {
		switch {
		case s[x] == 'z':
			z++
		case s[x] == 'y':
			y++
		case s[x] > 'u':
			v_x++
		case s[x] == 'u':
			u++
		case s[x] > 'k':
			l_t++
		case s[x] == 'k':
			k++
		case s[x] > 'i':
			j++
		case s[x] == 'i':
			i++
		default:
			a_h++
		}
	}
	ans := 0
	ans += z
	l := min(y, v_x)
	ans += l
	y -= l
	l = min(min(y, u), l_t)
	ans += l
	y -= l
	u -= l
	l = min(min(y, min(u, k)), j)
	ans += l
	y -= l
	u -= l
	k -= l
	l = min(min(y, min(u, min(k, i))), a_h)
	ans += l

	y -= l
	u -= l
	k -= l
	i -= l
	for y != 0 && u != 0 && k != 0 && i >= 2 {
		y--
		u--
		k--
		i -= 2
		ans++
	}
	for y != 0 && u != 0 && k >= 2 {
		y--
		u--
		k -= 2
		ans++
	}
	for y != 0 && u >= 2 {
		y--
		u -= 2
		ans++
	}
	for y >= 2 {
		y -= 2
		ans++
	}
	out(ans)
}
