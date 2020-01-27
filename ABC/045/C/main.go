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

	s := getString()
	l := len(s)

	v := make([]int, l)
	for i := 0; i < l; i++ {
		v[i] = int(s[i] - '0')
	}

	ans := 0
	pat := 1 << uint(l-1)
	for i := 0; i < pat; i++ {
		s := 0
		k := 0
		p := i
		for j := 0; j < l; j++ {
			s = s*10 + v[k]
			if p&1 == 1 {
				ans += s
				s = 0
			}
			k++
			p >>= 1
		}
		ans += s
	}

	out(ans)

}
