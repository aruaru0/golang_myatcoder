package main

import (
	"bufio"
	"fmt"
	"os"
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

	t := "keyence"

	for i := 0; i <= len(t); i++ {
		f0 := strings.HasPrefix(s, t[:i])
		f1 := strings.HasSuffix(s, t[i:])
		// out(t[:i], t[i:], s, f0, f1)
		if f0 && f1 {
			out("YES")
			return
		}
	}
	out("NO")
}