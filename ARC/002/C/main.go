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

type Data struct {
	key string
	cnt int
}

func nextPermutation(x sort.Interface) bool {
	n := x.Len() - 1
	if n < 1 {
		return false
	}
	j := n - 1
	for ; !x.Less(j, j+1); j-- {
		if j == 0 {
			return false
		}
	}
	l := n
	for !x.Less(j, l) {
		l--
	}
	x.Swap(j, l)
	for k, l := j+1, n; k < l; {
		x.Swap(k, l)
		k++
		l--
	}
	return true
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1000000)

	N := getInt()
	s := getString()

	ans := 100100
	t := []byte{'A', 'B', 'X', 'Y'}
	pat := 4 * 4 * 4 * 4
	for p := 0; p < pat; p++ {
		pp := p
		L := string(t[pp%4]) + string(t[(pp/4)%4])
		pp /= 16
		R := string(t[pp%4]) + string(t[(pp/4)%4])

		d := make([]int, N+1)
		for i := 1; i <= N; i++ {
			d[i] = i
		}

		for i := 1; i <= N; i++ {
			if i > 1 {
				if s[i-2] == L[0] && s[i-1] == L[1] {
					d[i] = min(d[i-2]+1, d[i])
				} else if s[i-2] == R[0] && s[i-1] == R[1] {
					d[i] = min(d[i-2]+1, d[i])
				} else {
					d[i] = min(d[i-1]+1, d[i])
				}
			} else {
				d[i] = d[i-1] + 1
			}
		}

		ans = min(ans, d[N])
	}
	out(ans)
}
