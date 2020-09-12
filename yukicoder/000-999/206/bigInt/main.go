package main

import (
	"bufio"
	"fmt"
	"math/big"
	"math/bits"
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

func print(x *big.Int) {
	b := x.Bytes()
	for _, e := range b {
		fmt.Printf("%8.8b ", e)
	}
	fmt.Println()
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1000000)
	L, M, N := getInt(), getInt(), getInt()
	a := big.NewInt(0)
	b := big.NewInt(0)
	for i := 0; i < L; i++ {
		x := getInt()
		a = a.SetBit(a, N-x, 1)
	}
	for i := 0; i < M; i++ {
		x := getInt()
		b = b.SetBit(b, N-x, 1)
	}
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	Q := getInt()
	c := big.NewInt(0)
	for i := 0; i < Q; i++ {
		c = c.Set(a)
		c = c.And(c, b)
		// print(a)
		// print(b)
		// print(c)
		bit := c.Bits()
		cnt := 0
		for _, e := range bit {
			cnt += bits.OnesCount(uint(e))
		}
		fmt.Fprintln(w, cnt)
		b = b.Rsh(b, 1)
	}

}
