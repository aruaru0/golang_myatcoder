package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)
var wr = bufio.NewWriter(os.Stdout)

func out(x ...interface{}) {
	fmt.Fprintln(wr, x...)
}

func getI() int {
	sc.Scan()
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return i
}

func getF() float64 {
	sc.Scan()
	i, e := strconv.ParseFloat(sc.Text(), 64)
	if e != nil {
		panic(e)
	}
	return i
}

func getInts(N int) []int {
	ret := make([]int, N)
	for i := 0; i < N; i++ {
		ret[i] = getI()
	}
	return ret
}

func getS() string {
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

// min for n entry
func nmin(a ...int) int {
	ret := a[0]
	for _, e := range a {
		ret = min(ret, e)
	}
	return ret
}

// max for n entry
func nmax(a ...int) int {
	ret := a[0]
	for _, e := range a {
		ret = max(ret, e)
	}
	return ret
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
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1000000)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	x := getI()

	if x >= len(cnt) {
		out(0, 0)
		return
	}
	out(cnt[x], tot[x])

	// var cnt [35]int
	// var tot [35]int
	// for i := 0; i < (1 << 31); i++ {
	// 	n := bits.OnesCount(uint(i))
	// 	cnt[n]++
	// 	tot[n] += i
	// }
	// for _, e := range cnt {
	// 	fmt.Fprint(wr, e, ",")
	// }
	// out()
	// for _, e := range tot {
	// 	fmt.Fprint(wr, e, ",")
	// }
	// out()
}

var cnt = []int{
	1, 31, 465, 4495, 31465, 169911, 736281, 2629575, 7888725, 20160075, 44352165, 84672315, 141120525, 206253075, 265182525, 300540195, 300540195, 265182525, 206253075, 141120525, 84672315, 44352165, 20160075, 7888725, 2629575, 736281, 169911, 31465, 4495, 465, 31, 1, 0, 0, 0,
}
var tot = []int{
	0, 2147483647, 64424509410, 934155386445, 8718783606820, 58851789346035, 306029304599382, 1275122102497425, 4371847208562600, 12569060724617475, 30724370660176050, 64521178386369705, 117311233429763100, 185742786263791575, 257182319442172950, 312292816465495725, 333112337563195440, 312292816465495725, 257182319442172950, 185742786263791575, 117311233429763100, 64521178386369705, 30724370660176050, 12569060724617475, 4371847208562600, 1275122102497425, 306029304599382, 58851789346035, 8718783606820, 934155386445, 64424509410, 2147483647, 0, 0, 0,
}
