package main

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestNewProgress(t *testing.T) {
	name := "test"
	pg := NewProgress(name)

	diff := cmp.Diff(pg.name, name)
	if diff != "" {
		t.Errorf("TestNewProgress: (-got +want)\n%s", diff)
	}
}

func TestWrite(t *testing.T) {
	pg := NewProgress("test")
	bytes := []byte("hello")

	_, err := pg.Write(bytes)
	if err != nil {
		t.Errorf("TestWrite: err %v", err)
	}

	diff := cmp.Diff(pg.size, int64(len(bytes)))
	if diff != "" {
		t.Errorf("TestWrite2: (-got +want)\n%s", diff)
	}

	_, err = pg.Write(bytes)
	if err != nil {
		t.Errorf("TestWrite3: err %v", err)
	}

	diff = cmp.Diff(pg.size, 2*int64(len(bytes)))
	if diff != "" {
		t.Errorf("TestWrite2: (-got +want)\n%s", diff)
	}
}

func TestSetTotal(t *testing.T) {
	pg := NewProgress("test")

	pg.SetTotal(100)
	diff := cmp.Diff(pg.total, int64(100))
	if diff != "" {
		t.Errorf("TestSeTotal: (-got +want)\n%s", diff)
	}

	pg.SetTotal(1000)
	diff = cmp.Diff(pg.total, int64(1000))
	if diff != "" {
		t.Errorf("TestSetTotal2: (-got +want)\n%s", diff)
	}
}

func TestCancel(t *testing.T) {
	pg := NewProgress("test")
	diff := cmp.Diff(pg.canceled, false)
	if diff != "" {
		t.Errorf("TestCancel2: (-got +want)\n%s", diff)
	}

	pg.Cancel()

	diff = cmp.Diff(pg.canceled, true)
	if diff != "" {
		t.Errorf("TestCancel2: (-got +want)\n%s", diff)
	}
}

func TestToItem(t *testing.T) {
	pg := NewProgress("test")
	item := pg.ToItem("key")
	expected := Item{
		ID:       "key",
		Name:     "test",
		Total:    0,
		Size:     0,
		Canceled: false,
	}

	diff := cmp.Diff(item, expected)
	if diff != "" {
		t.Errorf("TestToItem: (-got +want)\n%s", diff)
	}
}
