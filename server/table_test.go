package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
)

func progressCmpOpt() (cmp.Option, cmp.Option) {
	return cmp.AllowUnexported(Progress{}), cmpopts.IgnoreFields(Progress{}, "mux")
}

func TestNewTable(t *testing.T) {
	table := NewTable()
	opt1, opt2 := progressCmpOpt()

	diff := cmp.Diff(table.table, map[string]*Progress{}, opt1, opt2)
	if diff != "" {
		t.Errorf("\n%s", diff)
	}
	diff = cmp.Diff(table.keys, []string{})
}

func TestAddDelete(t *testing.T) {
	table := NewTable()
	opt1, opt2 := progressCmpOpt()

	key := "key"
	pg := NewProgress("test")
	key2 := "key2"
	pg2 := NewProgress("test2")

	// Add key pg
	table.Add(key, pg)

	diff := cmp.Diff(table.table, map[string]*Progress{key: pg}, opt1, opt2)
	if diff != "" {
		t.Errorf("\n%s", diff)
	}
	diff = cmp.Diff(table.keys, []string{key})
	if diff != "" {
		t.Errorf("\n%s", diff)
	}

	// Add key2 pg2
	table.Add(key2, pg2)

	diff = cmp.Diff(table.table, map[string]*Progress{key: pg, key2: pg2}, opt1, opt2)
	if diff != "" {
		t.Errorf("\n%s", diff)
	}
	diff = cmp.Diff(table.keys, []string{key, key2})
	if diff != "" {
		t.Errorf("\n%s", diff)
	}

	// Delete key pg
	table.Delete(key)

	diff = cmp.Diff(table.table, map[string]*Progress{key2: pg2}, opt1, opt2)
	if diff != "" {
		t.Errorf("\n%s", diff)
	}
	diff = cmp.Diff(table.keys, []string{key2})
	if diff != "" {
		t.Errorf("\n%s", diff)
	}

	// Delete key2 pg2
	table.Delete(key2)

	diff = cmp.Diff(table.table, map[string]*Progress{}, opt1, opt2)
	if diff != "" {
		t.Errorf("\n%s", diff)
	}
	diff = cmp.Diff(table.keys, []string{})
	if diff != "" {
		t.Errorf("\n%s", diff)
	}
}

func TestTableCancel(t *testing.T) {
	table := NewTable()
	key := "key"
	pg := NewProgress("test")

	table.Add(key, pg)
	table.Cancel(key)

	diff := cmp.Diff(pg.Canceled, true)
	if diff != "" {
		t.Errorf("\n%s", diff)
	}
}

func TestToItems(t *testing.T) {
	table := NewTable()
	key := "key"
	pg := NewProgress("test")
	opt := cmp.AllowUnexported(Item{})

	table.Add(key, pg)

	items := table.ToItems()
	diff := cmp.Diff(items, []Item{{ID: "key", Name: "test"}}, opt)
	if diff != "" {
		t.Errorf("\n%s", diff)
	}
}
