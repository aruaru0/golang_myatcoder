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

var pat = [10][5]string{
	{
		"###",
		"#.#",
		"#.#",
		"#.#",
		"###",
	}, {
		".#.",
		"##.",
		".#.",
		".#.",
		"###",
	}, {
		"###",
		"..#",
		"###",
		"#..",
		"###",
	}, {
		"###",
		"..#",
		"###",
		"..#",
		"###",
	}, {
		"#.#",
		"#.#",
		"###",
		"..#",
		"..#",
	}, {
		"###",
		"#..",
		"###",
		"..#",
		"###",
	}, {
		"###",
		"#..",
		"###",
		"#.#",
		"###",
	}, {
		"###",
		"..#",
		"..#",
		"..#",
		"..#",
	}, {
		"###",
		"#.#",
		"###",
		"#.#",
		"###",
	}, {
		"###",
		"#.#",
		"###",
		"..#",
		"###",
	}}

func check(n int, s []string) {
	f := 4*n - 3
	t := 4 * n

	ans := -1
	for k := 0; k < 10; k++ {
		ok := true
		for i := 0; i < 5; i++ {
			if pat[k][i] != s[i][f:t] {
				ok = false
				break
			}
		}
		if ok {
			ans = k
		}
	}
	fmt.Print(ans)
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1000000)

	N := getInt()
	s := make([]string, 5)
	for i := 0; i < 5; i++ {
		s[i] = getString()
	}
	for i := 1; i <= N; i++ {
		check(i, s)
	}
	out()
}
