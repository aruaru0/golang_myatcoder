package main

import (
	"bufio"
	"fmt"
	"os"
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

func bit(n, N int) ([]int, bool) {
	ret := make([]int, N)
	idx := 0
	cnt := []int{0, 0, 0, 0}
	for n > 0 {
		ret[idx] = n % 4
		n /= 4
		idx++
	}

	for i := 0; i < N; i++ {
		cnt[ret[i]]++
	}

	f := cnt[0] != 0 && cnt[1] != 0 && cnt[2] != 0

	return ret, f
}

const inf = 1001001001001

func main() {
	sc.Split(bufio.ScanWords)

	var A [3]int
	N := getInt()
	A[0], A[1], A[2] = getInt(), getInt(), getInt()
	l := make([]int, N)
	for i := 0; i < N; i++ {
		l[i] = getInt()
	}

	cnt := 1
	for i := 0; i < N; i++ {
		cnt *= 4
	}

	ans := inf

	for i := 0; i < cnt; i++ {
		b, flg := bit(i, N)
		if flg == false {
			continue
		}

		cost := 0
		for j := 0; j < 3; j++ {
			L := 0
			c := 0
			f := false
			for k, v := range b {
				if v == j {
					L += l[k]
					if f {
						c += 10
					} else {
						f = true
					}
				}
			}
			if L > A[j] {
				c += L - A[j]
			} else if L < A[j] {
				c += A[j] - L
			}
			// out(b, c, L, A[j])
			cost += c
		}
		// out("--", cost, "--")
		ans = min(ans, cost)
	}

	out(ans)
}
