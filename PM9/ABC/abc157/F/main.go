package main

import (
	"bufio"
	"fmt"
	"math"
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

// ２つの円の交点を求める
type point struct {
	x, y float64
}

const eps = 1e-9

func circleCrossPt(p0, p1 point, r1, r2 float64) ([]point, int) {
	dx := p1.x - p0.x
	dy := p1.y - p0.y
	l2 := dx*dx + dy*dy
	a := (l2 + r1*r1 - r2*r2) / 2.0
	D := l2*r1*r1 - a*a
	if D < 0 || l2 == 0 {
		return []point{}, 0
	}
	D = math.Sqrt(D)
	return []point{
		point{(a*dx+D*dy)/l2 + p0.x, (a*dy-D*dx)/l2 + p0.y},
		point{(a*dx-D*dy)/l2 + p0.x, (a*dy+D*dx)/l2 + p0.y}}, 2
}

func abs(a float64) float64 {
	if a < 0 {
		return -a
	}
	return a
}

func equal(a, b float64) bool {
	return abs(a-b) < eps
}

func max(a, b float64) float64 {
	if a > b {
		return a
	} else {
		return b
	}
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1000000)

	//out(circleCrossPt(point{0, -10}, point{0, 10}, 10, 10))
	//panic(-1)
	N, K := getInt(), getInt()
	p := make([]point, N)
	c := make([]float64, N)
	for i := 0; i < N; i++ {
		x, y, z := getInt(), getInt(), getInt()
		p[i] = point{float64(x), float64(y)}
		c[i] = float64(z)
	}

	//	out(N, K, p, c)

	L := float64(0)
	R := float64(400000)

	for n := 0; n < 100; n++ {
		T := (L + R) / 2.0
		ap := make([]point, 0)
		// 中心を登録
		for i := 0; i < N; i++ {
			ap = append(ap, p[i])
		}
		// 交点を登録
		for i := 0; i < N; i++ {
			for j := i + 1; j < N; j++ {
				pos, num := circleCrossPt(p[i], p[j], T/c[i], T/c[j])
				switch num {
				case 1:
					ap = append(ap, pos[0])
				case 2:
					ap = append(ap, pos[0])
					ap = append(ap, pos[1])
				}
			}
		}
		ans := false
		for _, v := range ap {
			cnt := 0
			for i := 0; i < N; i++ {
				dx := p[i].x - v.x
				dy := p[i].y - v.y
				r := math.Sqrt(dx*dx+dy*dy) * c[i]
				if r <= T+eps {
					cnt++
				}
			}
			if cnt >= K {
				ans = true
				break
			}
		}
		if ans == true {
			R = T
		} else {
			L = T
		}
		//		out(ans, L, R)
	}
	out(L)
}
