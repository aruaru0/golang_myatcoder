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

func modpow(a, p, m int) int {
	if a == 0 {
		return 0
	}
	if p == 0 {
		return 1
	}
	if p%2 == 0 {
		tmp := modpow(a, p/2, m)
		return tmp * tmp % m
	}
	return a * modpow(a, p-1, m) % m
}

const Mmax = 2000

var phi [Mmax + 20]int

func AtoB(a, b, upbound int) int {
	if b == 0 || a == 1 {
		return min(upbound, 1)
	}
	if b == 1 {
		return min(upbound, a)
	}
	if b == 2 {
		if a > 10 {
			return upbound
		}
		return min(modpow(a, a, 1<<60), upbound)
	}
	if a == 2 && b == 3 {
		return min(upbound, 1<<4)
	}
	if a == 3 && b == 3 {
		return min(upbound, modpow(3, 27, 1<<60))
	}
	if a == 2 && b == 4 {
		return min(upbound, 1<<16)
	}
	return upbound
}

func H4M(a, b, m int) int {
	if m == 1 {
		return 0
	}
	if a == 1 || b == 0 {
		return 1
	}

	t := AtoB(a, b-1, m+1)
	if t <= m {
		return modpow(a, t, m)
	}
	return modpow(a, m+(H4M(a, b-1, phi[m])-m)%phi[m], m)
}

// GCD : greatest common divisor (GCD) via Euclidean algorithm
func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}
func GetPhi() {
	for x := 1; x <= Mmax; x++ {
		for i := 1; i <= x; i++ {
			if GCD(i, x) == 1 {
				phi[x]++
			}
		}
	}

}

func main() {
	sc.Split(bufio.ScanWords)
	A, N, M := getInt(), getInt(), getInt()
	GetPhi()
	out(H4M(A, N, M))
}
