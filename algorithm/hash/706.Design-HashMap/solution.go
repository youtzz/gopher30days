package problem0706

import "container/list"

type MyHashMap struct {
	set []list.List
}
type entry struct {
	key, value int
}

const base = 769

func Constructor() MyHashMap {
	return MyHashMap{set: make([]list.List, base)}
}

func (*MyHashMap) hash(key int) int {
	return key % base
}

func (m *MyHashMap) Put(key int, value int) {
	var flag bool
	insert := &entry{key, value}
	h := m.hash(key)
	for e := m.set[h].Front(); e != nil; e = e.Next() {
		if e.Value.(*entry).key == key {
			e.Value = insert
			flag = true
		}
	}
	if !flag {
		m.set[h].PushBack(insert)
	}
}

func (m *MyHashMap) Get(key int) int {
	h := m.hash(key)
	for e := m.set[h].Front(); e != nil; e = e.Next() {
		if e.Value.(*entry).key == key {
			return e.Value.(*entry).value
		}
	}
	return -1
}

func (m *MyHashMap) Remove(key int) {
	h := m.hash(key)
	for e := m.set[h].Front(); e != nil; e = e.Next() {
		if e.Value.(*entry).key == key {
			m.set[h].Remove(e)
		}
	}
}
