package main

import (
	"errors"
	"fmt"
	"log"
)

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

// スタック
type stack struct {
	pos    int
	len    int
	maxlen int
	val    []interface{}
}

func (q *stack) init(v int) {
	q.maxlen = v
	q.pos = 0
	q.val = make([]interface{}, v)
}

func (q *stack) push(v interface{}) {
	if q.pos == q.maxlen {
		log.Fatal("stack is overflow")
	}
	q.val[q.pos] = v
	q.pos++
}

func (q *stack) pop() (interface{}, error) {
	if q.pos == 0 {
		return nil, errors.New("empty")
	}
	q.pos--
	ret := q.val[q.pos]
	return ret, nil
}

func (q *stack) empty() bool {
	if q.pos == 0 {
		return true
	}
	return false
}

func (q *stack) full() bool {
	if q.pos == q.maxlen {
		return true
	}
	return false
}

func main() {
	var s stack
	s.init(5)

	for i := 0; !s.full(); i++ {
		s.push(i)
	}

	for !s.empty() {
		v, _ := s.pop()
		fmt.Println(v.(int))
	}
}

/*func main() {

	var q queue
	q.init(10)
	q.push(10)
	q.pop()
	q.push(20)
	q.push(30)
	q.push(40)

	go func() {
		x := 0
		for {
			if !q.full() {
				q.push(x)
				x++
			}
			if x == 100 {
				break
			}
		}
	}()

	for {
		if !q.empty() {
			v, _ := q.pop()
			fmt.Println(v.(int))
			if v == 99 {
				break
			}
		}
	}
}
*/
