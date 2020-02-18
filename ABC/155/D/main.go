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

const inf = int(2e18)

func main() {
	sc.Split(bufio.ScanWords)

	N, K := getInt(), getInt()
	a := make([]int, N)
	for i := 0; i < N; i++ {
		a[i] = getInt()
	}

	sort.Ints(a)

	L := -inf
	R := inf
	for L+1 != R {
		M := (L + R) / 2
		cnt := 0
		for i := 0; i < N; i++ {
			if a[i] < 0 {
				l := i
				r := N
				for l+1 != r {
					m := (l + r) / 2
					if a[i]*a[m] <= M {
						r = m
					} else {
						l = m
					}
				}
				cnt += N - r
			} else {
				l := i
				r := N
				for l+1 != r {
					m := (l + r) / 2
					if a[i]*a[m] <= M {
						l = m
					} else {
						r = m
					}
				}
				cnt += l - i
			}
		}
		if cnt < K {
			L = M
		} else {
			R = M
		}
	}
	out(R)
}
