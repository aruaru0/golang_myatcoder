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
	sc.Buffer([]byte{}, 1000000)

	S, T := getString(), getString()
	ms := make(map[byte]byte)
	mt := make(map[byte]byte)

	ans := "Yes"
	for i := 0; i < len(S); i++ {
		s := S[i]
		t := T[i]
		v, ok := ms[s]
		if ok == false {
			ms[s] = t
		} else if v != t {
			ans = "No"
			break
		}
		v, ok = mt[t]
		if ok == false {
			mt[t] = s
		} else if v != s {
			ans = "No"
			break
		}
	}
	out(ans)
}
