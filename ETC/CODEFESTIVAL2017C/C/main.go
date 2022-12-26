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

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 100100)

	s := getString()

	l := 0
	r := len(s) - 1
	//n := len(s)
	cnt := 0
	ans := 0
	for {
		//out(l, r, cnt, string(s[l]), string(s[r]))
		if l >= r {
			break
		}
		if s[l] == s[r] {
			l++
			r--
			cnt += 2
			//out(string(s[l]))
		} else if s[l] == 'x' && s[r] != 'x' {
			//out("l++")
			l++
			cnt++
			ans++
		} else if s[l] != 'x' && s[r] == 'x' {
			//out("r--")
			r--
			cnt++
			ans++
		} else {
			out(-1)
			return
		}
	}
	//out(l + r)
	out(ans)
}
