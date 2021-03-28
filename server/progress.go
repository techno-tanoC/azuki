package main

import (
	"fmt"
	"sync"
)

type Item struct {
	Name     string `json:"name"`
	Total    int64  `json:"total"`
	Size     int64  `json:"size"`
	Canceled bool   `json:"canceled"`
}

type Progress struct {
	mux sync.Mutex
	Item
}

func NewProgress(name string) *Progress {
	item := Item{Name: name}
	return &Progress{Item: item}
}

func (pg *Progress) Write(p []byte) (int, error) {
	pg.mux.Lock()
	defer pg.mux.Unlock()

	if pg.Canceled {
		return 0, fmt.Errorf("canceled: %s", pg.Name)
	}

	n := len(p)
	pg.Size += int64(n)
	return n, nil
}

func (pg *Progress) SetName(name string) {
	pg.mux.Lock()
	defer pg.mux.Unlock()

	pg.Name = name
}

func (pg *Progress) SetTotal(total int64) {
	pg.mux.Lock()
	defer pg.mux.Unlock()

	if total > 0 {
		pg.Total = total
	} else {
		pg.Total = 0
	}
}

func (pg *Progress) Cancel() {
	pg.mux.Lock()
	defer pg.mux.Unlock()

	pg.Canceled = true
}

func (pg *Progress) ToItem() Item {
	pg.mux.Lock()
	defer pg.mux.Unlock()

	return pg.Item
}
