package main

import (
	"fmt"
	"sync"
)

type Progress struct {
	mux      sync.Mutex
	name     string
	total    int64
	size     int64
	canceled bool
}

func NewProgress(name string) *Progress {
	return &Progress{name: name}
}

func (pg *Progress) Write(p []byte) (int, error) {
	pg.mux.Lock()
	defer pg.mux.Unlock()

	if pg.canceled {
		return 0, fmt.Errorf("canceled: %s", pg.name)
	}

	n := len(p)
	pg.size += int64(n)
	return n, nil
}

func (pg *Progress) SetTotal(total int64) {
	pg.mux.Lock()
	defer pg.mux.Unlock()

	if total > 0 {
		pg.total = total
	} else {
		pg.total = 0
	}
}

func (pg *Progress) Cancel() {
	pg.mux.Lock()
	defer pg.mux.Unlock()

	pg.canceled = true
}

func (pg *Progress) ToItem(key string) Item {
	pg.mux.Lock()
	defer pg.mux.Unlock()

	return Item{
		ID:       key,
		Name:     pg.name,
		Total:    pg.total,
		Size:     pg.size,
		Canceled: pg.canceled,
	}
}
