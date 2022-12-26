package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	sc.Buffer(make([]byte, 1000), 500000)
	sc.Split(bufio.ScanWords)
	n, m, k := nextInt(), nextInt(), nextInt()
	a := make([]int, n)
	for i := range a {
		a[i] = nextInt()
	}
	bt := NewBinaryTrie(32)
	for i := 0; i < m; i++ {
		bt.Add(a[i], 1)
	}
	sumk := 0
	for i := 0; i < k; i++ {
		sumk += bt.GetSmallKth(i + 1)
	}
	ans := make([]int, n-m+1)
	ans[0] = sumk
	for i := 1; i <= n-m; i++ {
		idx := bt.LowerBound(a[i-1])
		if idx < k {
			sumk -= a[i-1]
			sumk += bt.GetSmallKth(k + 1)
		}
		bt.Sub(a[i-1], 1)
		bt.Add(a[i+m-1], 1)
		idx = bt.LowerBound(a[i+m-1])
		if idx < k {
			sumk += a[i+m-1]
			sumk -= bt.GetSmallKth(k + 1)
		}
		ans[i] = sumk
	}
	ansstr := make([]string, len(ans))
	for i := range ans {
		ansstr[i] = strconv.Itoa(ans[i])
	}
	fmt.Println(strings.Join(ansstr, " "))
}

type BinaryTrie struct {
	root    *BinaryTrieNode
	bitsize int
}

type BinaryTrieNode struct {
	cnt         int
	left, right *BinaryTrieNode
}

func NewBinaryTrie(bitsize int) *BinaryTrie {
	return &BinaryTrie{
		root:    NewBinaryTrieNode(),
		bitsize: bitsize,
	}
}

func NewBinaryTrieNode() *BinaryTrieNode {
	return &BinaryTrieNode{}
}

func (t *BinaryTrie) Add(x, num int) {
	tmp := t.root
	for i := t.bitsize - 1; i >= 0; i-- {
		tmp.cnt += num
		if (x>>i)&1 == 0 {
			if tmp.left == nil {
				tmp.left = NewBinaryTrieNode()
			}
			tmp = tmp.left
		} else {
			if tmp.right == nil {
				tmp.right = NewBinaryTrieNode()
			}
			tmp = tmp.right
		}
	}
	tmp.cnt += num
}

func (t *BinaryTrie) Sub(x, num int) {
	t.Add(x, -num)
}

func (t *BinaryTrie) Set(x, num int) {
	tmp := t.Find(x)
	t.Add(x, num-tmp)
}

func (t *BinaryTrie) Remove(x int) {
	t.Set(x, 0)
}

func (t *BinaryTrie) Find(x int) int {
	tmp := t.root
	for i := t.bitsize - 1; i >= 0; i-- {
		if (x>>i)&1 == 0 {
			if tmp.left == nil {
				return 0
			}
			tmp = tmp.left
		} else {
			if tmp.right == nil {
				return 0
			}
			tmp = tmp.right
		}
	}
	return tmp.cnt
}

func (t *BinaryTrie) GetMax() int {
	tmp := t.root
	if t.root.cnt == 0 {
		return -1
	}
	result := 0
	for i := 0; i < t.bitsize; i++ {
		result <<= 1
		if tmp.right == nil || tmp.right.cnt == 0 {
			tmp = tmp.left
		} else {
			tmp = tmp.right
			result += 1
		}
	}
	return result
}

func (t *BinaryTrie) GetMin() int {
	tmp := t.root
	if t.root.cnt == 0 {
		return -1
	}
	result := 0
	for i := 0; i < t.bitsize; i++ {
		result <<= 1
		if tmp.left == nil || tmp.left.cnt == 0 {
			tmp = tmp.right
			result += 1
		} else {
			tmp = tmp.left
		}
	}
	return result
}

func (t *BinaryTrie) GetSmallKth(k int) int {
	tmp := t.root
	if k > t.root.cnt {
		return -1
	}
	pos := k
	result := 0
	for i := 0; i < t.bitsize; i++ {
		result <<= 1
		if tmp.left == nil {
			tmp = tmp.right
			result += 1
		} else if tmp.left.cnt < pos {
			pos -= tmp.left.cnt
			tmp = tmp.right
			result += 1
		} else {
			tmp = tmp.left
		}
	}
	return result
}

func (t *BinaryTrie) GetBigKth(k int) int {
	if k > t.root.cnt {
		return -1
	}
	sKth := t.root.cnt - k + 1
	return t.GetSmallKth(sKth)
}

func (t *BinaryTrie) LowerBound(x int) int {
	tmp := t.root
	result := 0
	for i := t.bitsize - 1; i >= 0; i-- {
		if (x>>i)&1 == 0 {
			if tmp.left == nil {
				return result
			}
			tmp = tmp.left
		} else {
			if tmp.left != nil {
				result += tmp.left.cnt
			}
			if tmp.right == nil {
				return result
			}
			tmp = tmp.right
		}
	}
	return result
}

func (t *BinaryTrie) LowerBoundElm(x int) int {
	lb := t.LowerBound(x)
	if lb == 0 {
		return -1
	}
	return t.GetSmallKth(lb)
}

func (t *BinaryTrie) UpperBound(x int) int {
	return t.LowerBound(x + 1)
}

func (t *BinaryTrie) UpperBoundElm(x int) int {
	return t.LowerBoundElm(x + 1)
}

func nextInt() int {
	sc.Scan()
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return i
}

func nextString() string {
	sc.Scan()
	return sc.Text()
}
