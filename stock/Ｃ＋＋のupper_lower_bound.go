type pair struct {
	f, s int
}

// P :
type P []pair

func (p P) Len() int {
	return len(p)
}

func (p P) Swap(i, j int) {
	p[i], p[j] = p[j], p[i]
}

func (p P) Less(i, j int) bool {
	if p[i].f == p[j].f {
		return p[i].s < p[j].s
	}
	return p[i].f < p[j].f
}

func lower_bound(a P, x pair) int {
	l := 0
	r := len(a)
	for l <= r {
		m := (l + r) / 2
		if len(a) == m {
			break
		}
		if a[m].f >= x.f || (a[m].f == x.f && a[m].s >= x.s) {
			r = m - 1
		} else {
			l = m + 1
		}
	}
	return l
}

func upper_bound(a P, x pair) int {
	l := 0
	r := len(a)
	for l <= r {
		m := (l + r) / 2
		if len(a) == m {
			break
		}
		if a[m].f <= x.f || (a[m].f == x.f && a[m].s <= x.s) {
			l = m + 1
		} else {
			r = m - 1
		}
	}
	return l
}