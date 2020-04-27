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

func main() {
	sc.Split(bufio.ScanWords)
	s, t := []byte(getString()), []byte(getString())

	S := make([]int, len(s))
	for i := 0; i < len(s); i++ {
		S[i] = int(s[i])
	}

	T := make([]int, len(t))
	for i := 0; i < len(t); i++ {
		T[i] = int(t[i])
	}

	sort.Ints(S)
	sort.Sort(sort.Reverse(sort.IntSlice(T)))

	// out(string(t), string(s))
	L := min(len(S), len(T))
	for i := 0; i < L; i++ {
		if S[i] > T[i] {
			out("No")
			return
		} else if S[i] < T[i] {
			out("Yes")
			return
		}
		// out(i, ":", string(S[i]), string(T[i]))
	}
	if len(T) > L {
		out("Yes")
		return
	}
	out("No")
}
