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

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1000000)

	s := getString()
	K := getInt()
	sub := make([]string, 0)
	for i := 0; i < len(s); i++ {
		n := min(i+5, len(s))
		for j := i + 1; j <= n; j++ {
			sub = append(sub, s[i:j])
		}
	}

	sort.Strings(sub)
	prev := ""
	cnt := 0
	for _, v := range sub {
		if v == prev {
			continue
		}
		prev = v
		cnt++
		if cnt == K {
			break
		}
	}
	out(prev)
}
