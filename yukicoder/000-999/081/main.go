package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"sort"
	"strconv"
	"strings"
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

func getFloat() float64 {
	sc.Scan()
	f, _ := strconv.ParseFloat(sc.Text(), 64)
	return f
}

func main() {
	sc.Split(bufio.ScanWords)
	N := getInt()
	a := make([]int, N)
	b := make([]int, N)
	for i := 0; i < N; i++ {
		s := getString()
		sign := 0
		if s[0] == '-' {
			sign = 1
			s = s[1:]
		}
		sl := strings.Split(s, ".")
		if len(sl) == 1 {
			a[i], _ = strconv.Atoi(sl[0])
		} else {
			sl[1] = sl[1] + "0000000000"
			sl[1] = sl[1][:10]
			a[i], _ = strconv.Atoi(sl[0])
			b[i], _ = strconv.Atoi(sl[1])
		}
		if sign == 1 {
			a[i] = -a[i]
			b[i] = -b[i]
		}
	}
	e10 := big.NewInt(1e10)
	ans := big.NewInt(0)
	for i := 0; i < N; i++ {
		Y := big.NewInt(int64(a[i]))
		Y.Mul(Y, e10)
		ans.Add(ans, Y)
		ans.Add(ans, big.NewInt(int64(b[i])))
	}

	I := big.NewInt(0)
	F := big.NewInt(0)
	sign := 0
	if ans.Sign() < 0 {
		sign = 1
		ans.Sub(big.NewInt(0), ans)
	}
	I.Set(ans)
	F.Set(ans)
	I.Div(I, e10)
	F.Mod(F, e10)
	ansi := I.Int64()
	ansf := F.Int64()
	if sign == 1 {
		fmt.Printf("-%d.%10.10d\n", ansi, ansf)
	} else {
		fmt.Printf("%d.%10.10d\n", ansi, ansf)
	}
}
