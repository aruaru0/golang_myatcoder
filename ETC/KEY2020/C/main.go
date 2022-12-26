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

	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	N, K, S := getInt(), getInt(), getInt()
	S2 := S + 1
	if S == 1000000000 {
		S2 = 1
	}
	for i := 0; i < N; i++ {
		if i < K {
			fmt.Fprint(w, S, " ")
		} else {
			fmt.Fprint(w, S2, " ")
		}
	}
	fmt.Fprintln(w)
}
