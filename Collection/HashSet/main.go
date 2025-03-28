package HashSet

import List "go-collections/Collection/Linklist"

type set struct {
	arr [8192]List.List
}

func (m *set) Push(T int32) {
	i := T % 8192
	if !m.Find(T) {
		m.arr[i].Push(T)
	}
}

func (m *set) Find(T int32) bool {
	i := T % 8192
	return m.arr[i].Search(T)
}

func (m *set) Remove(T int32) {
	i := T % 8192
	m.arr[i].Remove(T, false)
}

func New() set {
	var s set
	return s
}
