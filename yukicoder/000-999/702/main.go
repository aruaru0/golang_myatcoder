package main

import (
	"fmt"
	"sort"
)

var x, y, z, w uint32 = 0, 1, 2, 3

func generate() uint32 {
	t := (x ^ (x << 11))
	x = y
	y = z
	z = w
	w = (w ^ (w >> 19)) ^ (t ^ (t >> 8))
	return w
}

const size = 10000001

func main() {
	// this template is new version.
	// use getI(), getS(), getInts(), getF()
	var seed int
	fmt.Scanf("%d", &seed)
	x = uint32(seed)
	a := make([]uint32, 0)
	l, r := 0, 0
	for i := 0; i < size; i++ {
		w = generate()
		if w < 0x70000000 {
			l++
		} else if w < 0x90000000 {
			a = append(a, w)
		} else {
			r++
		}
	}
	sort.Slice(a, func(i, j int) bool {
		return a[i] < a[j]
	})
	// fmt.Println(len(a), l, r, l+r+len(a))
	fmt.Println(a[size/2-l])
	//	fmt.Println(a[size/2])
}
