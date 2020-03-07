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
	A, B, C, D, E, F := getInt(), getInt(), getInt(), getInt(), getInt(), getInt()

	alim := F / (100 * A)
	blim := F / (100 * B)

	suger := (F * E) / 100
	clim := suger / C
	dlim := suger / D

	max := float64(-1)
	maxw := 0
	maxs := 0

	// 汚いので書き直し候補
	for a := 0; a <= alim; a++ {
		for b := 0; b <= blim; b++ {
			for c := 0; c <= clim; c++ {
				for d := 0; d <= dlim; d++ {
					w := a*A + b*B
					s := c*C + d*D
					f := w*100 + s
					if f == 0 {
						continue
					}
					rate := float64(s*100) / float64(s+w)
					//out(w, f, s, (s*1000)/(s+w))
					if f <= F && s <= E*w && max < rate {
						max = rate
						maxw = w * 100
						maxs = s
					}
				}
			}
		}
	}

	if E == 0 {
		out(A*100, 0)
	} else {
		out(maxw+maxs, maxs)
	}
}
