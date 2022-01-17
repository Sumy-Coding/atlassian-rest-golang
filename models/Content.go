package models

type Space struct {
	Id int64 `json:"id"`
}

type Content struct {
	Id     string `json:"id"`
	Type   string `json:"type"`
	Status string `json:"status"`
	Title  string `json:"title"`
	Space  Space  `json:"space"`
}
