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

var C [51][51]int // C[n][k] -> nCk

func combTable(N int) {
	for i := 1; i <= N; i++ {
		for j := 0; j <= N; j++ {
			if j == 0 || j == i {
				C[i][j] = 1
			} else {
				C[i][j] = C[i-1][j-1] + C[i-1][j]
			}
		}
	}
}

// 残念ながら、解説を写経しています（結構面倒）
func main() {
	sc.Split(bufio.ScanWords)

	N, A, B := getInt(), getInt(), getInt()
	combTable(N)
	v := make([]int, N)
	for i := 0; i < N; i++ {
		v[i] = getInt()
	}
	sort.Sort(sort.Reverse(sort.IntSlice(v)))

	sum := 0
	for i := 0; i < A; i++ {
		sum += v[i]
	}
	ave := float64(sum) / float64(A)
	fmt.Printf("%f\n", ave)

	a, apos := 0, 0
	for i := 0; i < N; i++ {
		if v[i] == v[A-1] {
			a++
			if i < A {
				apos++
			}
		}
	}

	cnt := 0
	if apos == A {
		for apos = A; apos <= B; apos++ {
			cnt += C[a][apos]
		}
	} else {
		cnt += C[a][apos]
	}

	out(cnt)
}
