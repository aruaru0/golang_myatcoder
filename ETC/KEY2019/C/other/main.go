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

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1000000)

	N := getInt()
	a := make([]int, N)
	b := make([]int, N)
	for i := 0; i < N; i++ {
		a[i] = getInt()
	}
	for i := 0; i < N; i++ {
		b[i] = getInt()
	}

	sum := 0
	d := make([]int, N)
	for i := 0; i < N; i++ {
		d[i] = a[i] - b[i]
		sum += d[i]
	}

	// out(d, sum)

	if sum < 0 {
		out(-1)
		return
	}

	sort.Ints(d)
	cnt := 0
	pos := len(d)
	rest := 0
	for i := 0; i < N; i++ {
		x := d[i]
		// out(x, rest, "cnt", cnt)
		if x >= 0 {
			// out("break", i)
			cnt += i
			break
		}
		for x < 0 {
			// out("---")
			if rest >= -x {
				rest += x
				break
			} else {
				x += rest
				pos--
				cnt++
				rest = d[pos]
				// out("rest", rest, cnt)
			}
		}
	}
	out(cnt)
}
