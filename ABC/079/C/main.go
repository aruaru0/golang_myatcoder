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

func main() {
	sc.Split(bufio.ScanWords)

	x := getString()
	a := make([]int, 4)
	for i, v := range x {
		a[i] = int(v - '0')
	}

	// out(a)
	n := 1 << 3
	for i := 0; i < n; i++ {
		op := "000" + strconv.FormatInt(int64(i), 2)
		op = op[len(op)-3:]
		sum := a[0]
		for i, v := range op {
			if v == '0' {
				sum += a[i+1]
			} else {
				sum -= a[i+1]
			}
		}
		// out(sum)
		if sum == 7 {
			fmt.Print(a[0])
			for i, v := range op {
				if v == '0' {
					fmt.Print("+")
				} else {
					fmt.Print("-")
				}
				fmt.Print(a[i+1])
			}
			out("=7")
			break
		}
	}

}
