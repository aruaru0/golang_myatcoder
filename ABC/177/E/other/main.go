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

var pfs map[int]int
var prime []int

const size = 1000100

func initPrime() {
	prime = make([]int, size+1)
	for i := 0; i <= size; i++ {
		prime[i] = i
	}
	for i := 2; i*i <= size; i++ {
		if prime[i] == i {
			for j := i; j <= size; j += i {
				prime[j] = i
			}
		}
	}
}

// PfsMap : 素因数分解し、マップを作成
func PfsMap(n int) {
	for n >= 2 {
		p := prime[n]
		pfs[p]++
		for n%p == 0 {
			n /= p
		}
	}
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

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1000000)
	N := getInt()
	a := getInts(N)

	// N := 100
	// a := make([]int, N)
	// for i := 0; i < N; i++ {
	// 	a[i] = rand.Intn(1000000)
	// }

	initPrime()
	pfs = make(map[int]int)

	sc := a[0]
	PfsMap(a[0])
	for i := 1; i < N; i++ {
		PfsMap(a[i])
		sc = GCD(sc, a[i])
	}
	// out(sc, pfs, a)
	flg := true
	for _, e := range pfs {
		if e != 1 {
			flg = false
			break
		}
	}
	if flg == true {
		out("pairwise coprime")
		return
	}
	if sc == 1 {
		out("setwise coprime")
		return
	}
	out("not coprime")

}
