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

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1000000)
	N := getInt()
	D := make([]int, 0)
	C := make([]int, 0)
	H := make([]int, 0)
	S := make([]int, 0)
	card := "A23456789TJQK"
	for i := 0; i < N; i++ {
		s := getString()
		v := 0
		for i, e := range card {
			if byte(e) == s[1] {
				v = i
				break
			}
		}
		switch s[0] {
		case 'D':
			D = append(D, v)
		case 'C':
			C = append(C, v)
		case 'H':
			H = append(H, v)
		case 'S':
			S = append(S, v)
		}
	}
	sort.Ints(D)
	sort.Ints(C)
	sort.Ints(H)
	sort.Ints(S)
	for _, e := range D {
		fmt.Printf("D%c ", card[e])
	}
	for _, e := range C {
		fmt.Printf("C%c ", card[e])
	}
	for _, e := range H {
		fmt.Printf("H%c ", card[e])
	}
	for _, e := range S {
		fmt.Printf("S%c ", card[e])
	}
	out()
}
