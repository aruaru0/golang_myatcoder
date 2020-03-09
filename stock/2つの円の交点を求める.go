import "math"

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
