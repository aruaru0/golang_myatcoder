package main

import (
	"bufio"
	"fmt"
	"math"
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

var x, y, z float64
var X, Y, Z []float64

func calcL(a, b, c int) float64 {
	bx := X[b] - X[a]
	by := Y[b] - Y[a]
	bz := Z[b] - Z[a]
	cx := X[c] - X[a]
	cy := Y[c] - Y[a]
	cz := Z[c] - Z[a]

	A := by*cz - cy*bz
	B := bz*cx - cz*bx
	C := bx*cy - cx*by
	D := -(A*X[a] + B*Y[a] + C*Z[a])

	ret := math.Abs(A*x + B*y + C*z + D)
	ret /= math.Sqrt(A*A + B*B + C*C)

	return ret
}

func getf() float64 {
	sc.Scan()
	f, _ := strconv.ParseFloat(sc.Text(), 64)
	return f
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1000000)
	N := getInt()
	x, y, z = getf(), getf(), getf()
	X = make([]float64, N)
	Y = make([]float64, N)
	Z = make([]float64, N)
	for i := 0; i < N; i++ {
		X[i], Y[i], Z[i] = getf(), getf(), getf()
	}

	ans := 0.0

	for i := 0; i < N; i++ {
		for j := i + 1; j < N; j++ {
			for k := j + 1; k < N; k++ {
				ret := calcL(i, j, k)
				ans += ret
			}
		}
	}
	out(ans)
}
