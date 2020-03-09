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

var memo [301][102]int
var cnt = 1

func rec(i, n, ng1, ng2, ng3 int) bool {
	if memo[i][n] != 0 {
		if memo[i][n] > 0 {
			return true
		} else {
			return false
		}
	}

	if i == ng1 || i == ng2 || i == ng3 {
		return false
	}

	if n >= 100 {
		return false
	}
	if i == 1 || i == 2 || i == 3 {
		return true
	}

	ret0 := rec(i-1, n+1, ng1, ng2, ng3)
	ret1 := rec(i-2, n+1, ng1, ng2, ng3)
	ret2 := rec(i-3, n+1, ng1, ng2, ng3)
	ret := ret0 || ret1 || ret2
	if ret {
		memo[i][n] = 1
	} else {
		memo[i][n] = -1
	}

	return ret
}

func main() {
	sc.Split(bufio.ScanWords)
	N, ng1, ng2, ng3 := getInt(), getInt(), getInt(), getInt()

	res := rec(N, 0, ng1, ng2, ng3)
	if res {
		out("YES")
	} else {
		out("NO")
	}
}
