package main

import (
	"sync"
)

type Table struct {
	mux   sync.Mutex
	table map[string]ProgressLike
	keys  []string
}

func NewTable() *Table {
	return &Table{
		table: map[string]ProgressLike{},
		keys:  []string{},
	}
}

func (t *Table) Add(key string, pg ProgressLike) {
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

		t.keys = []string{}
		for _, k := range t.keys {
			if k != key {
				t.keys = append(t.keys, k)
			}
		}
	}
}

func (t *Table) Cancel(key string) {
	t.table[key].Cancel()
	t.Delete(key)
}

func (t *Table) ToItems() []Item {
	t.mux.Lock()
	defer t.mux.Unlock()

	var items []Item
	for _, k := range t.keys {
		item := t.table[k].ToItem(k)
		items = append(items, item)
	}
	return items
}
