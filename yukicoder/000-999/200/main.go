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
	N := getInt()
	A := getInt()
	B := getInts(A)
	C := getInt()
	D := getInts(C)
	sort.Ints(B)
	sort.Ints(D)
	b := make([]int, 0)
	d := make([]int, 0)
	cnt := 0
L0:
	for i := 0; i < N; i++ {
		if len(b) == 0 {
			b = make([]int, A)
			copy(b, B)
		}
		if len(d) == 0 {
			d = make([]int, C)
			copy(d, D)
		}
		// out(i+1, b, d)
		// maxD := d[len(d)-1]
		// p := upperBound(b, maxD)
		// if p < len(b) {
		// 	// win
		// 	b = append(b[:p], b[p+1:]...)
		// 	d = d[:len(d)-1]
		// 	cnt++
		// 	continue
		// }

		for j := 0; j < len(b); j++ {
			maxB := b[j]
			p := lowerBound(d, maxB)
			if p > 0 {
				b = append(b[:j], b[j+1:]...)
				d = append(d[:p-1], d[p:]...)
				cnt++
				continue L0
			}
		}
		b = b[1:]
		d = d[:len(d)-1]
	}
	out(cnt)
}
