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

func solve(h int, w []int) {
	ans := abs(w[0] - h)
	for i := 0; i < len(w); i++ {
		ans = min(ans, abs(h-w[i]))
	}
	out(ans)
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1000000)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N, M := getI(), getI()
	h := getInts(N)
	w := getInts(M)

	if N == 1 {
		solve(h[0], w)
		return
	}

	sort.Ints(h)
	// out(h)
	dL := make([]int, N/2)
	dR := make([]int, N/2)
	for i := 1; i < N; i += 2 {
		dL[i/2] = h[i] - h[i-1]
		dR[i/2] = h[N-i] - h[N-i-1]
	}
	sL := make([]int, N/2)
	sR := make([]int, N/2)
	sL[0] = dL[0]
	sR[0] = dR[0]
	for i := 1; i < N/2; i++ {
		sL[i] = sL[i-1] + dL[i]
		sR[i] = sR[i-1] + dR[i]
	}
	sL = append([]int{0}, sL...)
	sR = append([]int{0}, sR...)
	// out(dL, dR)
	// out(sL, sR)

	ans := int(1e18)
	for i := 0; i < M; i++ {
		pos := lowerBound(h, w[i])
		// out("----", h, pos, w[i])
		if pos%2 == 0 {
			pair := pos
			n := N / 2
			tot := sL[pos/2] + sR[n-pos/2] + abs(w[i]-h[pair])
			ans = min(ans, tot)
			// out("left", pos/2, "right", n-pos/2, "tot", tot, w[i], "pair", h[pair])
		} else {
			pair := pos - 1
			n := N / 2
			tot := sL[pos/2] + sR[n-pos/2] + abs(w[i]-h[pair])
			ans = min(ans, tot)
			// out("left", pos/2, "right", n-pos/2, "tot", tot, w[i], "pair", h[pair])
		}
	}
	out(ans)
}
