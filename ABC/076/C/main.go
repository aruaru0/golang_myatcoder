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

func match(s, t string) bool {
	ret := true
	for i := 0; i < len(s); i++ {
		if s[i] == '?' {
			continue
		}
		if s[i] != t[i] {
			ret = false
			break
		}
	}
	return ret
}

func main() {
	sc.Split(bufio.ScanWords)
	s := getString()
	t := getString()

	N := len(s)
	M := len(t)
	pos := -1
	// search
	for i := 0; i <= N-M; i++ {
		f := match(s[i:i+M], t)
		// out(s[i:i+M], f)
		if f == true {
			pos = i
		}
	}
	// out(pos)
	if pos == -1 {
		out("UNRESTORABLE")
		return
	}
	ans := []byte(s[:pos] + t + s[pos+M:])
	for i := 0; i < N; i++ {
		if ans[i] == '?' {
			ans[i] = 'a'
		}
	}
	out(string(ans))
}
