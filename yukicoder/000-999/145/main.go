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

	m := make([]int, 26)
	for i := 0; i < N; i++ {
		m[int(s[i]-'a')]++
	}
	// out("[a b c d e f g h i j k l m n o p q r s t u v w x y z]")
	// out(m)
	// Z
	ans := m[25]
	// out(m[24], m[20], m[10], m[8], ans)
	// Y[VWX]
	cnt := 0
	for i := 21; i <= 23; i++ {
		cnt += m[i]
	}
	l := min(m[24], cnt)
	ans += l
	m[24] -= l
	// out(m[24], m[20], m[10], m[8], ans, cnt)

	// YU...
	cnt = 0
	for i := 11; i <= 19; i++ {
		cnt += m[i]
	}
	t := min(m[24], m[20])
	u := min(t, cnt)
	ans += u
	m[24] -= u
	m[20] -= u

	// out(m[24], m[20], m[10], m[8], ans, cnt)

	// YUK...
	cnt = m[9]
	t = min(m[24], min(m[20], m[10]))
	u = min(t, cnt)
	ans += u
	m[24] -= u
	m[20] -= u
	m[10] -= u

	// out(m[24], m[20], m[10], m[8], ans, cnt)

	// YUKI...
	cnt = 0
	for i := 0; i <= 7; i++ {
		cnt += m[i]
	}
	t = min(m[24], min(m[20], min(m[10], m[8])))
	u = min(t, cnt)
	m[24] -= u
	m[20] -= u
	m[10] -= u
	m[8] -= u
	// out(m[24], m[20], m[10], m[8], ans, cnt)
	ans += u
	a := []int{m[24], m[20], m[10], m[8]}
	// out(a)
	for i := 3; i >= 0; i-- {
		for a[i] >= 2 {
			flg := true
			for j := i - 1; j >= 0; j-- {
				if a[j] == 0 {
					flg = false
					break
				}
				a[j]--
			}
			if flg {
				ans++
			}
			a[i] -= 2
		}
		// out(a)
	}
	out(ans)
}
