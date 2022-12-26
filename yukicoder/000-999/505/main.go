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

const inf = int(1e18)

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1000000)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	N := getI()
	a := getInts(N)

	N0 := min(12, N)
	N1 := max(0, N-N0)
	x := 1
	for i := 0; i < N0-1; i++ {
		x *= 4
	}

	m := make(map[int]bool, 0)
	ans := -inf
	for i := 0; i < x; i++ {
		n := i
		tot := a[0]
		for j := 1; j < N0; j++ {
			op := n % 4
			switch op {
			case 0:
				tot += a[j]
			case 1:
				tot -= a[j]
			case 2:
				tot *= a[j]
			case 3:
				if a[j] == 0 {
					tot += a[j]
				} else {
					tot /= a[j]
				}
			}
			n /= 4
		}
		m[tot] = true
		ans = max(ans, tot)
	}

	if N1 == 0 {
		out(ans)
		return
	}

	x = 1
	for i := 0; i < N1; i++ {
		x *= 4
	}

	ans = -inf
	for e := range m {
		for i := 0; i < x; i++ {
			n := i
			tot := e
			// fmt.Print(e)
			for j := N0; j < N; j++ {
				op := n % 4
				switch op {
				case 0:
					// fmt.Print("+", a[j])
					tot += a[j]
				case 1:
					// fmt.Print("-", a[j])
					tot -= a[j]
				case 2:
					// fmt.Print("*", a[j])
					tot *= a[j]
				case 3:
					if a[j] == 0 {
						// fmt.Print(".", a[j])
						tot += a[j]
					} else {
						// fmt.Print("/", a[j])
						tot /= a[j]
					}
				}
				n /= 4
			}
			// fmt.Println("=", tot)
			ans = max(ans, tot)
		}
	}
	out(ans)
}
