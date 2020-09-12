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
	D := getInt()
	c := getString() + getString()
	c = "xxxxxxxxxxxxxxx" + c + "xxxxxxxxxxxxxxx"
	p := make([]int, 0)
	for i := 0; i < len(c); i++ {
		if c[i] != 'o' {
			p = append(p, i)
		}
	}

	ans := 0
	for _, e := range p {
		d := D
		flg := false
		cnt := 0
		l := 0
		for i := 0; i < len(c); i++ {
			if i == e {
				flg = true
			}
			if c[i] == 'o' {
				// fmt.Print("o")
				cnt++
				flg = false
			} else if flg == true && d > 0 {
				// fmt.Print("o")
				d--
				cnt++
			} else {
				// fmt.Print("x")
				cnt = 0
			}
			l = max(l, cnt)
		}
		// out()
		// if flg == true {
		// 	cnt += d
		// 	l = max(l, cnt)
		// }
		ans = max(ans, l)
	}
	out(ans)
}
