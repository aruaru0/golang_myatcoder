package main

import (
	"bufio"
	"fmt"
	"os"
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
	sc.Buffer([]byte{}, 1000000)
	s := getString()

	prev := s[0]
	pos := 0
	cnt := 0
	a := make([]int, len(s))
	for i := 0; i < len(s); i++ {
		// out(string(s[i]))
		if s[i] == prev {
			cnt++
		} else if s[i] == 'L' {
			// out("L", i, cnt)
			a[i-1] += cnt/2 + cnt%2
			a[i] += cnt / 2
			prev = 'L'
			pos = i
			cnt = 1
			// out(a)
		} else {
			// out(i, pos, cnt)
			a[pos-1] += cnt / 2
			a[pos] += cnt/2 + cnt%2
			prev = 'R'
			cnt = 1
			pos = 0
			// out(a)
		}
	}
	// out(pos)
	a[pos-1] += cnt / 2
	a[pos] += cnt/2 + cnt%2

	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	for i := 0; i < len(s); i++ {
		fmt.Fprint(w, a[i], " ")
	}
	fmt.Fprintln(w)
}
