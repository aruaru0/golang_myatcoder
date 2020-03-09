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
	Q := getInt()

	S := make([]byte, 100000*7)
	start := 300000
	end := start + len(s)
	for i := 0; i < len(s); i++ {
		S[i+start] = s[i]
	}

	rev := false
	for i := 0; i < Q; i++ {
		t := getInt()
		if t == 1 {
			rev = !rev
		} else {
			f, c := getInt(), getString()[0]
			if f == 1 {
				if rev == true {
					S[end] = c
					end++
				} else {
					start--
					S[start] = c
				}
			} else {
				if rev == true {
					start--
					S[start] = c
				} else {
					S[end] = c
					end++
				}
			}
		}
	}

	len := end - start
	ans := make([]byte, len)
	if rev == false {
		for i := 0; i < len; i++ {
			ans[i] = S[start+i]
		}
	} else {
		for i := 0; i < len; i++ {
			ans[i] = S[end-1-i]
		}
	}
	out(string(ans))
}
