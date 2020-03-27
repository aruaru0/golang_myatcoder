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

	t := getString()

	ok := true
	for i := 0; i < len(t); i++ {
		if t[0] != t[i] {
			ok = false
			break
		}
	}
	if ok {
		out(0)
		return
	}

	N := len(t)
	cnt := 1
	a := make([]byte, 0)
	a = append(a, t[N-1])
	N--
LOOP:
	for i := 0; i < len(t); i++ {
		a = append(a, t[N-1])
		//out(i, N, "-------", string(a))
		for _, v := range a {
			ok = true
			for j := 0; j < N; j++ {
				check := false
				for k := 0; k <= cnt; k++ {
					//out(j, k, string(t[k]))
					if t[j+k] == v {
						check = true
						break
					}
				}
				if check == false {
					ok = false
					break
				}
			}
			if ok == true {
				break LOOP
			}
		}
		N--
		cnt++
	}

	out(cnt)
}
