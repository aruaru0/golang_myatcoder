package main

import "fmt"

// SegmentTree :
type SegmentTree struct {
	offset int
	inf    interface{}
	data   []interface{}
	merge  Merger
}

// Merger :
type Merger func(a, b interface{}) interface{}

// InitSegmentTree :
func InitSegmentTree(a []interface{}, inf interface{}, merge Merger) *SegmentTree {
	n := len(a)
	size := 1
	for size < n {
		size *= 2
	}
	data := make([]interface{}, size*2)
	for j, i := 0, size-1; i < size+n-1; i++ {
		data[i] = a[j]
		j++
	}
	for i := size + n - 1; i < size*2; i++ {
		data[i] = inf
	}

	for i := size - 2; i >= 0; i-- {
		data[i] = merge(data[i*2+1], data[i*2+2])
	}

	return &SegmentTree{
		inf:    inf,
		offset: size,
		data:   data,
		merge:  merge,
	}
}

// GetRange :
func (tree *SegmentTree) GetRange(from, to int) interface{} {
	return tree.getRange(from, to, 0, 0, tree.offset)
}

// getRange :
func (tree *SegmentTree) getRange(from, to, index, left, right int) interface{} {
	if to <= left || right <= from {
		return tree.inf
	}
	if from <= left && right <= to {
		return tree.data[index]
	}
	med := (left + right) / 2
	lvalue := tree.getRange(from, to, index*2+1, left, med)
	rvalue := tree.getRange(from, to, index*2+2, med, right)

	return tree.merge(lvalue, rvalue)
}

// UpdateAt :
func (tree *SegmentTree) UpdateAt(index int, value interface{}) {
	idx := tree.offset - 1 + index
	tree.data[idx] = value
	for idx >= 1 {
		parent := (idx - 1) / 2
		left := parent*2 + 1
		right := parent*2 + 2
		// out(idx, tree.data[left], tree.data[right], left, right, parent)
		tree.data[parent] = tree.merge(tree.data[left], tree.data[right])
		idx = parent
	}
}

// IntMergerMin :
func IntMergerMin(a, b interface{}) interface{} {
	aInt := a.(int)
	bInt := b.(int)
	if aInt < bInt {
		return aInt
	}
	return bInt
}

// InitIntSegmentTree :
func InitIntSegmentTree(a []int, inf int, merge Merger) *SegmentTree {
	n := len(a)
	vec := make([]interface{}, n)
	for i := 0; i < n; i++ {
		vec[i] = interface{}(a[i])
	}
	return InitSegmentTree(vec, inf, merge)
}

// UpdateIntAt :
func (tree *SegmentTree) UpdateIntAt(index int, value int) {
	tree.UpdateAt(index, interface{}(value))
}

type data struct {
	a, b int
}

func dataMergerMin(a, b interface{}) interface{} {
	x := a.(data)
	y := b.(data)
	if x.a < y.a {
		return x
	}
	return y
}

func dataInitSegmentTree(a []data, inf data, merge Merger) *SegmentTree {
	n := len(a)
	vec := make([]interface{}, n)
	for i := 0; i < n; i++ {
		vec[i] = interface{}(a[i])
	}
	return InitSegmentTree(vec, inf, merge)
}

func (tree *SegmentTree) dataUpdateAt(index int, value data) {
	tree.UpdateAt(index, interface{}(value))
}

func main() {
	seg := InitIntSegmentTree(
		[]int{1, 9, 5, 3, 7, 2, 4, 6, 8},
		1e9, IntMergerMin)
	fmt.Println(seg)
	fmt.Println(seg.data[15:])
	seg.UpdateIntAt(0, 0)
	fmt.Println(seg)
	fmt.Println(seg.data[15:])

	fmt.Println(seg.GetRange(1, 2))
	fmt.Println(seg.GetRange(3, 5))
	fmt.Println(seg.GetRange(0, 9))
	fmt.Println(seg.GetRange(5, 7))
	fmt.Println(seg.GetRange(2, 4))
	fmt.Println(seg)
	seg.UpdateIntAt(1, 2)
	fmt.Println(seg)
	seg.UpdateIntAt(2, 0)
	fmt.Println(seg)

	fmt.Println(seg.GetRange(1, 4))

	segA := dataInitSegmentTree(
		[]data{data{1, 2}, data{2, 3}, data{5, 3}, data{10, 4}, data{1, 3}},
		data{1e9, 1e9}, dataMergerMin)
	fmt.Println(segA)
	fmt.Println(segA.GetRange(0, 4))
	fmt.Println(segA.GetRange(2, 4))
	fmt.Println(segA)
	segA.dataUpdateAt(0, data{0, 3})
	fmt.Println(segA)
}
