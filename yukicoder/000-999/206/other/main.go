package main

import (
	"bufio"
	"fmt"
	"math"
	"math/cmplx"
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

const n = 1 << 18

func FFT(x []complex128, n int) []complex128 {
	y := fft(x, n)
	for i := 0; i < n; i++ {
		y[i] = y[i] / complex(float64(n), 0.0)
	}
	return y
}

func IFFT(x []complex128, n int) []complex128 {
	y := make([]complex128, n)
	for i := 0; i < n; i++ {
		y[i] = cmplx.Conj(x[i])
	}
	y = fft(y, n)
	for i := 0; i < n; i++ {
		y[i] = cmplx.Conj(y[i])
	}
	return y
}

func fft(a []complex128, n int) []complex128 {
	x := make([]complex128, n)
	copy(x, a)

	j := 0
	for i := 0; i < n; i++ {
		if i < j {
			x[i], x[j] = x[j], x[i]
		}
		m := n / 2
		for {
			if j < m {
				break
			}
			j = j - m
			m = m / 2
			if m < 2 {
				break
			}
		}
		j = j + m
	}
	kmax := 1
	for {
		if kmax >= n {
			return x
		}
		istep := kmax * 2
		for k := 0; k < kmax; k++ {
			theta := complex(0.0, -1.0*math.Pi*float64(k)/float64(kmax))
			for i := k; i < n; i += istep {
				j := i + kmax
				temp := x[j] * cmplx.Exp(theta)
				x[j] = x[i] - temp
				x[i] = x[i] + temp
			}
		}
		kmax = istep
	}
}

func multPoly(P, Q []complex128) []complex128 {
	P = FFT(P, n)
	Q = FFT(Q, n)
	for i := 0; i < len(P); i++ {
		P[i] = P[i] * Q[i]
	}
	return IFFT(P, n)
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1000000)

	ac := make([]complex128, n)
	bc := make([]complex128, n)

	L, M, N := getInt(), getInt(), getInt()
	for i := 0; i < L; i++ {
		x := getInt()
		ac[x] += 1
	}
	for i := 0; i < M; i++ {
		x := getInt()
		bc[N-x] += 1
	}
	cc := multPoly(ac, bc)

	for i := 0; i < n; i++ {
		cc[i] /= 1.0 / n
	}

	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()
	Q := getInt()
	for i := 0; i < Q; i++ {
		fmt.Fprintln(w, (int)(real(cc[i+N])+0.1))
	}
}
