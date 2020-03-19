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

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func asub(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

func main() {
	sc.Split(bufio.ScanWords)

	Deg, Dis := getInt(), getInt()
	deg := Deg * 10
	dir := []string{
		"NNE", "NE", "ENE",
		"E", "ESE", "SE", "SSE",
		"S", "SSW", "SW", "WSW",
		"W", "WNW", "NW", "NNW", "N"}

	n := 0
	for i := 1125; i < 34875; i += 2250 {
		if i <= deg && deg < i+2250 {
			break
		}
		n++
	}

	r := []int{
		15,
		93,
		201,
		327,
		477,
		645,
		831,
		1029,
		1245,
		1467,
		1707,
		1959,
	}

	m := 0
	dis := Dis
	for _, v := range r {
		if dis < v {
			break
		}
		m++
	}

	if m == 0 {
		out("C", m)
	} else {
		out(dir[n], m)
	}
}
