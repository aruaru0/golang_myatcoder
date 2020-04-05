package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
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

// キュー
type queue struct {
	wpos, rpos int
	len        int
	maxlen     int
	val        []interface{}
}

func (q *queue) init(v int) {
	q.maxlen = v
	q.len = 0
	q.wpos = 0
	q.rpos = 0
	q.val = make([]interface{}, v)
}

func (q *queue) push(v interface{}) {
	if q.len == q.maxlen {
		log.Fatal("queue is overflow")
	}
	q.val[q.wpos] = v
	q.wpos++
	q.len++
	if q.wpos == q.maxlen {
		q.wpos = 0
	}
}

func (q *queue) pop() (interface{}, error) {
	if q.len == 0 {
		return nil, errors.New("empty")
	}
	ret := q.val[q.rpos]
	q.len--
	q.rpos++
	if q.rpos == q.maxlen {
		q.rpos = 0
	}
	return ret, nil
}

func (q *queue) empty() bool {
	if q.len == 0 {
		return true
	}
	return false
}

func (q *queue) full() bool {
	if q.len == q.maxlen {
		return true
	}
	return false
}

func main() {
	sc.Split(bufio.ScanWords)
	sc.Buffer([]byte{}, 1000000)

	K := getInt()

	var q queue
	q.init(300000)
	for i := 1; i < 10; i++ {
		q.push(i)
	}
	last := 0
	for i := 0; i < K; i++ {
		v, _ := q.pop()
		x := v.(int)
		// out(x)
		t := x % 10
		if t != 0 {
			q.push(x*10 + t - 1)
		}
		q.push(x*10 + t)
		if t != 9 {
			q.push(x*10 + t + 1)
		}
		last = x
	}
	out(last)
}
