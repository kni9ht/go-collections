package List

import (
	"fmt"
)

type node struct {
	data int32
	nxt  *node
	prv  *node
}

type List struct {
	head *node
	tail *node
	size int32
}

func (ls *List) Clear() {
	ls.head = nil
	ls.size = 0
	ls.tail = nil
}

func (ls *List) Pop() {
	ls.tail = ls.tail.prv
	ls.tail.nxt = nil
	ls.size--
}

func (ls *List) Remove(T int32, all bool) {
	st := ls.head
	for i := 0; i < int(ls.size); i++ {
		if st.data == T {
			if st.prv != nil {
				st.prv.nxt = st.nxt
			} else {
				ls.head = st.nxt
				ls.head.prv = nil
			}
			if st.nxt != nil {
				st.nxt.prv = st.prv
			} else {
				ls.tail = st.prv
				ls.tail.nxt = nil
			}
			ls.size--
			if !all {
				return
			}
		}
		st = st.nxt
	}
}

func (ls *List) Search(T int32) bool {
	st := ls.head
	for i := 0; i < int(ls.size); i++ {
		if st.data == T {
			return true
		}
		if st.nxt != nil {
			st = st.nxt
		} else {
			break
		}
	}
	return false
}

func (ls *List) Push(i int32) {
	var tmp = node{
		data: i,
		nxt:  nil,
	}
	if ls.tail != nil {
		tmp.prv = ls.tail
		ls.tail.nxt = &tmp
		ls.tail = &tmp
	} else {
		ls.head = &tmp
		ls.tail = &tmp
	}
	ls.size++
}

func (ls *List) Print() {
	var st *node = ls.head
	for i := 0; i < int(ls.size); i++ {
		fmt.Printf("%d ", st.data)
		if st.nxt != nil {
			st = st.nxt
		} else {
			break
		}
	}
	fmt.Println()
}

func New() List {
	var l List
	return l
}
