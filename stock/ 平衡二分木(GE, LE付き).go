import (
	"fmt"
	"math"
	"strings"
)

// 平衡二分探索木
var (
	y uint = 88172645463325252
)

const inf = math.MaxInt64

func xorshift() uint {
	y ^= y << 7
	y ^= y >> 9
	return y
}

type treapNode struct {
	value    int
	priority uint
	count    int
	left     *treapNode
	right    *treapNode
}

func newTreapNode(v int) *treapNode {
	return &treapNode{v, xorshift(), 1, nil, nil}
}

func treapRotateRight(n *treapNode) *treapNode {
	l := n.left
	n.left = l.right
	l.right = n
	return l
}

func treapRotateLeft(n *treapNode) *treapNode {
	r := n.right
	n.right = r.left
	r.left = n
	return r
}

func treapInsert(n *treapNode, v int) *treapNode {
	if n == nil {
		return newTreapNode(v)
	}
	if n.value == v {
		n.count++
		return n
	}
	if n.value > v {
		n.left = treapInsert(n.left, v)
		if n.priority > n.left.priority {
			n = treapRotateRight(n)
		}
	} else {
		n.right = treapInsert(n.right, v)
		if n.priority > n.right.priority {
			n = treapRotateLeft(n)
		}
	}
	return n
}

func treapDelete(n *treapNode, v int) *treapNode {
	if n == nil {
		panic("node is not found!")
	}
	if n.value > v {
		n.left = treapDelete(n.left, v)
		return n
	}
	if n.value < v {
		n.right = treapDelete(n.right, v)
		return n
	}

	// n.value == v
	if n.count > 1 {
		n.count--
		return n
	}

	if n.left == nil && n.right == nil {
		return nil
	}

	if n.left == nil {
		n = treapRotateLeft(n)
	} else if n.right == nil {
		n = treapRotateRight(n)
	} else {
		// n.left != nil && n.right != nil
		if n.left.priority < n.right.priority {
			n = treapRotateRight(n)
		} else {
			n = treapRotateLeft(n)
		}
	}
	return treapDelete(n, v)
}

func treapCount(n *treapNode) int {
	if n == nil {
		return 0
	}
	return n.count + treapCount(n.left) + treapCount(n.right)
}

func treapString(n *treapNode) string {
	if n == nil {
		return ""
	}
	result := make([]string, 0)
	if n.left != nil {
		result = append(result, treapString(n.left))
	}
	result = append(result, fmt.Sprintf("%d:%d", n.value, n.count))
	if n.right != nil {
		result = append(result, treapString(n.right))
	}
	return strings.Join(result, " ")
}

func treapMin(n *treapNode) int {
	if n.left != nil {
		return treapMin(n.left)
	}
	return n.value
}

func treapLEMax(n *treapNode, v int) int {
	ret := -inf
	t := n
	for t != nil {
		if t.value <= v {
			ret = max(ret, t.value)
		}
		if t.value > v {
			t = t.left
		} else {
			t = t.right
		}
	}

	return ret
}

func treapGEMin(n *treapNode, v int) int {
	ret := inf
	t := n
	for t != nil {
		if t.value >= v {
			ret = min(ret, t.value)
		}
		if t.value < v {
			t = t.right
		} else {
			t = t.left
		}
	}

	return ret
}

func treapMax(n *treapNode) int {
	if n.right != nil {
		return treapMax(n.right)
	}
	return n.value
}

type treap struct {
	root *treapNode
	size int
}

func (t *treap) Insert(v int) {
	t.root = treapInsert(t.root, v)
	t.size++
}

func (t *treap) Delete(v int) {
	t.root = treapDelete(t.root, v)
	t.size--
}

func (t *treap) String() string {
	return treapString(t.root)
}

func (t *treap) Count() int {
	return t.size
}

func (t *treap) Min() int {
	return treapMin(t.root)
}

func (t *treap) LEMax(v int) int {
	return treapLEMax(t.root, v)
}

func (t *treap) Max() int {
	return treapMax(t.root)
}

func (t *treap) GEMin(v int) int {
	return treapGEMin(t.root, v)
}

/*
func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1000000)

	t := &treap{}
	a := make([]int, 0)
	for i := 0; i < 10; i++ {
		x := rand.Intn(100)
		a = append(a, x)
		t.Insert(x)
	}
	// out(t.LE(20))
	// return
	sort.Ints(a)
	out(t)
	out(a)

	out("min", t.Min(), "max", t.Max())
	for i := 0; i < 100; i++ {
		out(i, "LEMax", t.LEMax(i), "GEMax", t.GEMin(i))
	}
}
*/