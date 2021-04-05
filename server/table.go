package main

import (
	"sync"
)

type Table struct {
	mux   sync.Mutex
	table map[string]*Progress
	keys  []string
}

func NewTable() *Table {
	return &Table{
		table: map[string]*Progress{},
		keys:  []string{},
	}
}

func (t *Table) Add(key string, pg *Progress) {
	t.mux.Lock()
	defer t.mux.Unlock()

	t.table[key] = pg
	t.keys = append(t.keys, key)
}

func (t *Table) Delete(key string) {
	t.mux.Lock()
	defer t.mux.Unlock()

	if _, present := t.table[key]; present {
		delete(t.table, key)

		keys := []string{}
		for _, k := range t.keys {
			if k != key {
				keys = append(keys, k)
			}
		}
		t.keys = keys
	}
}

func (t *Table) Cancel(key string) {
	if _, present := t.table[key]; present {
		t.table[key].Cancel()
		t.Delete(key)
	}
}

func (t *Table) ToItems() []Item {
	t.mux.Lock()
	defer t.mux.Unlock()

	items := make([]Item, 0)
	for _, k := range t.keys {
		item := t.table[k].ToItem(k)
		items = append(items, item)
	}
	return items
}
