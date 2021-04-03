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
	key := "key"
	pg := NewProgress("test")
	opt1, opt2 := progressCmpOpt()

	table.Add("key", pg)

	diff := cmp.Diff(table.table, map[string]*Progress{key: pg}, opt1, opt2)
	if diff != "" {
		t.Errorf("\n%s", diff)
	}
	diff = cmp.Diff(table.keys, []string{"key"})
	if diff != "" {
		t.Errorf("\n%s", diff)
	}

	table.Delete(key)

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

	diff := cmp.Diff(pg.canceled, true)
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
