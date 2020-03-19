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

func conv(c byte) int {
	ret := -1
	if c >= 'A' && c <= 'Z' {
		c = c - 'A' + 'a'
	}

	switch {
	case c == 'b' || c == 'c':
		ret = 1
	case c == 'd' || c == 'w':
		ret = 2
	case c == 't' || c == 'j':
		ret = 3
	case c == 'f' || c == 'q':
		ret = 4
	case c == 'l' || c == 'v':
		ret = 5
	case c == 's' || c == 'x':
		ret = 6
	case c == 'p' || c == 'm':
		ret = 7
	case c == 'h' || c == 'k':
		ret = 8
	case c == 'n' || c == 'g':
		ret = 9
	case c == 'z' || c == 'r':
		ret = 0
	}
	return ret
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1000000)

	N := getInt()
	ans := ""
	for i := 0; i < N; i++ {
		s := getString()
		flg := false
		for _, v := range s {
			r := conv(byte(v))
			if r != -1 {
				ans += string(r + '0')
				flg = true
			}
		}
		if flg == true && i != N-1 {
			ans += string(" ")
		}
	}

	t := len(ans)
	for i := len(ans) - 1; i >= 0; i-- {
		if ans[i] == ' ' {
			t = i
		} else {
			break
		}
	}
	out(ans[:t])
}
