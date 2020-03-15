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

func solve(N int) int {
	n := 1
	tbl := []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}
	for i := 0; i < N; i++ {
		n *= tbl[i]
	}

	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	for i := 0; i < n; i++ {
		x := i
		s := make([]int, N)
		skip := false
		for j := 0; j < N; j++ {
			b := x % tbl[j]
			x /= tbl[j]
			s[N-1-j] = b
			if b > N-j {
				skip = true
				break
			}
		}
		if skip {
			continue
		}

		flag := true
		max := 0
		for j := 0; j < N; j++ {
			if s[j] > j {
				flag = false
				break
			}
			if max+1 < s[j] {
				flag = false
				break
			}
			if max < s[j] {
				max = s[j]
			}
		}
		if flag == true {
			for _, v := range s {
				fmt.Fprintf(w, "%c", v+'a')
			}
			fmt.Fprintln(w)
		}
	}
	return 0
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1000000)

	N := getInt()

	solve(N)
}
