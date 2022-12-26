package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"sync"
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

const div = 60

func main() {
	sc.Split(bufio.ScanWords)
	var A [210000]uint
	var C [21000]uint

	N, S, X, Y, Z := getInt(), getInt(), getInt(), getInt(), getInt()
	pre := S
	for i := 0; i < N; i++ {
		abase := i / div
		apos := i % div
		if pre%2 != 0 {
			A[abase] |= 1 << apos
		}
		pre = (X*pre + Y) % Z
	}

	// fmt.Printf("%64.64b\n", A[0])
	Q := getInt()
	for l := 0; l < Q; l++ {
		sk, tk, uk, vk := getInt()-1, getInt()-1, getInt()-1, getInt()-1
		sb := sk / div
		sp := sk % div
		tb := tk / div
		tp := tk % div
		ub := uk / div
		up := uk % div
		vb := vk / div
		// vp := vk % div

		cp := up
		cb := 0
		C[cb] = 0
		for i := sb; i <= tb; i++ {
			target := div - 1
			if i == tb {
				target = tp
			}
			for sp <= target {
				LEN := min(target-sp+1, div-cp)
				C[cb] |= (((1 << LEN) - 1) & (A[i] >> sp)) << cp

				sp += LEN
				cp += LEN
				if cp >= div {
					cb++
					C[cb] = 0
					cp = 0
				}
			}
			sp = 0
		}

		wg := sync.WaitGroup{}
		f := func(from, to int) {
			for i := from; i <= to; i++ {
				A[i] ^= C[i-ub]
			}
			wg.Done()
		}
		wg.Add(2)
		NN := (vb - ub) / 2
		go f(ub, ub+NN)
		go f(ub+NN+1, vb)
		wg.Wait()
		// for i := ub; i <= vb; i++ {
		// 	A[i] ^= C[i-ub]
		// }
	}

	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()
	for i := 0; i < N; i++ {
		b := i / div
		p := i % div
		if (A[b]>>p)&1 != 0 {
			fmt.Fprint(w, "O")
		} else {
			fmt.Fprint(w, "E")
		}
	}
	fmt.Fprintln(w)
}
