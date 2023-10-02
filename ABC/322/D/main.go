package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)
var wr = bufio.NewWriter(os.Stdout)

func out(x ...interface{}) {
	fmt.Fprintln(wr, x...)
}

func getI() int {
	sc.Scan()
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return i
}

func getF() float64 {
	sc.Scan()
	i, e := strconv.ParseFloat(sc.Text(), 64)
	if e != nil {
		panic(e)
	}
	return i
}

func getInts(N int) []int {
	ret := make([]int, N)
	for i := 0; i < N; i++ {
		ret[i] = getI()
	}
	return ret
}

func getS() string {
	sc.Scan()
	return sc.Text()
}

// min, max, asub, absなど基本関数
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

// min for n entry
func nmin(a ...int) int {
	ret := a[0]
	for _, e := range a {
		ret = min(ret, e)
	}
	return ret
}

// max for n entry
func nmax(a ...int) int {
	ret := a[0]
	for _, e := range a {
		ret = max(ret, e)
	}
	return ret
}

func chmin(a *int, b int) bool {
	if *a < b {
		return false
	}
	*a = b
	return true
}

func chmax(a *int, b int) bool {
	if *a > b {
		return false
	}
	*a = b
	return true
}

func asub(a, b int) int {
	if a > b {
		return a - b
	}
	return b - a
}

func abs(a int) int {
	if a >= 0 {
		return a
	}
	return -a
}

func lowerBound(a []int, x int) int {
	idx := sort.Search(len(a), func(i int) bool {
		return a[i] >= x
	})
	return idx
}

func upperBound(a []int, x int) int {
	idx := sort.Search(len(a), func(i int) bool {
		return a[i] > x
	})
	return idx
}

func f(sx, sy int, target [4][4]bool, table [10][10]bool) (bool, [10][10]bool) {
	var ret [10][10]bool
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			ret[i][j] = table[i][j]
		}
	}

	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			px, py := sx+i, sy+j
			if target[i][j] && table[px][py] {
				return false, ret
			}
			if target[i][j] {
				ret[px][py] = target[i][j]
			}
		}
	}

	return true, ret
}

func r(s [4]string, rot int) [4][4]bool {
	var ret [4][4]bool
	if rot == 0 {
		for i := 0; i < 4; i++ {
			for j := 0; j < 4; j++ {
				if s[i][j] == '#' {
					ret[i][j] = true
				}
			}
		}
	} else if rot == 1 {
		for i := 0; i < 4; i++ {
			for j := 0; j < 4; j++ {
				if s[i][j] == '#' {
					ret[j][3-i] = true
				}
			}
		}
	} else if rot == 2 {
		for i := 0; i < 4; i++ {
			for j := 0; j < 4; j++ {
				if s[i][j] == '#' {
					ret[3-i][3-j] = true
				}
			}
		}
	} else if rot == 3 {
		for i := 0; i < 4; i++ {
			for j := 0; j < 4; j++ {
				if s[i][j] == '#' {
					ret[3-j][i] = true
				}
			}
		}
	}
	return ret
}

func ini() [10][10]bool {
	var ret [10][10]bool
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			if 3 <= i && i <= 6 && 3 <= j && j <= 6 {
			} else {
				ret[i][j] = true
			}
		}
	}
	return ret
}

func display(a [10][10]bool) {
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			x := 0
			if a[i][j] {
				x = 1
			}
			fmt.Fprint(wr, x)
		}
		out()
	}
}

func display_rot(a [4][4]bool) {
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			x := 0
			if a[i][j] {
				x = 1
			}
			fmt.Fprint(wr, x)
		}
		out()
	}
}

func main() {
	defer wr.Flush()
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, math.MaxInt32)
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	p := make([][4]string, 3)
	for i := 0; i < 3; i++ {
		for j := 0; j < 4; j++ {
			p[i][j] = getS()
		}
	}

	table0 := ini()
	for i0 := 0; i0 <= 6; i0++ {
		for j0 := 0; j0 <= 6; j0++ {
			for k0 := 0; k0 < 4; k0++ {
				target0 := r(p[0], k0)
				ret, table1 := f(i0, j0, target0, table0)
				if ret == false {
					continue
				}
				for i1 := 0; i1 <= 6; i1++ {
					for j1 := 0; j1 <= 6; j1++ {
						for k1 := 0; k1 < 4; k1++ {
							target1 := r(p[1], k1)
							ret, table2 := f(i1, j1, target1, table1)
							if ret == false {
								continue
							}
							for i2 := 0; i2 <= 6; i2++ {
								for j2 := 0; j2 <= 6; j2++ {
									for k2 := 0; k2 < 4; k2++ {
										target2 := r(p[2], k2)
										ret, table3 := f(i2, j2, target2, table2)
										if ret == true {
											sum := 0
											for i := 0; i < 10; i++ {
												for j := 0; j < 10; j++ {
													if table3[i][j] {
														sum++
													}
												}
											}
											if sum == 100 {
												// display_rot(target0)
												// display(table1)
												// display_rot(target1)
												// display(table2)
												// display_rot(target2)
												// display(table3)
												// out(i0, j1, k0, ":", i1, j1, k1, ":", i2, j2, k2)
												out("Yes")
												return
											}
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}

	out("No")
}
