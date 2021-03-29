package main

type Item struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Total    int64  `json:"total"`
	Size     int64  `json:"size"`
	Canceled bool   `json:"canceled"`
}

type IntoItem interface {
	ToItem(string) Item
}
